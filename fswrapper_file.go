package gitfs

import (
	"github.com/go-git/go-billy/v5"
	"io/fs"
)

var _ fs.File = &fsFileWrapper{}

// fsFileWrapper wraps billy.File and implements fs.File interface
type fsFileWrapper struct {
	billy.File
	statFunc func() (fs.FileInfo, error)
}

func (f *fsFileWrapper) Stat() (fs.FileInfo, error) {
	return f.statFunc()
}

func (f *fsFileWrapper) Read(bytes []byte) (int, error) {
	return f.File.Read(bytes)
}

func (f *fsFileWrapper) Close() error {
	return f.File.Close()
}
