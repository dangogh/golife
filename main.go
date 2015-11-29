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
			relNeighbors[Cell{x, y}] = struct{}{}
		}
	}
	fmt.Printf("neighborgrid is %#v\n", relNeighbors)
}

func main() {
	g := make(Grid, 5)

	c := Cell{5, 5}
	g[c] = struct{}{}

	fmt.Printf("grid is %#v\n", g)
}
