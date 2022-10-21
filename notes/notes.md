# Notes

## Fyne Widget

```go
// To create a widget we must implement the Widget interface
type Widget interface {
  // Base functionality and state for all fyne widgets
  // (size, position etc).
  // Initialized with widget.ExtendBaseWidget(widget)
  CanvasObject
  // Createa the renderer which defins how the widget looks
  CreateRenderer() WidgetRenderer
}

type WidgetRenderer interface {
  // Deprecated: Ignore
  BackgroundColor() color.Color
  // Internal use only
  Destroy()
  //Calculate the position of the individual objects
  // that make up the widget
  Layout(Size)
  // Minimum dimensions that this widget occupies
  MinSize() Size
  // All objects that sohuld be drawn
  Objects() []CanvasObject
  // An update occurred and the widget needs to be redrawn
  Refresh()
}

// Mouse events
type Mouseable interface {
  MouseDown(*MouseEvent)
  MouseUp(*MouseEvent)
}

type Scrollable interface {
  Scrolled(*ScrollEvent)
}

type Hoverable interface {
  MouseIn(*MouseEvent)
  MouseMoved(*MouseEvent)
  MouseOut(*MouseEvent)
}
```
