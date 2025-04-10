package mypkgs

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
type IntMatrix_EbtnDisplay_Options struct {
	TileSize             CoordInts
	TileSpacing          CoordInts
	ShowTileOutline      bool
	TileOLThickness      float32
	ShowTileDLine0       bool
	TileDLine0_Thickness float32
	ShowTileDLine1       bool
	TileDLine1_Thickness float32
	BoardMargin          CoordInts
	BoardPosition        CoordInts
}

*/

func (node *Node) ShowOnImat(screen *ebiten.Image, ui_helper *UI_Helper, imat IntMatrix, options IntMatrix_EbtnDisplay_Options, colors []color.Color) {
	node.showOnImat_Helper(screen, ui_helper, imat, options, colors, 0, false)
}

func (node *Node) showOnImat_Helper(screen *ebiten.Image, ui_helper *UI_Helper, imat IntMatrix, options IntMatrix_EbtnDisplay_Options, colors []color.Color, num int, active bool) {
	if active {
		tempColorVal := num
		for tempColorVal > len(colors)-1 {
			tempColorVal = tempColorVal % (len(colors) - 1)
		}
		imat.DrawAGridTile_With_Lines(screen, node.Postion, colors[tempColorVal], &options)
		if node.ChildPTR != nil {

			node.ChildPTR.showOnImat_Helper(screen, ui_helper, imat, options, colors, num+1, true)
		}
	} else {
		if node.ParentPTR != nil {
			node.ParentPTR.showOnImat_Helper(screen, ui_helper, imat, options, colors, 0, false)
		} else {

			if node.ChildPTR != nil {
				// tempColorVal := num
				imat.DrawAGridTile_With_Lines(screen, node.Postion, colors[0], &options)

				node.ChildPTR.showOnImat_Helper(screen, ui_helper, imat, options, colors, num+1, true)
			}
		}
	}
}
