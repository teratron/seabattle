package lang

import "github.com/teratron/seabattle/pkg/utils"

type Lang map[string]string

func (l *Lang) Decode(decoder utils.Decoder) error {
	return decoder.Decode(l)
}
