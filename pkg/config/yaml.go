package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func (cfg *Config) Decode(path string) error {
	file, err := os.Open(path)
	if err == nil {
		defer func() {
			err = file.Close()
		}()
		var info os.FileInfo
		if info, err = file.Stat(); err == nil && !info.IsDir() {
			err = yaml.NewDecoder(file).Decode(&cfg)
		}
	}
	return err
}

/*
func (cfg *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return nil
}

func (d *decoder) callUnmarshaler(n *node, u Unmarshaler) (good bool) {
	terrlen := len(d.terrors)
	err := u.UnmarshalYAML(func(v interface{}) (err error) {
		defer handleErr(&err)
		d.unmarshal(n, reflect.ValueOf(v))
		if len(d.terrors) > terrlen {
			issues := d.terrors[terrlen:]
			d.terrors = d.terrors[:terrlen]
			return &TypeError{issues}
		}
		return nil
	})
	if e, ok := err.(*TypeError); ok {
		d.terrors = append(d.terrors, e.Errors...)
		return false
	}
	if err != nil {
		fail(err)
	}
	return true
}*/
