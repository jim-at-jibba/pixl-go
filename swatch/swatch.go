package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/jim-at-jibba/pixl/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *apptype.State, color color.Color, SwatchIndex int, clickHandler func(s *Swatch)) {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		clickHandler: clickHandler,
		SwatchIndex:  SwatchIndex,
	}

	swatch.ExtendBaseWidget(swatch)
	return swatch
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}