package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"brickbreaker/game"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("🎮 Brick Breaker (Go + Ebiten)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
