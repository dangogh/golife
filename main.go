package main

import (
	"fmt"
	"github.com/dangogh/golife/grid"
)

func main() {
	g := &grid.Grid{
		grid.Cell{4, 4}: struct{}{},
		grid.Cell{4, 5}: struct{}{},
		grid.Cell{4, 6}: struct{}{},
		grid.Cell{5, 4}: struct{}{},
		grid.Cell{5, 6}: struct{}{},
		grid.Cell{6, 4}: struct{}{},
		grid.Cell{6, 5}: struct{}{},
		grid.Cell{6, 6}: struct{}{},
	}

	g.NextGen()

	fmt.Printf("grid is %s\n", g)
}
