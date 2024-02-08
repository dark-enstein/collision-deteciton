package object

import (
	"github.com/dark-enstein/collission-detection/experiments"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/rs/xid"
	"image/color"
)

type Square struct {
	vs []*Coordinates
	id string
}

func NewSquare(startingVertice *Coordinates, sideLength float32) *Square {
	v2, v3, v4 := startingVertice.CloneWithDisplace(experiments.AXIS_X, sideLength), startingVertice.CloneWithDisplace(experiments.AXIS_Y, sideLength), startingVertice.CloneWithDisplace(experiments.AXIS_Y, sideLength).Displace(experiments.AXIS_X, sideLength)
	return &Square{
		vs: []*Coordinates{startingVertice, v2.(*Coordinates), v3.(*Coordinates), v4.(*Coordinates)},
		id: xid.New().String(),
	}
}

func (s *Square) Type() int {
	return OBJECT_SQUARE
}

func (s *Square) Render(e *ebiten.Image) {
	startingX := s.vs[0].x
	startingY := s.vs[0].y
	width := s.vs[1].x - s.vs[0].x
	height := s.vs[2].y - s.vs[0].y
	vector.StrokeRect(e, startingX, startingY, width, height, 2, color.White, false)
}

func (s *Square) Displace(axis int, delta float32) experiments.Object {
	for i := 0; i < len(s.vs); i++ {
		switch axis {
		case experiments.AXIS_X:
			s.vs[i].x += delta
		case experiments.AXIS_Y:
			s.vs[i].y += delta
		}
	}
	return s
}

func (s *Square) ID() string {
	return s.id
}

func (s *Square) CloneWithDisplace(axis int, delta float32) experiments.Object {
	return &Square{}
}

func (s *Square) Clone() experiments.Object {
	return &Square{}
}

func (s *Square) Bounds(index int) (float32, float32) {
	return s.vs[index].X(), s.vs[index].Y()
}

func (s *Square) VertexLen() int {
	return len(s.vs)
}
