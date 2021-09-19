package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/dangogh/golife/grid"
)

func clear() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	flag.Parse()

	fh, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	cells, err := readCells(fh)
	if err != nil {
		log.Fatal(err)
	}

	g := grid.NewGrid(cells...)

	clear()

	fmt.Println(g.String())

	for g.NextGen() {
		time.Sleep(time.Second)
		clear()

		fmt.Println(g.String())
	}

}

func readCells(r io.Reader) ([]grid.Cell, error) {
	var cells []grid.Cell

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var x, y int64

		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		if err != nil {
			return cells, err
		}

		cells = append(cells, grid.Cell{X: x, Y: y})
	}

	return cells, nil
}
