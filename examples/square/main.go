package main

import (
	"github.com/floriankarydes/xybotsim/pkg/xybotsim"
)

func main() {
	simulator := xybotsim.NewSimulator(10)
	robot, _ := simulator.AddRobot("robot", 5, 5, 1)

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

	simulator.Show()
}
