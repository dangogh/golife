package grid

import (
	"fmt"
	"sort"
)

type extrema struct {
	minX, minY int
	maxX, maxY int
}

func (c Cell) String() string {
	return fmt.Sprintf("[%3d,%3d]", c.X, c.Y)
}

func (g Grid) bounds() extrema {
	if len(g) == 0 {
		return extrema{}
	}
	c := g[0]
	e := extrema{c.X, c.Y, c.X, c.Y}
	for c := range g {
		if c.X < e.minX {
			e.minX = c.X
		}
		if c.X > e.maxX {
			e.maxX = c.X
		}
	}
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
			for i := prev.Y; i < c.Y; i++ {
				s += "\n"
			}
		}
		prev = &c
		s += c.String()
	}
	return s
}
