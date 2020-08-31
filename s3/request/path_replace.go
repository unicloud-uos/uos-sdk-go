package request

// PathReplace replaces path elements using field buffers
type PathReplace struct {
	// May mutate path slice
	path     []byte
	rawPath  []byte
	fieldBuf []byte
}

func NewPathReplace(path string) PathReplace {
	return PathReplace{
		path:    []byte(path),
		rawPath: []byte(path),
	}
}

func (r *PathReplace) Encode() (path string, rawPath string) {
	return string(r.path), string(r.rawPath)
}
