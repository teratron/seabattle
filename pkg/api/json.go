package api

import (
	"encoding/json"
	"os"
)

type JSONFile struct {
	File string
}

// Decode
func (j *JSONFile) Decode(d Decoder) error {
	file, err := os.OpenFile(j.File, os.O_RDONLY, 0)
	if err == nil {
		defer func() { err = file.Close() }()
		err = json.NewDecoder(file).Decode(d)
	}
	return err
}

// Encode
func (j *JSONFile) Encode(e Encoder) error {
	file, err := os.OpenFile(j.File, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err == nil {
		defer func() { err = file.Close() }()
		err = json.NewEncoder(file).Encode(e)
	}
	return err
}
