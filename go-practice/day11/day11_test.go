package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay11_CopyStruct(t *testing.T) {
	type a struct {
		Name string
		Age  int
	}

	type b struct {
		Name string
		Old  int
	}

	d := Day11{}
	x := &a{}

	y := &b{
		Name: "b",
		Old:  2,
	}

	err := d.CopyStruct(x, y)
	require.NoError(t, err)
}
