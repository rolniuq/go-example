package pool

import "testing"

func BenchmarkPoolAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PoolAlloc()
	}
}

func BenchmarkMakeBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeBuf()
	}
}
