package experiments

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	AXIS_X = iota
	AXIS_Y
)

type Collision interface {
	WillCollide(one Object) bool
}

type Object interface {
	Type() int
	ID() string
	Displace(axis int, delta float32) Object
	CloneWithDisplace(axis int, delta float32) Object
	Clone() Object
	Render(e *ebiten.Image)
	VertexLen() int
	Bounds(index int) (x, y float32)
}
