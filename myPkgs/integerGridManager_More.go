package mypkgs

import "image/color"

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

func (igd *IntegerGridManager) GetIMatDisplayOptions(colors []color.Color, ShowLines []bool, lineThicknesses []float32) IntMatrix_EbtnDisplay_Options {
	temp := IntMatrix_EbtnDisplay_Options{
		TileSize:             CoordInts{igd.Tile_Size.X, igd.Tile_Size.Y},
		TileSpacing:          CoordInts{igd.Margin.X, igd.Margin.Y},
		BoardMargin:          CoordInts{igd.BoardMargin.X, igd.BoardMargin.Y},
		BoardPosition:        CoordInts{igd.BoardPosition.X, igd.BoardPosition.Y},
		ShowTileOutline:      ShowLines[0],
		ShowTileDLine0:       ShowLines[1],
		ShowTileDLine1:       ShowLines[2],
		TileOLThickness:      lineThicknesses[0],
		TileDLine0_Thickness: lineThicknesses[1],
		TileDLine1_Thickness: lineThicknesses[2],
	}
	return temp
}
