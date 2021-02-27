package api

import (
	"fmt"
	"path/filepath"
)

type FileError struct {
	Err error
	DecodeEncoder
}

func (f FileError) Error() string {
	return fmt.Sprintf("file type error: %v\n", f.Err)
}

func GetFileType(file string) DecodeEncoder {
	ext := filepath.Base(filepath.Ext(file))
	switch ext {
	case ".json":
		return &JSONFile{File: file}
	case ".yml":
		return &YAMLFile{File: file}
	default:
		return &FileError{Err: fmt.Errorf("extension isn't defined: %s", ext)}
	}
}
