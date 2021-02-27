package api

type DecodeEncoder interface {
	Decoder
	Encoder
}

type Decoder interface {
	Decode(Decoder) error
}

type Encoder interface {
	Encode(Encoder) error
}
