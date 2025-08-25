package json

import "github.com/bytedance/sonic"

func Marshal(val any) ([]byte, error) {
	return sonic.Marshal(val)
}

func Unmashal(data []byte, val any) error {
	return sonic.Unmarshal(data, val)
}
