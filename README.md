# XY Bot Simulator

[![Platform Tests](https://github.com/floriankarydes/xybotsim/actions/workflows/platform_tests.yml/badge.svg?branch=main)](https://github.com/floriankarydes/xybotsim/actions/workflows/platform_tests.yml) [![GitHub tag (Latest by date)](https://img.shields.io/github/v/tag/floriankarydes/xybotsim)](https://github.com/floriankarydes/xybotsim/releases) ![GitHub all releases](https://img.shields.io/github/downloads/floriankarydes/xybotsim/total)

Run extremely simple simulations of ["turtlebot"](https://www.turtlebot.com) style robots.

## Features

- 🤖 Test robot controllers written in [Go](https://go.dev).
- 💯 Run hundreds of robots simultaneously.
- 💥 Simulate robot collisions.
- 📘 Connect your code with just a few lines.
- 📺 Visualize simulations in real-time.

But don't get your hopes too high, this will be dots moving on a grid 😅 If you are looking for a proper robotics simulator, you should checkout [Gazebo](https://gazebosim.org/home).

XY Bot Simulator's design is actually very minimalist :

- Simulated world is a 2D squared grid.
- Robots move from cell to cell with a chosen velocity.
- Movement is discrete & limited to adjacent cells (no diagonal).

## Prerequisites

To run your simulations using XY Bot Simulator you will need Go version 1.17 or later, a C compiler and your system's development tools. Although XY Bot Simulator is virtually supported on any platforms with Go and a C compiler, it has been fully tested **only with macOS 13 on Apple Silicon**.

Using the standard Go tools you can install XY Bot Simulator's core library using:

```bash
go get github.com/floriankarydes/xybotsim
```

## Getting Started

Open a new file `my_simulation.go` and you're ready to setup your first simulation!

```go
package main

import (
  "github.com/floriankarydes/xybotsim"
)

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
```

And you can run that simply as:

```bash
go run my_simulation.go
```

It should look like this:

<img src="https://j.gifs.com/RlJn8q.gif" alt="First simulation preview" width="400" height="420">

The black square is a robot and the cells marked as occupied are greyed out. When moving from one cell to another, a robot occupy both cell for the duration of the movement.

## Going further

Please checkout the [examples](examples) available to learn how to run more complex simulation scenarios. Full documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/floriankarydes/xybotsim).
