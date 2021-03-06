package util

import (
	"encoding/json"
	"os"
)

type FileJSON struct {
	name string
}

// Decode
func (j *FileJSON) Decode(data interface{}) error {
	file, err := os.OpenFile(j.name, os.O_RDONLY, 0600)
	if err == nil {
		defer func() { err = file.Close() }()
		err = json.NewDecoder(file).Decode(data)
	}
	return err
}

// Encode
func (j *FileJSON) Encode(data interface{}) error {
	file, err := os.OpenFile(j.name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err == nil {
		defer func() { err = file.Close() }()
		err = json.NewEncoder(file).Encode(data)
	}
	return err
}
