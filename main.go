package main

import (
	"github.com/dark-enstein/collission-detection/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(engine.WINDOW_WIDTH, engine.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Experiments")
	if err := ebiten.RunGame(engine.NewGame()); err != nil {
		log.Fatal(err)
	}
}
