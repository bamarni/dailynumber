package dailynumber

import (
	"fmt"
	"time"
)

type Generator struct {
	counter Counter
}

func New(counter Counter) *Generator {
	return &Generator{
		counter: counter,
	}
}

func (g *Generator) Generate() string {
	year, month, day := time.Now().In(time.UTC).Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	number := int(time.Since(midnight).Minutes()) + g.counter.Increment(midnight)

	return fmt.Sprintf("%c%d", 65+number%26, number/26)
}
