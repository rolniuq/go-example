package day8

import (
	"go-practice/day8/mathutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	d := mathutil.Sum(1, 2)
	require.Equal(t, d, 3)
}
