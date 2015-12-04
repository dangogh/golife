package grid

import (
//"fmt"
)

type neighborCount map[Cell]int

var relNeighbors Grid

func init() {
	r := []int{-1, 0, 1}
	relNeighbors := new(Grid)
	for _, x := range r {
		for _, y := range r {
			if x != 0 || y != 0 {
				relNeighbors.AddCoords(x, y)
			}
		}
	}
}

func getRelNeighbors() Grid {
	//fmt.Printf("relneighbors: %v\n", relNeighbors)
	return relNeighbors
}

func (g Grid) countNeighbors() neighborCount {
	// for each cell, increment neighbor count of each of its neighbors
	ncount := make(neighborCount, len(g.set)*8)
	for c, _ := range g.set {
		for n, _ := range relNeighbors.set {
			nc := Cell{c.X + n.X, c.Y + n.Y}
			ncount[nc]++
		}
	}
	return ncount
}
