package game

import (
	"fmt"
	"image/color"
	

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600

	BrickRows    = 5
	BrickCols    = 10
	BrickW       = 60
	BrickH       = 20
	BrickPadding = 8
	BrickTopOff  = 60
)

// ────────────────────────────────────────────────────────────

type Game struct {
	paddle  *Paddle
	ball    *Ball
	bricks  []*Brick
	lives   int
	score   int
	whitePx *ebiten.Image // 1×1 white pixel used for all rectangles
}

// Constructor ───────────────────────────────────────────────
func NewGame() *Game {
	g := &Game{
		paddle: NewPaddle(),
		ball:   NewBall(),
		lives:  3,
		score:  0,
		whitePx: func() *ebiten.Image {
			img := ebiten.NewImage(1, 1)
			img.Fill(color.White)
			return img
		}(),
	}

	// build brick grid
	for r := 0; r < BrickRows; r++ {
		for c := 0; c < BrickCols; c++ {
			x := float64(c)*(BrickW+BrickPadding) + BrickPadding
			y := BrickTopOff + float64(r)*(BrickH+BrickPadding)
			g.bricks = append(g.bricks, NewBrick(x, y))
		}
	}
	return g
}

// ────────────────────────────────────────────────────────────
// Update – game logic (called every tick)
func (g *Game) Update() error {
	g.paddle.Update()

	// move ball
	g.ball.X += g.ball.DX * g.ball.Speed
	g.ball.Y += g.ball.DY * g.ball.Speed

	// wall collisions
	if g.ball.X-g.ball.Radius < 0 || g.ball.X+g.ball.Radius > ScreenWidth {
		g.ball.DX *= -1
	}
	if g.ball.Y-g.ball.Radius < 0 {
		g.ball.DY *= -1
	}

	// paddle collision
	if g.ball.Y+g.ball.Radius >= g.paddle.Y &&
		g.ball.Y+g.ball.Radius <= g.paddle.Y+g.paddle.Height &&
		g.ball.X >= g.paddle.X && g.ball.X <= g.paddle.X+g.paddle.Width {
		g.ball.DY = -g.ball.DY
		offset := (g.ball.X - (g.paddle.X+g.paddle.Width/2)) / (g.paddle.Width/2)
		g.ball.DX = offset
	}

	// brick collisions
	for _, br := range g.bricks {
		if br.Broken {
			continue
		}
		if g.ball.X >= br.X && g.ball.X <= br.X+br.Width &&
			g.ball.Y-g.ball.Radius <= br.Y+br.Height && g.ball.Y+g.ball.Radius >= br.Y {
			br.Broken = true
			g.score++
			g.ball.DY *= -1
			break
		}
	}

	// ball fell below screen
	if g.ball.Y-g.ball.Radius > ScreenHeight {
		g.lives--
		if g.lives <= 0 {
			*g = *NewGame() // restart entire game
			return nil
		}
		g.ball.Reset() // just reset ball
	}

	return nil
}

// ────────────────────────────────────────────────────────────
// Draw – render everything
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255}) // background

	// paddle
	g.drawRect(screen, g.paddle.X, g.paddle.Y, g.paddle.Width, g.paddle.Height, color.White)

	// ball
	g.drawRect(screen, g.ball.X-g.ball.Radius, g.ball.Y-g.ball.Radius, g.ball.Radius*2, g.ball.Radius*2,
		color.RGBA{255, 100, 100, 255})

	// bricks
	for _, br := range g.bricks {
		if br.Broken {
			continue
		}
		g.drawRect(screen, br.X, br.Y, br.Width, br.Height, color.RGBA{100, 200, 250, 255})
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d  Lives: %d", g.score, g.lives))
}

func (g *Game) Layout(outsideW, outsideH int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// helper: draw colored rectangle via 1×1 white pixel
func (g *Game) drawRect(dst *ebiten.Image, x, y, w, h float64, clr color.Color) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(w, h)      // scale first
	op.GeoM.Translate(x, y)  // then move into place
	r, gcol, b, a := clr.RGBA()
	op.ColorM.Scale(float64(r)/0xffff, float64(gcol)/0xffff, float64(b)/0xffff, float64(a)/0xffff)
	dst.DrawImage(g.whitePx, op)
}

// ────────────────────────────────────────────────────────────
