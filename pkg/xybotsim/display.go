package xybotsim

import (
	"time"

	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type renderer struct {
	simulatorHandle *Simulator
	imageCache      *image.RGBA
	freeColor       color.Color
	occupiedColor   color.Color
	robotColor      color.Color
}

func (r *renderer) draw(w int, h int) image.Image {

	var screenSize, worldSize int
	if w < h {
		screenSize = w
	} else {
		screenSize = h
	}
	worldSize = int(r.simulatorHandle.worldSize)

	img := r.imageCache
	if img == nil || img.Bounds().Size().X != screenSize || img.Bounds().Size().Y != screenSize {
		img = image.NewRGBA(image.Rect(0, 0, screenSize, screenSize))
		r.imageCache = img
	}

	for pY := 0; pY < screenSize; pY++ {
		cY := pY * worldSize / screenSize
		for pX := 0; pX < screenSize; pX++ {
			cX := pX * worldSize / screenSize
			if r.simulatorHandle.occupancyGrid[cX][cY] {
				img.Set(pX, pY, r.occupiedColor)
			} else {
				img.Set(pX, pY, r.freeColor)
			}
		}
	}

	cellSize := screenSize / worldSize
	robotSize := cellSize * 8 / 10
	marginSize := (cellSize - robotSize) / 2
	for _, robot := range r.simulatorHandle.robotRegister {
		pYMin := int(robot.currentCell.y)*screenSize/worldSize + marginSize
		pYMax := pYMin + robotSize
		pXMin := int(robot.currentCell.x)*screenSize/worldSize + marginSize
		pXMax := pXMin + robotSize
		for pY := pYMin; pY < pYMax; pY++ {
			for pX := pXMin; pX < pXMax; pX++ {
				img.Set(pX, pY, r.robotColor)
			}
		}
	}
	return img
}

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
