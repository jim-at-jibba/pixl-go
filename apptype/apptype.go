package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasArea     fyne.Position
	PxRows, PxCols int
	PxSize         int
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int
	filePath       string
}

func (state *State) SetFilePath(path string) {
	state.filePath = path
}
