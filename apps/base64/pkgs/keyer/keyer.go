package keyer

import "fmt"

type Keyer interface {
	GetKey() (string, error)
}

type keyer struct{}

func NewKeyer() Keyer {
	return &keyer{}
}

func (k *keyer) GetKey() (string, error) {
	var key string
	if _, err := fmt.Scanln(&key); err != nil {
		return "", err
	}

	return key, nil
}
