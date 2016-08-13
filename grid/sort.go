package grid

import ()

type byCoord []Cell

func (b byCoord) Len() int {
	return len([]Cell(b))
}

func (b byCoord) Less(i, j int) bool {
	switch {
	case b[i].X < b[j].X:
		return true
	case b[i].X > b[j].X:
		return false
	case b[i].Y < b[j].Y:
		return true
	case b[i].Y > b[j].Y:
		return false
	default:
		return true
	}
}

func (b byCoord) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
