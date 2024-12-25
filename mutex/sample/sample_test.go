package sample

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrongSample(t *testing.T) {
	res := RaceConditionSample()
	require.Equal(t, "Wrong sample 1000", res)
}
