package main

import (
	"image/color"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
  Keys(g) 
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  g.s = screen
  g.p.img.Fill(color.White)
  g.s.DrawImage(g.p.img, g.p.op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.x, g.y
}

func main() {
  g := GameInit()
	ebiten.SetWindowSize(g.x, g.x)
	ebiten.SetWindowTitle("Grim")

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
