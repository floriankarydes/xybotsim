# XY Bot Simulator

Extremely simple simulation of a set of ["turtlebot"](https://www.turtlebot.com) style robots.

## Features

- 🤖 Test robot controllers written in [Go](https://go.dev) with software-in-the-loop simulations.
- 💯 Run large scale multi-agent scenarios with hundreds of robots.
- 💥 Simulate robot collisions with obstacles and other robots.
- 📘 Connect your code easily with the most simple [API](#getting-started).
- 📺 Visualize simulations in real-time with a minimal GUI.

But don't get your hopes too high, this will be dots moving on a grid 😅 If you are looking for a proper robotics simulator, you should checkout [Gazebo](https://gazebosim.org/home).

XY Bot Simulator's design is actually very minimalist :

- Simulated world is a 2D squared grid.
- Robot moves from cell to cell with a chosen velocity.
- Movement is discrete & limited to adjacent cells (no diagonal).

## Prerequisites

To run your simulations using XY Bot Simulator you will need Go version 1.14 or later, a C compiler and your system's development tools.

Using the standard Go tools you can install XY Bot Simulator's core library using:

```bash
go get github.com/floriankarydes/xybotsim
```

## Getting Started

Open a new file `my_simulation.go` and you're ready to setup your first simulation!

```go
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
```

And you can run that simply as:

```bash
go run my_simulation.go
```

It should look like this:

<img src="https://j.gifs.com/RlJn8q.gif" alt="First simulation preview" width="400" height="420">

## Going further

Please checkout the [examples](examples) available to learn how to run more complex simulation scenarios.
