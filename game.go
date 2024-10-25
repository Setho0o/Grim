package main

import (
	"os"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
  x int = 1920 
  y int = 1080
  px int = 0
  py int = 0
  speed int = 10
) 

type Game struct {
  x int 
  y int 
  s *ebiten.Image // screen
  p Player
}

type Player struct {
  x int 
  y int 
  img *ebiten.Image
  op *ebiten.DrawImageOptions
}

func GameInit() Game {
  return Game {
    x: x,
    y: y,
    s: nil,
    p: Player {
      x: px,
      y: py,
      img: ebiten.NewImage(10,10),
      op: &ebiten.DrawImageOptions{},
    },
  }
}

func MoveTo(x, y, dx, dy int) (float64, float64){
  return float64(-(x - dx)), float64(-(y - dy))
}

func Keys(g *Game) {
  if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
    g.p.op.GeoM.Translate(MoveTo(g.p.x,g.p.y,0,-speed))
  }
  if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
    g.p.op.GeoM.Translate(MoveTo(g.p.x,g.p.y,0,speed))
  }
  if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
    g.p.op.GeoM.Translate(MoveTo(g.p.x,g.p.y,-speed,0))
  }
  if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
    g.p.op.GeoM.Translate(MoveTo(g.p.x,g.p.y,speed,0))
  }
  if ebiten.IsKeyPressed(ebiten.KeyEscape) {
    os.Exit(0)
  }
}



