package funcs

import (
	"sync"

	"github.com/satori/go.uuid"
)

var (
	utilNS     *UtilFuncs
	utilNSInit sync.Once
)

// UtilNS - the util namespace
func UtilNS() *UtilFuncs {
	utilNSInit.Do(func() { utilNS = &UtilFuncs{} })
	return utilNS
}

// AddUtilFuncs -
func AddUtilFuncs(f map[string]interface{}) {
	f["util"] = UtilNS
}

// UtilFuncs -
type UtilFuncs struct{}

func (f *UtilFuncs) UUID() (string, error) {
	retval := uuid.Must(uuid.NewV4())

	return retval.String(), nil
}

