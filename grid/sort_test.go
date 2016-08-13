package grid

import (
	"sort"
	"testing"
)

var fixtures = [][][]Cell{
	[][]Cell{
		// given
		[]Cell{Cell{0, 0}, Cell{3, 3}, Cell{2, 2}},
		// expect
		[]Cell{Cell{0, 0}, Cell{2, 2}, Cell{3, 3}},
	},
	[][]Cell{
		// given
		[]Cell{Cell{0, 2}, Cell{0, 3}, Cell{0, 1}},
		// expect
		[]Cell{Cell{0, 1}, Cell{0, 2}, Cell{0, 3}},
	},
	[][]Cell{

		// given
		[]Cell{Cell{3, 2}, Cell{0, 3}, Cell{0, 1}},
		// expect
		[]Cell{Cell{0, 1}, Cell{0, 3}, Cell{3, 2}},
	},
}

func TestSort(t *testing.T) {
	for _, f := range fixtures {
		given, expect := f[0], f[1]
		got := make([]Cell, len(f[0]))
		copy(got, f[0])
		sort.Sort(byCoord(got))

		for i := range expect {
			if got[i] != expect[i] {
				t.Errorf("Given %s, expected %s, got %s", given, expect, got)
				break
			}
		}
	}
}
