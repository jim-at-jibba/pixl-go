package ui

import (
	"fyne.io/fyne/v2"
	"github.com/jim-at-jibba/pixl/apptype"
	"github.com/jim-at-jibba/pixl/pxcanvas"
	"github.com/jim-at-jibba/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
