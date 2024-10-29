package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.Keys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.s = screen
  g.MapEditor()  
  g.UtilScreen()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.x, g.y
}

func main() {
	g := GameInit()
	ebiten.SetWindowSize(g.x, g.x)
	ebiten.SetWindowTitle("Grim")
  ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
