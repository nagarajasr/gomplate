package funcs

import (
	"path"
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

// Base -
func (f *PathFuncs) Base(in interface{}) (string, error) {
	return path.Base(conv.ToString(in)), nil
}

// Dir -
func (f *PathFuncs) Dir(in interface{}) (string, error) {
	return path.Dir(conv.ToString(in)), nil
}
