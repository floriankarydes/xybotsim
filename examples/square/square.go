package main

import (
	"github.com/floriankarydes/xybotsim"
)

// Extremely simple robot controller simulating moving a robot in a 1x1 square pattern.
func main() {

	// Create a new simulator with a 10x10 world grid.
	simulator := xybotsim.NewSimulator(10)

	// Spawn a new robot in simulator world on cell (5,5).
	// Set robot velocity to 1 cell per second
	robot, _ := simulator.AddRobot("robot", 5, 5, 1)

	// Command the robot in a 1x1 square pattern (async)
	go func() {
		for i := 0; i < 4*3; i++ {
			switch i % 4 {
			case 0:
				robot.EnqueueCommand(xybotsim.North)
			case 1:
				robot.EnqueueCommand(xybotsim.East)
			case 2:
				robot.EnqueueCommand(xybotsim.South)
			case 3:
				robot.EnqueueCommand(xybotsim.West)
			}
		}
	}()

	// Display the simulated world on screen
	simulator.Show()
}
