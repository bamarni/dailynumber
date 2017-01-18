package dailynumber

import (
	"sync"
	"time"
)

type Counter interface {
	Increment(since time.Time) int
}

type MemoryCounter struct {
	mutex   sync.Mutex
	counter int
	since   time.Time
}

func (c *MemoryCounter) Increment(since time.Time) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if since != c.since {
		c.since = since
		c.counter = 0
	}
	c.counter++
	return c.counter
}
