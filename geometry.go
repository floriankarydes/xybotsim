package xybotsim

import (
	"errors"
	"math"
)

type direction uint8

// Cardinal directions.
const (
	North direction = iota + 1
	East
	South
	West
)

// Cell on 2D grid.
type cell struct {
	x uint
	y uint
}

// Get next cell along a given direction.
// In case of overflow/underflow, return same cell.
func (c *cell) getAdjacent(d direction) (cell, error) {
	switch d {
	case North:
		if c.y == 0 {
			return *c, errors.New("underflow prevented")
		}
		return cell{c.x, c.y - 1}, nil
	case East:
		if c.x == math.MaxUint {
			return *c, errors.New("overflow prevented")
		}
		return cell{c.x + 1, c.y}, nil
	case South:
		if c.y == math.MaxUint {
			return *c, errors.New("overflow prevented")
		}
		return cell{c.x, c.y + 1}, nil
	case West:
		if c.x == 0 {
			return *c, errors.New("underflow prevented")
		}
		return cell{c.x - 1, c.y}, nil
	}
	return *c, nil
}
