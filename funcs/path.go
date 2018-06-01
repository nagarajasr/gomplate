package funcs

import (
	"path/filepath"
	"sync"

	"github.com/hairyhenderson/gomplate/conv"
)

var (
	pf     *PathFuncs
	pfInit sync.Once
)

// PathNS - the Path namespace
func PathNS() *PathFuncs {
	pfInit.Do(func() { pf = &PathFuncs{} })
	return pf
}

// AddPathFuncs -
func AddPathFuncs(f map[string]interface{}) {
	f["path"] = PathNS
}

// PathFuncs -
type PathFuncs struct {
}

// Basename -
func (f *PathFuncs) Basename(in interface{}) (string, error) {
	return filepath.Base(conv.ToString(in)), nil
}

// Dirname -
func (f *PathFuncs) Dirname(in interface{}) (string, error) {
	return filepath.Dir(conv.ToString(in)), nil
}
