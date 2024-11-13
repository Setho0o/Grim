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

func MatrixInit() Matrix {
  xy := [x/blocksize]Yes {
    N: [y / blocksize]Node{

    },
  }
  matrix := Matrix {
    x: xy,
    blocksize: blocksize,
  } 
	for i, row := range matrix.x {
		for j, e := range row.N {
			x := false
			e = Node{x: j * blocksize, y: i * blocksize, fill: &x}
			e.Bounds()
		}
	}
	return matrix
} 

func (g *Game) MapEditor() {
  g.BlockMode()
 // g.PointMode()
}
func (g *Game) BlockMode() {
  for _, row := range g.m.x {
		for _, e := range row.y {
			n := e
			if *n.fill {
				g.FillNode(n)
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
/*
	x, y := ClosestNode(ebiten.CursorPosition())
  x := g.m.x[x]
	node := g.m[y][x]
  currentNode = &node 
	g.FillNode(node)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		*node.fill = true
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) {
		*node.fill = false
	}
*/
}

/*
func (g *Game) PointMode() {
  for i, row := range matrix {
		for j, _ := range row {
			n := matrix[i][j]
			if *n.fill {
				g.FillNode(n)
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
/*
		}
	}
	x, y := ClosestNode(ebiten.CursorPosition())
	node := g.m[y][x]
  currentNode = &node 
	g.FillNode(node)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		*node.fill = true
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) {
		*node.fill = false
	}
}
*/
func (g *Game) UtilScreen() {
  
	img := ebiten.NewImage(g.x/3, g.y/3)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(MoveTo(0, 0, g.x-g.x/3, 0))
	img.Fill(tan)

  DrawText(img,"blocksize: "+strconv.Itoa(blocksize),arcadeFontSrc, 10, 10 ) // fix scaling 
  DrawText(img,"x: "+strconv.Itoa(x),arcadeFontSrc, 10, 30 ) // fix scaling 
  DrawText(img,"y: "+strconv.Itoa(y),arcadeFontSrc, 10, 50 ) // fix scaling 
  cx, cy := ebiten.CursorPosition() 
  DrawText(img,"cursor pos: "+strconv.Itoa(cx)+" "+strconv.Itoa(cy),arcadeFontSrc, 10, 70 ) // fix scaling 
  DrawText(img,"current node: "+strconv.Itoa(currentNode.x / blocksize)+" "+strconv.Itoa(currentNode.y / blocksize),arcadeFontSrc, 10, 90 ) 
	
 // DrawText(img,"nodes filled: "+strconv.Itoa(g.FilledNodes()),arcadeFontSrc, 10, 110 ) // fix scaling 
  g.s.DrawImage(img, &op)

}
/*
func (g *Game) FilledNodes() int {
  nodesFilled := 0
	for i, row := range matrix {
		for j, _ := range row {
			n := matrix[i][j]
      if *n.fill {
        nodesFilled++
      }
		}
	}
  return nodesFilled
}
*/
func (g *Game) FillNode(n Node) {
	for _, x := range n.bx {
		for _, y := range n.by {
			g.s.Set(x, y, color.White)
		}
	}
}
