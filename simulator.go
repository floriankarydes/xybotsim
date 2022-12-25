// Run extremely simple simulation of a set of "turtlebot" style robots.
package xybotsim

import (
	"errors"
	"sync"
	"time"
)

// Create a new simulator instance with a given world size.
// World is a square grid of dimension `size` x `size`.
func NewSimulator(size uint) (s Simulator) {
	s = Simulator{
		worldSize:     size,
		worldGrid:     make([][]bool, size),
		robotRegister: make(map[string]*Robot),
	}
	for i := range s.worldGrid {
		s.worldGrid[i] = make([]bool, size)
	}
	return
}

// Spawn a new robot on a give cell in the simulator's world grid.
// New robots are registered under a unique string ID on creation.
// Robots have constant velocity expressed in cell per second.
func (s *Simulator) AddRobot(id string, x uint, y uint, velocity float32) (*Robot, error) {
	s.robotMutex.Lock()
	defer s.robotMutex.Unlock()
	s.worldMutex.Lock()
	defer s.worldMutex.Unlock()
	if _, idExist := s.robotRegister[id]; idExist {
		return nil, errors.New("ID already exists")
	}
	if velocity < 0 {
		return nil, errors.New("velocity cannot be negative")
	}
	if s.worldGrid[x][y] {
		return nil, errors.New("cell is already occupied")
	}
	s.worldGrid[x][y] = true
	s.robotRegister[id] = newRobot(x, y, velocity)
	s.startRobot(s.robotRegister[id])
	return s.robotRegister[id], nil
}

// Remove an existing robot from the simulator's world.
func (s *Simulator) DeleteRobot(id string) error {
	s.robotMutex.Lock()
	defer s.robotMutex.Unlock()
	if _, idExist := s.robotRegister[id]; !idExist {
		return errors.New("ID does not exist")
	}
	s.stopRobot(s.robotRegister[id])
	position := s.robotRegister[id].getPosition()
	delete(s.robotRegister, id)
	s.worldGrid[position.x][position.y] = false
	return nil
}

// List all robots living in the simulator's world.
func (s *Simulator) ListRobots() (robotArray []*Robot) {
	s.robotMutex.RLock()
	defer s.robotMutex.RUnlock()
	for _, robotPtr := range s.robotRegister {
		robotArray = append(robotArray, robotPtr)
	}
	return robotArray
}

// Simulator for multi-robots scenario on 2D grid.
type Simulator struct {
	worldSize     uint              // Size (width & height) of the square world grid.
	worldGrid     [][]bool          // Occupancy map for the world. True if cell is occupied.
	worldMutex    sync.RWMutex      // Ensure thread-safe access of the occupancy map.
	robotRegister map[string]*Robot // Register all robot living on the world grid.
	robotMutex    sync.RWMutex      // Ensure thread-safe access to the robots' register.
}

// Set cell status to occupied if free, send error if not.
func (s *Simulator) occupyCell(c cell) error {
	s.worldMutex.Lock()
	defer s.worldMutex.Unlock()
	if c.x >= s.worldSize || c.y >= s.worldSize {
		return errors.New("out of world")
	}
	if s.worldGrid[c.x][c.y] {
		return errors.New("already occupied")
	}
	s.worldGrid[c.x][c.y] = true
	return nil
}

// Set cell status to free if occupied, send error if not.
func (s *Simulator) freeCell(c cell) error {
	s.worldMutex.Lock()
	defer s.worldMutex.Unlock()
	if c.x > s.worldSize || c.y > s.worldSize {
		return errors.New("out of world")
	}
	if !s.worldGrid[c.x][c.y] {
		return errors.New("not occupied")
	}
	s.worldGrid[c.x][c.y] = false
	return nil
}

// Start robot simulation thread.
// Handle incoming command, move robot according to velocity & check for collisions.
// When colliding, the "moving" robot wheel slides on ground and robot stay at same position.
//
// See AddRobot
func (s *Simulator) startRobot(r *Robot) {

	// Flush all previous signal or commands.
	for len(r.stopSignal) > 0 {
		<-r.stopSignal
	}
	for len(r.commandQueue) > 0 {
		<-r.commandQueue
	}

	go func() {
		for {
			// Wait for new command.
			select {
			case <-r.stopSignal:
				return
			case d := <-r.commandQueue:
				// New incoming command is available in queue.
				currentCell := r.getPosition()
				targetCell, _ := currentCell.getAdjacent(d)
				// Try to move according to command. Slip if running against another robot.
				if s.occupyCell(targetCell) == nil {
					select {
					case <-r.stopSignal:
						s.freeCell(targetCell)
						return
					case <-time.After(r.stepDuration):
					}
					s.freeCell(currentCell)
					r.setPosition(targetCell)
				} else {
					select {
					case <-r.stopSignal:
						return
					case <-time.After(r.stepDuration):
					}
				}
			}
		}
	}()
}

// Stop robot simulation thread.
func (s *Simulator) stopRobot(r *Robot) {
	r.stopSignal <- true
}
