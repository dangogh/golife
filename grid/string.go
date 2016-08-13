package grid

import (
	"fmt"
	"sort"
)

func (c Cell) String() string {
	return fmt.Sprintf("[%3d,%3d]", c.X, c.Y)
}

func (g Grid) String() string {
	var a []Cell
	for c, _ := range g {
		a = append(a, c)
	}
	sort.Sort(byCoord(a))
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
