package gitfs

import (
	"io/fs"
	"os"
)

var _ fs.DirEntry = &fsDirEntryWrapper{}

// os.FileInfo wrapper that implements the fs.DirEntry interface
type fsDirEntryWrapper struct {
	os.FileInfo
}

func (f fsDirEntryWrapper) Name() string {
	return f.FileInfo.Name()
}

func (f fsDirEntryWrapper) IsDir() bool {
	return f.FileInfo.IsDir()
}

func (f fsDirEntryWrapper) Type() fs.FileMode {
	return f.FileInfo.Mode()
}

func (f fsDirEntryWrapper) Info() (fs.FileInfo, error) {
	return f.FileInfo, nil
}
