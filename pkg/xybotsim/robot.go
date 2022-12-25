package xybotsim

import (
	"errors"
	"time"
)

type Robot struct {
	currentCell  cell
	stepDuration time.Duration
	commandQueue chan Direction
	stopSignal   chan bool
}

func newRobot(x uint, y uint, velocity float32) (r Robot) {
	r = Robot{
		currentCell:  cell{x, y},
		stepDuration: time.Duration(float32(time.Second) / velocity),
		commandQueue: make(chan Direction, 100),
		stopSignal:   make(chan bool, 1),
	}
	return r
}

func (r *Robot) GetPosition() (x uint, y uint) {
	return r.currentCell.x, r.currentCell.y
}

func (r *Robot) EnqueueCommand(d Direction) error {
	if len(r.commandQueue) == cap(r.commandQueue) {
		return errors.New("queue is full")
	}
	r.commandQueue <- d
	return nil
}
