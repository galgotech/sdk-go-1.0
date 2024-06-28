package util

import (
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
)

var defaultIncludePaths atomic.Value

// IncludePaths will return the search path for non-absolute import file
func IncludePaths() []string {
	return defaultIncludePaths.Load().([]string)
}

// SetIncludePaths will update the search path for non-absolute import file
func SetIncludePaths(paths []string) {
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			panic(fmt.Errorf("%s must be an absolute file path", path))
		}
	}

	defaultIncludePaths.Store(paths)
}

func init() {
	// No execute set include path to suport webassembly
	if WebAssembly() {
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	SetIncludePaths([]string{wd})
}
