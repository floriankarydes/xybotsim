package xybotsim

import (
	"errors"
	"sync"
	"time"
)

type Simulator struct {
	worldSize      uint
	occupancyGrid  [][]bool
	occupancyMutex sync.RWMutex
	robotRegister  map[string]*Robot
	robotMutex     sync.RWMutex
}

func NewSimulator(size uint) (s Simulator) {
	s = Simulator{
		worldSize:     size,
		occupancyGrid: make([][]bool, size),
		robotRegister: make(map[string]*Robot),
	}
	for i := range s.occupancyGrid {
		s.occupancyGrid[i] = make([]bool, size)
	}
	return
}

func (s *Simulator) AddRobot(id string, x uint, y uint, velocity float32) (*Robot, error) {
	s.robotMutex.Lock()
	defer s.robotMutex.Unlock()
	s.occupancyMutex.Lock()
	defer s.occupancyMutex.Unlock()
	if _, idExist := s.robotRegister[id]; idExist {
		return nil, errors.New("ID already exists")
	}
	if velocity < 0 {
		return nil, errors.New("velocity cannot be negative")
	}
	if s.occupancyGrid[x][y] {
		return nil, errors.New("cell is already occupied")
	}
	s.occupancyGrid[x][y] = true
	r := newRobot(x, y, velocity)
	s.robotRegister[id] = &r
	s.startRobot(s.robotRegister[id])
	return s.robotRegister[id], nil
}

func (s *Simulator) DeleteRobot(id string) error {
	s.robotMutex.Lock()
	defer s.robotMutex.Unlock()
	if _, idExist := s.robotRegister[id]; !idExist {
		return errors.New("ID does not exist")
	}
	s.stopRobot(s.robotRegister[id])
	position := s.robotRegister[id].currentCell
	delete(s.robotRegister, id)
	s.occupancyGrid[position.x][position.y] = false
	return nil
}

func (s *Simulator) ListRobots() (robotArray []*Robot) {
	s.robotMutex.RLock()
	defer s.robotMutex.RUnlock()
	for _, robotPtr := range s.robotRegister {
		robotArray = append(robotArray, robotPtr)
	}
	return robotArray
}

func (s *Simulator) occupyCell(c cell) error {
	s.occupancyMutex.Lock()
	defer s.occupancyMutex.Unlock()
	if c.x >= s.worldSize || c.y >= s.worldSize {
		return errors.New("out of world")
	}
	if s.occupancyGrid[c.x][c.y] {
		return errors.New("already occupied")
	}
	s.occupancyGrid[c.x][c.y] = true
	return nil
}

func (s *Simulator) freeCell(c cell) error {
	s.occupancyMutex.Lock()
	defer s.occupancyMutex.Unlock()
	if c.x > s.worldSize || c.y > s.worldSize {
		return errors.New("out of world")
	}
	if !s.occupancyGrid[c.x][c.y] {
		return errors.New("not occupied")
	}
	s.occupancyGrid[c.x][c.y] = false
	return nil
}

func (s *Simulator) startRobot(r *Robot) {
	for len(r.stopSignal) > 0 {
		<-r.stopSignal
	}
	for len(r.commandQueue) > 0 {
		<-r.commandQueue
	}
	go func() {
		for {
			select {
			case <-r.stopSignal:
				return
			case d := <-r.commandQueue:
				targetCell, _ := r.currentCell.getAdjacent(d)
				if s.occupyCell(targetCell) == nil {
					select {
					case <-r.stopSignal:
						s.freeCell(targetCell)
						return
					case <-time.After(r.stepDuration):
					}
					s.freeCell(r.currentCell)
					r.currentCell = targetCell
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

func (s *Simulator) stopRobot(r *Robot) {
	r.stopSignal <- true
}
