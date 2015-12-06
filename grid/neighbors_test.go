package grid

import (
	//"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	if len(getRelNeighbors()) != 8 {
		t.Errorf("should be 8 relative neighbor entries, not %d", len(getRelNeighbors()))
	}

	for c, _ := range getRelNeighbors() {
		if c.X < -1 || c.X > 1 || c.Y < -1 || c.Y > 1 {
			t.Errorf("relative neighbor should be within -1, 1 not %v\n", c)
		}
	}
}

func TestNoNeighbors(t *testing.T) {
	g := Grid{}
	c := Cell{4, 3}
	g[c] = S{}

	nc := g.countNeighbors()
	if nc[c] != 0 {
		t.Errorf("lone cell should have 0 neighbors, not %d", nc[c])
	}
	// each relative neighbor should have exactly one
	for r, _ := range getRelNeighbors() {
		neighbor := Cell{c.X + r.X, c.Y + r.Y}
		if nc[neighbor] != 1 {
			t.Errorf("neighbors of lone cell should have 1 neighbor, not %d", nc[neighbor])
		}
	}
}

func TestOneNeighbor(t *testing.T) {
	g := Grid{}
	c0 := Cell{4, 3}
	c1 := Cell{3, 3}
	g[c0] = S{}
	g[c1] = S{}

	nc := g.countNeighbors()
	if nc[c0] != 1 {
		t.Errorf("each of 2 adjacent cells should have 1 neighbors, not %d", nc[c0])
	}
	if nc[c1] != 1 {
		t.Errorf("each of 2 adjacent cells should have 1 neighbors, not %d", nc[c1])
	}
	// each relative neighbor should have exactly one
	for r, _ := range getRelNeighbors() {
		neighbor := Cell{c0.X + r.X, c0.Y + r.Y}
		//fmt.Printf("neighbor of %v is %v\n", c0, neighbor)
		if nc[neighbor] != 1 && nc[neighbor] != 2 {
			t.Errorf("neighbors of adjacent cells should have 1 or 2 neighbors, not %d", nc[neighbor])
		}
	}
}
