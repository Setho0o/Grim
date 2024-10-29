package main

import (
	"os"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	blocksize int = 60
	x         int = 1920
	y         int = 1080
)

var (
	px     int                                   = 0
	py     int                                   = 0
	speed  int                                   = 10
	matrix [y/blocksize + 1][x/blocksize + 1]Node // 18, 32
)


type Game struct {
	x int
	y int
	s *ebiten.Image // screen
	p Player
	m [y/blocksize + 1][x/blocksize + 1]Node
}

type Player struct {
	x   int
	y   int
	img *ebiten.Image
	op  *ebiten.DrawImageOptions
}

func GameInit() Game {
	return Game{
		x: x,
		y: y,
		s: nil,
		p: Player{
			x:   px,
			y:   py,
			img: ebiten.NewImage(blocksize,blocksize),
			op:  &ebiten.DrawImageOptions{},
		},
		m: MatrixInit(),
	}
}

func MoveTo(x, y, dx, dy int) (float64, float64) {
	return float64(-(x - dx)), float64(-(y - dy))
}

func (g *Game) Keys() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.p.op.GeoM.Translate(MoveTo(g.p.x, g.p.y, 0, -speed))
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.p.op.GeoM.Translate(MoveTo(g.p.x, g.p.y, 0, speed))
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.p.op.GeoM.Translate(MoveTo(g.p.x, g.p.y, -speed, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.p.op.GeoM.Translate(MoveTo(g.p.x, g.p.y, speed, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}
