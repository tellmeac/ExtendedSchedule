package cache

import (
	"github.com/vmihailenco/msgpack/v5"
)

// Codec codec interface for cache
type Codec interface {
	// Encode returns the encoded byte array of v.
	Encode(v interface{}) ([]byte, error)

	// Decode analyzes the encoded data and stores the result into the v.
	Decode(data []byte, v interface{}) error
}

var (
	// codec the default codec for the cache
	codec Codec = &msgpackCodec{}
)

type msgpackCodec struct{}

func (*msgpackCodec) Encode(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (*msgpackCodec) Decode(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}

// DefaultCodec returns default codec.
func DefaultCodec() Codec {
	return codec
}
