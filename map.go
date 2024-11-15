package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type Node struct {
	x    int
	y    int
	bx   []int
	by   []int
	fill *bool
  adj []Node
}
func (c *Node) Bounds() {
	for x := c.x; x < c.x+blocksize; x++ {
		c.bx = append(c.bx, x)
	}
	for y := c.y; y < c.y+blocksize; y++ {
		c.by = append(c.by, y)
	}
}
func (n *Node) AdjNodes() {
}

func ClosestNode(x, y int) (int, int) {
	return x / blocksize, y / blocksize

}

func (g *Game) MapEditor() {
	for i, row := range g.m {
		for j, _ := range row {
			n := g.m[i][j]
			if *n.fill {
				g.FillCell(n)
			}
			g.s.Set(n.x, n.y, tan)

			/* tiles
			   for i := 0; i < len(loc.bx); i++ {
			     g.s.set(loc.bx[i], loc.y, c)
			   }
			   for i := 0; i < len(loc.by); i++ {
			     g.s.set(loc.x, loc.by[i], c)
			   }
			*/
		}
	}
  
	x, y := ClosestNode(ebiten.CursorPosition())
	cell := g.m[y][x]
  CurrentNode = cell	

  g.FillCell(cell)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		*cell.fill = true
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) {
		*cell.fill = false
	}
}

func (g *Game) UtilScreen() {

	img := ebiten.NewImage(g.x/3, g.y/3)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(MoveTo(0, 0, g.x-g.x/3, 0))

	img.Fill(tan)
	g.s.DrawImage(img, &op)
  
  cx, cy := ebiten.CursorPosition()  

  DrawText(g.s, "x - " + strconv.Itoa(x), arcadeFontSrc,20,float64(g.x-g.x/3),20)
  DrawText(g.s, "y -" + strconv.Itoa(y), arcadeFontSrc,20,float64(g.x-g.x/3),40)
  DrawText(g.s, "cursor - "+strconv.Itoa(cx)+" "+strconv.Itoa(cy), arcadeFontSrc,20,float64(g.x-g.x/3),60)
  DrawText(g.s, "blocksize -"+strconv.Itoa(blocksize), arcadeFontSrc,20,float64(g.x-g.x/3),80)
  DrawText(g.s, "current node -"+strconv.Itoa(cx/blocksize )+" "+strconv.Itoa(cy/blocksize), arcadeFontSrc,20,float64(g.x-g.x/3),100)



}

func (g *Game) FillCell(n Node) {
	for _, x := range n.bx {
		for _, y := range n.by {
			g.s.Set(x, y, color.White)
		}
	}
}
