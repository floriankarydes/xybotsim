package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/floriankarydes/xybotsim"
)

// Multi-agent simulations with robot spawned randomly in the map.
func main() {
	var worldSize int = 50
	var robotCount int = 100
	var maxVelocity float32 = 5
	var maxPeriod float32 = 5

	rand.Seed(time.Now().UnixNano())

	simulator := xybotsim.NewSimulator(uint(worldSize))

	// Manage several robots lifecycle (i.e. spawn, move, remove) asynchronously.
	var robotIdList []string
	for i := 0; i < robotCount; i++ {

		robotId := fmt.Sprintf("robot_%v", i)

		// Add new robots at random time, random place & random velocity
		go func() {
			time.Sleep(time.Duration(maxPeriod * float32(time.Second) * rand.Float32()))

			robotIdList = append(robotIdList, robotId)
			robot, _ := simulator.AddRobot(
				robotId,
				uint(rand.Intn(worldSize)),
				uint(rand.Intn(worldSize)),
				rand.Float32()*maxVelocity)

			// Once spawned, feed the robot with random command at irregular rate.
			go func() {
				for {
					time.Sleep(time.Duration(maxPeriod * float32(time.Second) * rand.Float32()))
					if robot == nil {
						return
					}
					switch rand.Intn(3) {
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
		}()

		// After a while, remove the robot from the world.
		go func() {
			time.Sleep(time.Duration(maxPeriod * float32(time.Second) * (3 + rand.Float32())))
			simulator.DeleteRobot(robotId)
		}()

	}

	// Display the simulated world on screen
	simulator.Show()
}
