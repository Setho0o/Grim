package main

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) MapEditor() {
	for i, row := range g.m {
		for j, _ := range row {
			n := g.m[i][j]
			if *n.fill {
				g.FillNode(n)
			}
			g.s.Set(n.x, n.y, tan)
		}
	}

	x, y := ClosestNode(ebiten.CursorPosition())
	node := g.m[y][x]
	CurrentNode = node
  
  g.LineMode()
  g.SavedLines()
}

func (g *Game) LineMode() {
  if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
    HomeNode = CurrentNode
	}
  
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) {
    cnode = append(cnode, CurrentNode)
    hnode = append(hnode, HomeNode)
    HomeNode = CurrentNode
	}

  vector.StrokeLine(g.s,float32(HomeNode.x),float32(HomeNode.y),float32(CurrentNode.x),float32(CurrentNode.y),2,tan,false)
}
func (g *Game) SavedLines() {
  for i := range cnode {
    c := cnode[i]
    h := hnode[i]
    vector.StrokeLine(g.s,float32(h.x),float32(h.y),float32(c.x),float32(c.y),2,tan,false)
  }
}
func (g *Game) BlockMode() {
  g.FillNode(CurrentNode)
  
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		*CurrentNode.fill = true
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) {
		*CurrentNode.fill = false
	}

}

func (g *Game) UtilScreen() {
	img := ebiten.NewImage(g.x/3, g.y/3)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(MoveTo(0, 0, g.x-g.x/3, 0))

	img.Fill(tan)
	g.s.DrawImage(img, &op)

	cx, cy := ebiten.CursorPosition()

	DrawText(g.s, "x - "+strconv.Itoa(x), arcadeFontSrc, 20, float64(g.x-g.x/3), 20)
	DrawText(g.s, "y -"+strconv.Itoa(y), arcadeFontSrc, 20, float64(g.x-g.x/3), 40)
	DrawText(g.s, "cursor - "+strconv.Itoa(cx)+" "+strconv.Itoa(cy), arcadeFontSrc, 20, float64(g.x-g.x/3), 60)
	DrawText(g.s, "blocksize -"+strconv.Itoa(blocksize), arcadeFontSrc, 20, float64(g.x-g.x/3), 80)
	DrawText(g.s, "current node -"+strconv.Itoa(cx/blocksize)+" "+strconv.Itoa(cy/blocksize), arcadeFontSrc, 20, float64(g.x-g.x/3), 100)
}
