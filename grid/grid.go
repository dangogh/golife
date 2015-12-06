package grid

// S represents an empty struct -- just a cheap placeholder
type S struct{}

type Cell struct {
	X, Y int
}

type Grid map[Cell]S

func (c Cell) OffsetBy(offset Cell) (ret Cell) {
	return Cell{c.X+offset.X, c.Y+offset.Y}
}

func (g Grid) NextGen() {
	nc := g.countNeighbors()
	// for each live cell, if not enough or too many neighbors, kill it
	for livecell := range g {
		n, ok := nc[livecell]
		if !ok || n < 2 || n > 3 {
			delete(g, livecell)
		}
		// remove from neighbor count map
		delete(nc, livecell)
	}
	// only non-live still in neighbor count -- for each with 3 neighbors emerges
	for deadcell, n := range nc {
		if n == 3 {
			g[deadcell] = S{}
		}
	}
}
