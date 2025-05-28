package mem

import "testing"

func BenchmarkStackAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StackAlloc(1)
	}
}

func BenchmarkHeapAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HeapAlloc(1)
	}
}
