package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/floriankarydes/xybotsim/pkg/xybotsim"
)

func main() {
	var worldSize int = 50
	var robotCount int = 100
	var maxVelocity float32 = 5
	var maxPeriod float32 = 5

	rand.Seed(time.Now().UnixNano())

	simulator := xybotsim.NewSimulator(uint(worldSize))

	var robotIdList []string
	for i := 0; i < robotCount; i++ {

		robotId := fmt.Sprintf("robot_%v", i)

		go func() {
			time.Sleep(time.Duration(maxPeriod * float32(time.Second) * rand.Float32()))

			robotIdList = append(robotIdList, robotId)
			robot, _ := simulator.AddRobot(
				robotId,
				uint(rand.Intn(worldSize)),
				uint(rand.Intn(worldSize)),
				rand.Float32()*maxVelocity)

			go func() {
				for {
					time.Sleep(time.Duration(maxPeriod * float32(time.Second) * rand.Float32()))
					if robot == nil {
						return
					}
					robot.EnqueueCommand(xybotsim.Direction(1 + rand.Intn(3)))
				}
			}()
		}()

		go func() {
			time.Sleep(time.Duration(maxPeriod * float32(time.Second) * (3 + rand.Float32())))
			simulator.DeleteRobot(robotId)
		}()

	}

	simulator.Show()
}
