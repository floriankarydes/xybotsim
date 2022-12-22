# XY Bot Simulator

Extremely simple simulation of a set of ["turtlebot"](https://www.turtlebot.com) style robots.

## Features

- 🤖 Test robot controllers written in [Go](https://go.dev) with software-in-the-loop simulations
- 💯 Run large scale multi-robots scenarios
- 💥 Monitor robot collisions with obstacles and other robots
- 📘 Connect your code easily with the most simple [API](#api)

But don't get your hopes too high, this will be dots moving on a grid 😅 If you are looking for an actual simulator for
robots, you should checkout [Gazebo](https://gazebosim.org/home).

XY Bot Simulator's specification is very minimalist :

- Simulated world is a 2D squared grid.
- Robot moves from cell to cell with a chosen velocity.
- Movement is limited to adjacent cells (no diagonal).
- Expose two simple interfaces
-
  ```go
  type Robot interface {
    GetPosition() (x uint, y uint)
    EnqueueCommand(d Direction) error
  }

  type Simulator interface {
    AddRobot(id string) (*Robot, error)
    DeleteRobot(id string) error
    ListRobots() []*Robot
  }
  ```

## Getting Started

⚠️ This is a work in progress ⚠️
