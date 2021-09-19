package grid

func (g Grid) countNeighbors() map[Cell]int {
	// for each cell, increment neighbor count of each of its neighbors
	ncount := make(map[Cell]int, g.Len()*8)

	g.Do(func(c interface{}) {
		c.(Cell).applyToNeighbors(func(nc Cell) {
			ncount[nc]++
		})
	})

	return ncount
}

func (c Cell) applyToNeighbors(fn func(Cell)) {
	for x := c.X - 1; x <= c.X+1; x++ {
		for y := c.Y - 1; y <= c.Y+1; y++ {
			nc := Cell{X: x, Y: y}
			if c != nc {
				fn(nc)
			}
		}
	}
}
