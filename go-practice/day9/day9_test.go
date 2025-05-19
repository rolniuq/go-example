package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay9_Fib(t *testing.T) {
	d9 := Day9{}
	res := d9.Fib(10)
	require.Equal(t, 55, res)
}
