package api

import (
	"os"

	"gopkg.in/yaml.v2"
)

type YAMLFile struct {
	File string
}

// Decode
func (y *YAMLFile) Decode(d Decoder) error {
	file, err := os.OpenFile(y.File, os.O_RDONLY, 0)
	if err == nil {
		defer func() { err = file.Close() }()
		err = yaml.NewDecoder(file).Decode(d)
	}
	return err
}

// Encode
func (y *YAMLFile) Encode(e Encoder) error {
	file, err := os.OpenFile(y.File, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err == nil {
		defer func() { err = file.Close() }()
		err = yaml.NewEncoder(file).Encode(e)
	}
	return err
}
