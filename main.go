package main

import (
	"fmt"
	"github.com/dangogh/golife/grid"
)

func main() {
	g := new(grid.Grid)

	c := grid.Cell{5, 5}
	g.AddCell(c)

	g.NextGen()

	fmt.Printf("grid is %#v\n", g)
}
