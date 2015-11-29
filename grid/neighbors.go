package grid

import (
//"fmt"
)

type neighborCount map[Cell]int

var relNeighbors Grid = Grid{
	Cell{-1, -1}: S{}, Cell{-1, 0}: S{}, Cell{-1, 1}: S{},
	Cell{0, -1}: S{}, Cell{0, 1}: S{},
	Cell{1, -1}: S{}, Cell{1, 0}: S{}, Cell{1, 1}: S{},
}

func _init() {
	r := []int{-1, 0, 1}
	relNeighbors := make(Grid, len(r)^2-1)
	for _, x := range r {
		for _, y := range r {
			if x != 0 || y != 0 {
				relNeighbors[Cell{x, y}] = S{}
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
	ncount := make(neighborCount, len(g)*8)
	for c, _ := range g {
		for n, _ := range relNeighbors {
			nc := Cell{c.X + n.X, c.Y + n.Y}
			ncount[nc]++
		}
	}
	return ncount
}
