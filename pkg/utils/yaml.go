package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type YAMLFile struct {
	file string
}

/*var Unmarshal = func(data interface{}) error {
	file, err := os.OpenFile(y.file, os.O_RDONLY, 0)
	//fmt.Println(file.Fd())
	if err == nil {
		defer func() { err = file.Close() }()
		err = yaml.NewDecoder(file).Decode(data)
	}
	return nil
}

func (y *YAMLFile) UnmarshalYAML(unmarshal func(interface{}) error) error {
	file, err := os.OpenFile(y.file, os.O_RDONLY, 0)
	if err == nil {
		defer func() { err = file.Close() }()
		d := yaml.NewDecoder(file)
		err = d.Decode(data)
	}
	//unmarshal()
	return nil
}*/

// Decode
func (y *YAMLFile) Decode(data interface{}) error {
	file, err := os.OpenFile(y.file, os.O_RDONLY, 0666)
	if err == nil {
		defer func() { err = file.Close() }()
		err = yaml.NewDecoder(file).Decode(data)
	}
	return err
}

// Encode
func (y *YAMLFile) Encode(data interface{}) error {
	file, err := os.OpenFile(y.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err == nil {
		defer func() { err = file.Close() }()
		err = yaml.NewEncoder(file).Encode(data)
	}
	return err
}
