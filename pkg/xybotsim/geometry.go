package xybotsim

import (
	"errors"
	"math"
)

type Direction uint8

const (
	North Direction = iota + 1
	East
	South
	West
)

type cell struct {
	x uint
	y uint
}

func (c *cell) getAdjacent(d Direction) (cell, error) {
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
