package grid

import "github.com/golang-collections/collections/set"

// S represents an empty struct -- just a cheap placeholder
type S struct{}

// Cell is a single live cell in the grid.
type Cell struct {
	X, Y int64
}

// Grid is the collection of live cells
type Grid struct {
	set *set.Set
}

func NewGrid(cells ...Cell) *Grid {
	ii := make([]interface{}, len(cells))

	for i, c := range cells {
		ii[i] = c
	}

	return &Grid{set: set.New(ii...)}
}

func (g *Grid) Insert(c Cell) {
	g.set.Insert(c)
}

func (g *Grid) Remove(c Cell) {
	g.set.Remove(c)
}

func (g *Grid) Do(fn func(interface{})) {
	g.set.Do(fn)
}

func (g Grid) Len() int {
	return g.set.Len()
}

func (g Grid) Has(c Cell) bool {
	return g.set.Has(c)
}
