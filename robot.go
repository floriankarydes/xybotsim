package xybotsim

import (
	"errors"
	"sync"
	"time"
)

// Get coordinates of the cell occupied by robot.
func (r *Robot) GetPosition() (x uint, y uint) {
	c := r.getPosition()
	return c.x, c.y
}

// Send a new command (i.e unit step in a given cardinal direction) to robot.
// Commands are processed one after the other at a rate constrained by the robot velocity.
//
// See AddRobot.
func (r *Robot) EnqueueCommand(d direction) error {
	if len(r.commandQueue) == cap(r.commandQueue) {
		return errors.New("queue is full")
	}
	r.commandQueue <- d
	return nil
}

// Robot that lives on the world grid of a Simulator.
//
// See AddRobot.
type Robot struct {
	position      cell           // Position of the robot on the world grid
	stepDuration  time.Duration  // Duration required to complete one unit movement.
	commandQueue  chan direction // Queue for incoming command input.
	stopSignal    chan bool      // Signal to stop the robot's simulation.
	positionMutex sync.RWMutex   // Ensure thread-safe access of robot's position.
}

// Create a new robot instance.
//
// See AddRobot.
func newRobot(x uint, y uint, velocity float32) *Robot {
	r := Robot{
		position:     cell{x, y},
		stepDuration: time.Duration(float32(time.Second) / velocity),
		commandQueue: make(chan direction, 100),
		stopSignal:   make(chan bool, 1),
	}
	return &r
}

// Set robot's position on world grid.
func (r *Robot) setPosition(c cell) {
	r.positionMutex.Lock()
	defer r.positionMutex.Unlock()
	r.position = c
}

// Get robot's position on world grid.
func (r *Robot) getPosition() (c cell) {
	r.positionMutex.RLock()
	defer r.positionMutex.RUnlock()
	return r.position
}
