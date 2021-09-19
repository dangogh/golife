package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	g := NewGrid()

	assert.Equal(t, g.Len(), 0)
}

func TestInsertAndRemove(t *testing.T) {
	g := NewGrid(Cell{1, 10})

	g.Insert(Cell{0, 0})

	assert.True(t, g.Has(Cell{0, 0}), "g.Has")
	assert.True(t, g.Has(Cell{1, 10}), "g.Has")
	assert.False(t, g.Has(Cell{0, 1}), "g.Has")
	assert.Equal(t, g.Len(), 2)

	g.Remove(Cell{1, 10})

	assert.True(t, g.Has(Cell{0, 0}), "g.Has")
	assert.False(t, g.Has(Cell{1, 10}), "g.Has")
	assert.False(t, g.Has(Cell{0, 1}), "g.Has")
	assert.Equal(t, g.Len(), 1)

}
