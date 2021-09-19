package grid

import (
	"log"
	"math"
)

func (g Grid) bounds() [][]int64 {
	var minX, maxX, minY, maxY int64

	minX, maxX = math.MaxInt64, math.MinInt64
	minY, maxY = math.MaxInt64, math.MinInt64

	g.Do(func(i interface{}) {
		c, ok := i.(Cell)
		if !ok {
			log.Panicf("Invalid type in set %t", i)
		}

		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	})

	return [][]int64{{minX, minY}, {maxX, maxY}}
}

func (g Grid) String() string {
	var s string

	b := g.bounds()
	maxX, maxY := 2*b[1][0], 2*b[1][1]

	// print rows top-to-bottom
	for row := int64(0); row <= maxY; row++ {
		for col := int64(0); col <= maxX; col++ {
			ch := " "
			if g.Has(Cell{X: col, Y: row}) {
				ch = "*"
			}

			s += " " + ch
		}

		s += "\n"
	}

	return s
}
