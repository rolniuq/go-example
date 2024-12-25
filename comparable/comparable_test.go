package comparabletest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerson(t *testing.T) {
	p := NewPerson[string]()
	p.SetName("Alice")
	p.SetAge(20)
	p.SetBonus("Hello")

	require.Equal(t, "Alice", p.GetName())
	require.Equal(t, 20, p.GetAge())
	require.True(t, p.IsAdult())
	require.Equal(t, "Hello", p.GetBonus())
}
