package xybotsim

import (
	"time"

	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

// Display a minimal visualization of the simulator's world on screen.
//
// This should be called near the end of a main() function as it will block until window is closed.
func (s *Simulator) Show() {
	application := app.New()
	window := application.NewWindow("XY Bot Simulator")
	window.Resize(fyne.NewSize(500, 500))

	r := renderer{
		simulatorHandle: s,
		freeColor:       color.RGBA{0, 0, 0, 0},
		occupiedColor:   color.RGBA{0, 0, 0, 20},
		robotColor:      color.RGBA{0, 0, 0, 250},
	}
	worldCanvas := canvas.NewRaster(r.draw)
	worldCanvas.ScaleMode = canvas.ImageScalePixels
	window.SetContent(worldCanvas)

	go func() {
		tick := time.NewTicker(time.Second / 6)
		for range tick.C {
			worldCanvas.Refresh()
		}
	}()

	window.ShowAndRun()
}

// Persistent object to render new frame of the simulation.
type renderer struct {
	simulatorHandle *Simulator  // Handle to the simulator to be rendered.
	imageCache      *image.RGBA // Previous frame.
	freeColor       color.Color // Color of free cells.
	occupiedColor   color.Color // Color of occupied cells.
	robotColor      color.Color // Color of robots.
}

func (r *renderer) draw(w int, h int) image.Image {

	// Parse screen size (i.e. pixels) & world size (i.e. cells).
	var screenSize, worldSize int
	if w < h {
		screenSize = w
	} else {
		screenSize = h
	}
	worldSize = int(r.simulatorHandle.worldSize)

	// Initialize image cache if needed
	img := r.imageCache
	if img == nil || img.Bounds().Size().X != screenSize || img.Bounds().Size().Y != screenSize {
		img = image.NewRGBA(image.Rect(0, 0, screenSize, screenSize))
		r.imageCache = img
	}

	// Draw the occupancy map (i.e. all cells were robot cannot move to) in screen space
	for pY := 0; pY < screenSize; pY++ {
		cY := pY * worldSize / screenSize
		for pX := 0; pX < screenSize; pX++ {
			cX := pX * worldSize / screenSize
			if r.simulatorHandle.worldGrid[cX][cY] {
				img.Set(pX, pY, r.occupiedColor)
			} else {
				img.Set(pX, pY, r.freeColor)
			}
		}
	}

	// Draw the robots in screen space
	cellSize := screenSize / worldSize
	robotSize := cellSize * 8 / 10
	marginSize := (cellSize - robotSize) / 2
	for _, robot := range r.simulatorHandle.robotRegister {
		pYMin := int(robot.position.y)*screenSize/worldSize + marginSize
		pYMax := pYMin + robotSize
		pXMin := int(robot.position.x)*screenSize/worldSize + marginSize
		pXMax := pXMin + robotSize
		for pY := pYMin; pY < pYMax; pY++ {
			for pX := pXMin; pX < pXMax; pX++ {
				img.Set(pX, pY, r.robotColor)
			}
		}
	}
	return img
}
