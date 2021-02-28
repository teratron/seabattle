package lang

import "github.com/teratron/seabattle/pkg/api"

type Lang map[string]string

func (l *Lang) Decode(decoder api.Decoder) error {
	return decoder.Decode(l)
}
