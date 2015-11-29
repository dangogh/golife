package main

import (
	"fmt"
)

type Cell struct {
	x, y int
}

type Grid map[Cell]struct{}
type neighborCount map[Cell]int

var relNeighbors Grid

func init() {
	r := []int{-1, 0, 1}
	relNeighbors := make(Grid, len(r)^2-1)
	for _, x := range r {
		for _, y := range r {
			if x != 0 || y != 0 {
				relNeighbors[Cell{x, y}] = struct{}{}
			}
		}
	}
}

func (g Grid) countNeighbors() neighborCount {
	// for each cell, increment neighbor count of each of its neighbors
	ncount := make(neighborCount, len(g)*8)
	for c, _ := range g {
		for n, _ := range relNeighbors {
			nc := Cell{c.x + n.x, c.y + n.y}
			ncount[nc]++
		}
	}
	return ncount
}

func (g Grid) nextGen() {
	nc := g.countNeighbors()
	// for each live cell, if not enough or too many neighbors, kill it
	for c := range g {
		if nc[c] < 2 || nc[c] > 3 {
			delete(g, c)
		}
		delete(nc, c)
	}
	// only non-live still in neighbor count -- for each with 3 neighbors emerges
	for deadcell, n := range nc {
		if n == 3 {
			g[deadcell] = struct{}{}
		}
	}
}

func main() {
	g := make(Grid, 5)

	c := Cell{5, 5}

	g[c] = struct{}{}
	g.nextGen()

	fmt.Printf("grid is %#v\n", g)
}
