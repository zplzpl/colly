package buffer

import "bytes"

const _size = 4096 // by default, create 4096 bytes buffers

type Buffer struct {
	*bytes.Buffer
	pool Pool
}

// Free returns the Buffer to its Pool.
//
// Callers must not retain references to the Buffer after calling Free.
func (b *Buffer) Free() {
	b.pool.put(b)
}