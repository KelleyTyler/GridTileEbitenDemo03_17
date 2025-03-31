package mypkgs

import (
	"fmt"

	//"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	passable bool
	position CoordInts
	img      *ebiten.Image
	value    int
	name     string
}
type TileGrid struct {
}

func (t *Tile) ToString() string {
	return fmt.Sprintf("name: %s\nPOS:%3d,%3d\nPassable:%t VALUE: %2d", t.name, t.position.X, t.position.Y, t.passable, t.value)
}
func (t *Tile) PrintTile() {
	fmt.Printf("%s", t.ToString())
}

// func (t *Tile)GetTILE() {}

// func (imat IntMatrix) ClearAnArea(c1_X, c1_Y, c2_X, c2_Y, val int) {

// 	if c1_Y < len(imat) && c1_Y > -1 {
// 		if c2_Y > -1 && c2_Y < len(imat) {
// 			for i := c1_Y; i < c2_Y; i++ {
// 				for j := c1_X; j < c2_X; j++ {
// 					imat[i][j] = val
// 				}
// 			}
// 		}
// 	}
// }

func (imat IntMatrix) ClearAnArea(c0_X, c0_Y, c1_X, c1_Y, val int) {

	if c0_Y < len(imat) && c0_Y > -1 {
		if c1_Y > -1 && c1_Y < len(imat) {
			for i := c0_Y; i < c1_Y; i++ {
				for j := c0_X; j < c1_X; j++ {
					imat[i][j] = val
				}
			}
		}
	}
}
