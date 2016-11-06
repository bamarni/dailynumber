package dailynumber

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

type Generator struct {
	counter int
	number  int
	mutex   sync.Mutex
}

func (g *Generator) Generate() string {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	year, month, day := time.Now().In(time.UTC).Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	currentNumber := int(time.Since(midnight).Minutes())
	if currentNumber < g.number {
		g.counter = 0
	}
	g.number = currentNumber
	g.counter++

	number := int64(g.number + g.counter)

	return strings.ToUpper(strconv.FormatInt(number, 18))
}
