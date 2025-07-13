package game

type Brick struct {
	X, Y          float64
	Width, Height float64
	Broken        bool
}

func NewBrick(x, y float64) *Brick {
	return &Brick{
		X:      x,
		Y:      y,
		Width:  BrickW,
		Height: BrickH,
	}
}
