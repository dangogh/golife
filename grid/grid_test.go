package grid

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	g := Grid{}
	g.NextGen()

	if len(g.set) != 0 {
		t.Error("NextGen of empty grid should be empty")
	}
}

func TestOne(t *testing.T) {
	g := Grid{}
	g.AddCoords(0, 0)

	if len(g.set) != 1 {
		t.Errorf("added one cell, but have %d", len(g.set))
	}

	g.NextGen()
	if len(g.set) != 0 {
		t.Errorf("NextGen of one cell should be empty, not %d", len(g.set))
	}

}

func gridSame(g1, g2 Grid) bool {
	if len(g1.set) != len(g2.set) {
		return false
	}
	for k1, v1 := range g1.set {
		v2, ok := g2.set[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func TestBlinker(t *testing.T) {
	// 2 stages to blinker
	blinker_stages := [][]int{{2, 1}, {2, 2}, {2, 3}}
	blinker := make([]Grid, 3)
	blinker[0] = Grid{}
	blinker[0].AddCoords(2, 1)
	blinker.AddCoords(2, 2)
	blinker.AddCoords(2, 3)

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
