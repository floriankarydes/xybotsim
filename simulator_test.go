package xybotsim

import (
	"testing"
	"time"
)

func TestRobotCommand(t *testing.T) {
	var size uint = 1000
	simulator := NewSimulator(size)

	id := "robotId"
	velocity := float32(100)
	srcPos := cell{0, 0}
	targetPos := cell{1, 1}
	robot, _ := simulator.AddRobot(id, srcPos.x, srcPos.y, velocity)
	robot.EnqueueCommand(East)
	robot.EnqueueCommand(South)
	time.Sleep(3 * robot.stepDuration)
	if robot.getPosition() != targetPos {
		t.Fatalf("Robot do not respond to command inputs.")
	}
}

func TestRobotVelocity(t *testing.T) {
	var size uint = 1000
	simulator := NewSimulator(size)

	var velocity float32 = 100
	srcPos := cell{0, 0}
	targetPos := cell{1, 1}
	robot, _ := simulator.AddRobot("robotId", srcPos.x, srcPos.y, velocity)
	robot.EnqueueCommand(East)
	robot.EnqueueCommand(South)
	time.Sleep(1 * robot.stepDuration)
	if robot.getPosition() == targetPos {
		t.Fatalf("Robot teleport on command input (i.e. velocity is not simulated).")
	}
}

func TestRobotCollision(t *testing.T) {
	var size uint = 10
	simulator := NewSimulator(size)

	var velocity float32 = float32(100 * size)
	minPos := cell{0, 0}
	maxPos := cell{size - 1, 0}

	robot1, _ := simulator.AddRobot("robot_01", minPos.x, minPos.y, velocity)
	for i := 0; i < 100; i++ {
		robot1.EnqueueCommand(East)
	}

	robot2, _ := simulator.AddRobot("robot_02", maxPos.x, maxPos.y, velocity)
	for i := 0; i < 100; i++ {
		robot2.EnqueueCommand(West)
	}

	time.Sleep(time.Duration(2 * size * uint(robot1.stepDuration)))
	if robot1.getPosition() == maxPos || robot2.getPosition() == minPos {
		t.Fatalf("Robot do not collide and go through each others.")
	}
}

func TestWorldBound(t *testing.T) {
	var size uint = 10
	simulator := NewSimulator(size)

	var velocity float32 = float32(100 * size)
	minPos := cell{0, 0}
	maxPos := cell{size - 1, 0}
	robot, _ := simulator.AddRobot("robotId", minPos.x, minPos.y, velocity)
	for i := 0; i < 20; i++ {
		robot.EnqueueCommand(East)
	}

	time.Sleep(time.Duration(3 * size * uint(robot.stepDuration)))
	curPos := robot.getPosition()
	if curPos != maxPos {
		t.Fatalf("Robot should block on world bound %v but get %v.", maxPos, curPos)
	}
}
