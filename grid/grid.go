package grid

type Cell struct {
	X, Y int
}

// S represents an empty struct -- just a cheap placeholder
type Grid struct {
	set map[Cell]struct{}
}

func (g *Grid) AddCell(c Cell) bool {
	found := g.ExistsCell(c)
	if !found {
		g.set[c] = struct{}{}
	}
	return !found
}

func (g *Grid) DeleteCell(c Cell) bool {
	found := g.ExistsCell(c)
	delete(g.set, c)
	return !found
}

func (g *Grid) ExistsCell(c Cell) bool {
	_, found := g.set[c]
	return found
}

func (g *Grid) AddCoords(x, y int) bool {
	return g.AddCell(Cell{x, y})
}

func (g *Grid) DeleteCoords(x, y int) bool {
	return g.DeleteCell(Cell{x, y})
}

func (g *Grid) ExistsCoords(x, y int) bool {
	return g.ExistsCell(Cell{x, y})
}

func (g Grid) NextGen() {
	nc := g.countNeighbors()
	// for each live cell, if not enough or too many neighbors, kill it
	for livecell := range g.set {
		n, ok := nc[livecell]
		if !ok || n < 2 || n > 3 {
			delete(g.set, livecell)
		}
		// remove from neighbor count map
		delete(nc, livecell)
	}
	// only non-live still in neighbor count -- for each with 3 neighbors emerges
	for deadcell, n := range nc {
		if n == 3 {
			g.DeleteCell(deadcell)
		}
	}
}
