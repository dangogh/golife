package grid

import (
	"sort"
	"testing"
)

var fixtures = [][][]Cell{
	{
		// given
		{{0, 0}, {3, 3}, {2, 2}},
		// expect
		{{0, 0}, {2, 2}, {3, 3}},
	},
	{
		// given
		{{0, 2}, {0, 3}, {0, 1}, {0, 2}},
		// expect
		{{0, 1}, {0, 2}, {0, 2}, {0, 3}},
	},
	{
		// given
		{{3, 2}, {0, 3}, {0, 1}},
		// expect
		{{0, 1}, {0, 3}, {3, 2}},
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
				t.Errorf("Given %v, expected %v, got %v", given, expect, got)
				break
			}
		}
	}
}
