package api

/*type PathYAML string

func (p PathYAML) Decode(u yaml.Unmarshaler) error {
	file, err := os.Open(string(p))
	if err == nil {
		defer func() {
			err = file.Close()
		}()
		var info os.FileInfo
		if info, err = file.Stat(); err == nil && !info.IsDir() {
			err = yaml.NewDecoder(file).Decode(&u)
		}
	}
	return err
}*/
