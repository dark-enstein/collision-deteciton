package object

import (
	"github.com/dark-enstein/collission-detection/experiments"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/xid"
)

type Window struct {
	dimension []*Coordinates //[0] - width, [1] - height
	id        string
}

func NewWindow(width, height float32) *Window {
	return &Window{dimension: []*Coordinates{
		NewCoordinates(0, 0),
		NewCoordinates(width, 0),
		NewCoordinates(0, height),
		NewCoordinates(width, height),
	},
		id: xid.New().String(),
	}
}

func (w *Window) Type() int {
	return OBJECT_WINDOW
}

func (w *Window) Displace(axis int, delta float32) experiments.Object {
	return nil
}

func (w *Window) CloneWithDisplace(axis int, delta float32) experiments.Object {
	return nil
}

func (w *Window) Clone() experiments.Object {
	return nil
}

func (w *Window) Render(e *ebiten.Image) {}

func (w *Window) Bounds(index int) (float32, float32) {
	return w.dimension[index].X(), w.dimension[index].Y()
}

func (w *Window) VertexLen() int {
	return len(w.dimension)
}

func (w *Window) ID() string {
	return w.id
}
