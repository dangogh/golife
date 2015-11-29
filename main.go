package main

import (
	"fmt"
	"github.com/dangogh/golife/grid"
)

func main() {
	g := make(grid.Grid, 5)

	c := grid.Cell{5, 5}

	g[c] = struct{}{}
	g.NextGen()

	fmt.Printf("grid is %#v\n", g)
}
