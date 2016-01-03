package grid

import (
	"fmt"
	"sort"
)

type ByCoord []Cell

func (b ByCoord) Len() int {
	return len([]Cell(b))
}

func (b ByCoord) Less(i, j int) bool {
	if b[i].X > b[j].X || b[i].Y > b[j].Y {
		return false
	}
	return true
}

func (b ByCoord) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (c Cell) String() string {
	return fmt.Sprintf("[%3d,%3d]", c.X, c.Y)
}

func (g Grid) String() string {
	a := make([]Cell, 0, len(g))
	for c, _ := range g {
		a = append(a, c)
	}
	sort.Sort(ByCoord(a))
	var prev *Cell = nil
	s := ""
	for _, c := range a {
		if prev != nil {
			if prev.Y != c.Y {
				s += "\n"
			} else {
				s += " "
			}
		}
		prev = &c
		s += c.String()
	}
	return s
}
