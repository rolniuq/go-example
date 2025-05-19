package day8

import (
	"fmt"
	"go-practice/day8/mathutil"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	cases := []struct {
		name       string
		a, b, want int
	}{
		{"both positive", 1, 2, 3},
		{"both negative", -1, -2, -3},
		{"zeros", 0, 0, 0},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := mathutil.Sum(tc.a, tc.b)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestDivide(t *testing.T) {
	cases := []struct {
		name       string
		a, b, want int
	}{
		{"two numbers", 100, 2, 50},
		{"two numbers", 50, 2, 25},
		{"zero", 0, 2, 0},
		{"zero", 2, 0, 0},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := mathutil.Divide(tc.a, tc.b)
			if tc.b == 0 {
				require.Error(t, err)
				require.Equal(t, fmt.Sprintf("cannot divide %d by %d", tc.a, tc.b), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}

type fakeProcessor struct{}

func (fp fakeProcessor) Process(i int) int {
	return i * 2
}

func TestProcessItems(t *testing.T) {
	t.Run("process items", func(t *testing.T) {
		items := []int{1, 2, 3}
		fp := &fakeProcessor{}
		d := mathutil.ProcessItems(items, fp.Process)
		require.Equal(t, d, []int{2, 4, 6})
	})
}

func FuzzSum(f *testing.F) {
	f.Add(1, 2)
	f.Add(-1, 5)

	f.Fuzz(func(t *testing.T, a, b int) {
		d1 := mathutil.Sum(a, b)
		d2 := mathutil.Sum(b, a)
		require.Equal(t, d1, d2)
	})
}

func BenchmarkProcessItems(b *testing.B) {
	items := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		items[i] = rand.Intn(1000)
	}

	fn := func(x int) int { return x * x }
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mathutil.ProcessItems(items, fn)
	}
}
