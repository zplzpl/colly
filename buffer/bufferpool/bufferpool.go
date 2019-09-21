package bufferpool

import "github.com/gocolly/colly/buffer"

var (
	_pool = buffer.NewPool()
	// Get retrieves a buffer from the pool, creating one if necessary.
	Get = _pool.Get
)
