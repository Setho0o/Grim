package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	x         int = 1920
	y         int = 1080
)

var (
	blocksize int = 15 // block size must be the same as block litmit or problems 
  blocklimit int = 15
	px     int                                    = 0
	py     int                                    = 0
	speed  int                                    = 10
	mx int = x/blocksize + 1
	my int = y/blocksize + 1
  CurrentNode Node
)

type Game struct {
	x int
	y int
	s *ebiten.Image // screen
	p Player
	m [][]Node
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
			img: ebiten.NewImage(blocksize, blocksize),
			op:  &ebiten.DrawImageOptions{},
		},
		m: MatrixInit(blocksize),
	}
}

func MatrixInit(size int) [][]Node{
  m  := make([][]Node, my)
  
  for i := range m {
    m[i] = make([]Node, mx)
  }

	for i, row := range m {
		for j, _ := range row {
			x := false
			m[i][j] = Node{x: j * blocksize, y: i * blocksize, fill: &x}
			m[i][j].Bounds()
		}
	}

  return m
}
func (g *Game) NewMatrix(NewSize int) {
  if NewSize > blocklimit {
    size := &blocksize

    NewMatrix := MatrixInit(NewSize)
    OldMatrix := g.m 
    if len(NewMatrix) > len(OldMatrix) {
      for i, e := range OldMatrix {
        for j, _ := range e {
          if *OldMatrix[i][j].fill {
            *NewMatrix[i][j].fill = true
          }
        }
      }
    } else {
      for i, e := range NewMatrix {
        for j, _ := range e {
         if *OldMatrix[i][j].fill {
            *NewMatrix[i][j].fill = true
          } 
        }
      }
    }
    *size = NewSize
    g.m = NewMatrix
  }
}

func MoveTo(x, y, dx, dy int) (float64, float64) {
	return float64(-(x - dx)), float64(-(y - dy))
}

func (g *Game) Keys() {
  x, y := ebiten.Wheel()
  g.NewMatrix(blocksize + int(x))
  g.NewMatrix(blocksize - int(y))



	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.p.op.GeoM.Translate(MoveTo(g.p.x, g.p.y, 0, -speed))


    g.NewMatrix(blocksize + 1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.p.op.GeoM.Translate(MoveTo(g.p.x, g.p.y, 0, speed))
    g.NewMatrix(blocksize - 1)
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
