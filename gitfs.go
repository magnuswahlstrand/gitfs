package gitfs

import (
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"io/fs"
)

var _ fs.FS = &FS{}
var _ fs.StatFS = &FS{}

// FS wraps a git work tree. Compatible with fs.FS and fs.StatFS
type FS struct {
	worktree *git.Worktree
}

func (d *FS) Stat(name string) (fs.FileInfo, error) {
	return d.worktree.Filesystem.Stat(name)
}

func (d *FS) ReadDir(name string) ([]fs.DirEntry, error) {
	f, err := d.worktree.Filesystem.ReadDir(name)
	if err != nil {
		return nil, err
	}

	var entries []fs.DirEntry
	for _, e := range f {
		entries = append(entries, fsDirEntryWrapper{
			FileInfo: e,
		})
	}

	return entries, nil
}

func (d *FS) Open(name string) (fs.File, error) {
	f, err := d.worktree.Filesystem.Open(name)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	return &fsFileWrapper{
		File: f,
		statFunc: func() (fs.FileInfo, error) {
			return d.worktree.Filesystem.Stat(name)
		},
	}, nil
}

func New(URL string) (fs.StatFS, error) {
	repo, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL: URL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to clone repository: %w", err)
	}

	// Get the working path (worktree)
	worktree, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get worktree: %w", err)
	}

	return &FS{worktree: worktree}, nil
}
