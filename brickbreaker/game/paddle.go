package game

import "github.com/hajimehoshi/ebiten/v2"

type Paddle struct {
	X, Y          float64
	Width, Height float64
	Speed         float64
}

func NewPaddle() *Paddle {
	return &Paddle{
		X:      350,
		Y:      560,
		Width:  100,
		Height: 12,
		Speed:  6,
	}
}

func (p *Paddle) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && p.X > 0 {
		p.X -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && p.X+p.Width < ScreenWidth {
		p.X += p.Speed
	}
}
