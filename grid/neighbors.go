package grid

import (
	_ "fmt"
)

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
			nc := Cell{c.X + n.X, c.Y + n.Y}
			ncount[nc]++
		}
	}
	return ncount
}
