package game

import "math"

type Ball struct {
	X, Y          float64
	DX, DY        float64
	Radius        float64
	Speed         float64
}

func NewBall() *Ball {
	b := &Ball{
		Radius: 8,
		Speed:  4,
	}
	b.Reset()
	return b
}

func (b *Ball) Reset() {
	b.X = ScreenWidth / 2
	b.Y = ScreenHeight / 2
	angle := -math.Pi / 4
	b.DX = math.Cos(angle)
	b.DY = math.Sin(angle)
}
