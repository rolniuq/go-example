package base64

import (
	"base64parser/tools/json"
	"base64parser/tools/pointer"
	eb64 "encoding/base64"
)

type Base64 interface {
	EncodeToString(data any) (*string, error)
	DecodeToString(data string) (*string, error)
}

type base64 struct{}

func NewBase64() Base64 {
	return &base64{}
}

func (b *base64) EncodeToString(data any) (*string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return pointer.NewPointer(eb64.StdEncoding.EncodeToString(bytes)), nil
}

func (b *base64) DecodeToString(data string) (*string, error) {
	bytes, err := eb64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return pointer.NewPointer(string(bytes)), nil
}
