package xybotsim

import (
	"math"
	"testing"
)

func TestAdjacentNoWrap(t *testing.T) {
	limitCells := [2]cell{{0, 0}, {math.MaxUint, math.MaxUint}}
	allDirections := [4]direction{North, East, South, West}
	for _, d := range allDirections {
		for _, sourceCell := range limitCells {
			targetCell, err := sourceCell.getAdjacent(d)
			if distance(sourceCell, targetCell) != 1 && err == nil {
				t.Fatalf("%#v.getAdjacent(%#v) fail to return adjacent cell without throwing an error.", sourceCell, d)
			}
		}
	}

}

func TestAdjacentCell(t *testing.T) {
	argCell := cell{5, 5}
	adjacentCells := map[direction]cell{
		North: {5, 4},
		East:  {6, 5},
		South: {5, 6},
		West:  {4, 5},
	}
	for d, adjacentCell := range adjacentCells {
		retCell, err := argCell.getAdjacent(d)
		if adjacentCell != retCell || err != nil {
			t.Fatalf("%#v.getAdjacent(%#v) = %#v, %v, want match for %#v, nil.", argCell, d, retCell, err, adjacentCell)
		}
	}
}

func distance(a cell, b cell) uint {
	var d uint = 0

	if a.x < b.x {
		d += b.x - a.x
	} else {
		d += a.x - b.x
	}

	if a.y < b.y {
		d += b.y - a.y
	} else {
		d += a.y - b.y
	}

	return d
}
