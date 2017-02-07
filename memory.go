package dailynumber

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Memory struct {
	mutex   sync.Mutex
	buckets [8]int
	day     int
}

func (m *Memory) Generate() string {
	i := rand.Intn(8)

	m.mutex.Lock()
	defer m.mutex.Unlock()

	if day := time.Now().Day(); day != m.day {
		m.buckets = *new([8]int)
		m.day = day
	}
	m.buckets[i]++

	return fmt.Sprintf("%c%d", 65+i, m.buckets[i])
}
