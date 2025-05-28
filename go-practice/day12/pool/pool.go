package pool

import "sync"

var bufPool = sync.Pool{
	New: func() any {
		return make([]byte, 1024)
	},
}

func MakeBuf() []byte {
	return make([]byte, 1024)
}

func PoolAlloc() []byte {
	buf := bufPool.Get().([]byte)
	// use buf
	bufPool.Put(buf)

	return buf
}
