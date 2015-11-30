package grid

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	g := Grid{}
	g.NextGen()

	if len(g) != 0 {
		t.Error("NextGen of empty grid should be empty")
	}
}

func TestOne(t *testing.T) {
	g := Grid{}
	g[Cell{0, 0}] = S{}

	if len(g) != 1 {
		t.Errorf("added one cell, but have %d", len(g))
	}

	g.NextGen()
	if len(g) != 0 {
		t.Errorf("NextGen of one cell should be empty, not %d", len(g))
	}

}

func gridSame(g1, g2 Grid) bool {
	if len(g1) != len(g2) {
		return false
	}
	for k1, v1 := range g1 {
		v2, ok := g2[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func TestBlinker(t *testing.T) {
	// 2 stages to blinker
	blinker := []Grid{
		Grid{
			Cell{1, 1}: S{},
			Cell{2, 1}: S{},
			Cell{3, 1}: S{},
		},
		Grid{
			Cell{1, 1}: S{},
			Cell{1, 2}: S{},
			Cell{1, 3}: S{},
		},
	}

	stage := 0
	g := blinker[stage]
	for ii := 0; ii < 4; ii++ {
		g.NextGen()
		stage = 1 - stage
		if !gridSame(g, blinker[stage]) {
			t.Errorf("NextGen should be %v, not %v", blinker[stage], g)
		}
	}
}
