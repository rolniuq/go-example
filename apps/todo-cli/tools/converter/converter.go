package converter

import "encoding/json"

func ByteToType[T any](b []byte) (*T, error) {
	var result T

	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func StringToType[T any](s string) (*T, error) {
	var result T

	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func TypeToString[T any](t T) (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
