package object

import (
	"github.com/dark-enstein/collission-detection/experiments"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/xid"
	"log"
)

type Coordinates struct {
	x, y float32
	id   string
}

func NewCoordinates(x, y float32) *Coordinates {
	return &Coordinates{
		x:  x,
		y:  y,
		id: xid.New().String(),
	}
}

func (c *Coordinates) X() float32 {
	return c.x
}

func (c *Coordinates) Y() float32 {
	return c.y
}

func (c *Coordinates) Type() int {
	return OBJECT_COORDINATE
}

func (c *Coordinates) Displace(axis int, delta float32) experiments.Object {
	switch axis {
	case experiments.AXIS_X:
		c.x += delta
	case experiments.AXIS_Y:
		c.y += delta
	}
	return c
}

func (c *Coordinates) ID() string {
	return c.id
}

func (c *Coordinates) Clone() experiments.Object {
	return NewCoordinates(c.x, c.y)
}

func (c *Coordinates) CloneWithDisplace(axis int, delta float32) experiments.Object {
	switch axis {
	case experiments.AXIS_X:
		return NewCoordinates(c.x+delta, c.y)
	case experiments.AXIS_Y:
		return NewCoordinates(c.x, c.y+delta)
	}
	log.Println("axis not found")
	return &Coordinates{}
}

func (c *Coordinates) Render(e *ebiten.Image) {
}

func (c *Coordinates) Bounds(index int) (float32, float32) {
	return c.x, c.y
}

func (c *Coordinates) VertexLen() int {
	return 1
}
