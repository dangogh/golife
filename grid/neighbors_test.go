package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNeighbors(t *testing.T) {
	g := NewGrid()

	g.Insert(Cell{1, 1})

	counts := g.countNeighbors()

	assert.Equal(t, len(counts), 8)
	assert.Contains(t, counts, Cell{0, 0})
	assert.Contains(t, counts, Cell{2, 0})
}
