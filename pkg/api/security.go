package api

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Scope struct {
	Root string
}

func (s *Scope) LoadFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	path := getNormPath(name, s.Root)
	if !isInProject(path, s.Root) {
		return nil, &OutOfScopeError{}
	}

	file, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func getNormPath(name string, root string) string {
	var absPath string
	if path.IsAbs(name) {
		absPath = name
	} else {
		absPath = filepath.Join(root, name)
	}

	return filepath.Clean(absPath)
}

func isInProject(norm string, root string) bool {
	return strings.HasPrefix(norm, root)
}
