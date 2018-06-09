package funcs

// func Abs(path string) (string, error)
func Base(path string) string
func Clean(path string) string
func Dir(path string) string
func Ext(path string) string
func FromSlash(path string) string
func IsAbs(path string) bool
func Join(elem ...string) string
func Match(pattern, name string) (matched bool, err error)
func Rel(basepath, targpath string) (string, error)
func Split(path string) (dir, file string)
func ToSlash(path string) string
func VolumeName(path string) string
