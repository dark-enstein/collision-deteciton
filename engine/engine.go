package engine

import (
	"github.com/dark-enstein/collission-detection/engine/object"
	"github.com/dark-enstein/collission-detection/experiments"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	WINDOW_WIDTH  = 640
	WINDOW_HEIGHT = 480
)

type Game struct {
	window *object.Window
	object experiments.Object
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	if g.object == nil {
		coord := object.NewCoordinates(WINDOW_WIDTH/3, WINDOW_HEIGHT/3)
		g.object = object.NewSquare(coord, 20)
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.object.Displace(experiments.AXIS_X, 5)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.object.Displace(experiments.AXIS_X, -5)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.object.Displace(experiments.AXIS_Y, 5)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.object.Displace(experiments.AXIS_Y, -5)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.object.Render(screen)
	//ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
