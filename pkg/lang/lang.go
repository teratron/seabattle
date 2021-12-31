package lang

import "github.com/teratron/seabattle/pkg/util"

type Lang map[string]string

func (l *Lang) Decode(decoder util.Decoder) error {
	return decoder.Decode(l)
}
