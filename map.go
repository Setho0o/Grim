package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Node struct {
	x    int
	y    int
	bx   []int
	by   []int
	fill *bool
}

func (c *Node) Bounds() {
	for x := c.x; x < c.x+blocksize; x++ {
		c.bx = append(c.bx, x)
	}
	for y := c.y; y < c.y+blocksize; y++ {
		c.by = append(c.by, y)
	}
}

func ClosestNode(x, y int) (int, int) {
	return x / blocksize, y / blocksize

}

func MatrixInit() [y/blocksize + 1][x/blocksize + 1]Node {
	for i, row := range matrix {
		for j, _ := range row {
			x := false
			matrix[i][j] = Node{x: j * blocksize, y: i * blocksize, fill: &x}

			matrix[i][j].Bounds()
		}
	}
	return matrix
}

func (g *Game) MapEditor() {
	for i, row := range matrix {
		for j, _ := range row {
			n := matrix[i][j]
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

}

func (g *Game) FillCell(n Node) {
	for _, x := range n.bx {
		for _, y := range n.by {
			g.s.Set(x, y, color.White)
		}
	}
}
