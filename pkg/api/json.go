package api

import (
	"encoding/json"
	"os"
)

type JSONFile struct {
	file string
}

// Decode
func (j *JSONFile) Decode(data interface{}) error {
	file, err := os.OpenFile(j.file, os.O_RDONLY, 0666)
	if err == nil {
		defer func() { err = file.Close() }()
		err = json.NewDecoder(file).Decode(data)
	}
	return err
}

// Encode
func (j *JSONFile) Encode(data interface{}) error {
	file, err := os.OpenFile(j.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err == nil {
		defer func() { err = file.Close() }()
		err = json.NewEncoder(file).Encode(data)
	}
	return err
}
