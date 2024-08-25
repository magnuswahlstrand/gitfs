package gitfs

import (
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
	"path/filepath"
)

type InMemoryGitDirectory struct {
	path      string
	repoURL   string
	worktree  *git.Worktree
	Filenames []string
}

func (d *InMemoryGitDirectory) Open(filename string) (io.ReadCloser, error) {
	f, err := d.worktree.Filesystem.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	return f, nil
}

func NewGit(URL string, path string) (*InMemoryGitDirectory, error) {
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

	// Check that the path exists and is a path
	dir, err := worktree.Filesystem.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read path: %w", err)
	}
	if !dir.IsDir() {
		return nil, fmt.Errorf("%q is not a path", path)
	}

	files, err := worktree.Filesystem.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read path: %w", err)
	}

	fileNames := []string{}
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}

		file, err := worktree.Filesystem.Open(filepath.Join(path, fileInfo.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed during loading of file names: %w", err)
		}
		fileNames = append(fileNames, file.Name())
	}

	return &InMemoryGitDirectory{
		worktree:  worktree,
		Filenames: fileNames,
		repoURL:   URL,
		path:      path,
	}, nil
}
