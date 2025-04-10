package mypkgs

import "image/color"

/*
type IntMatrix_EbtnDisplay_Options struct {
	TileSize             CoordInts
	TileSpacing          CoordInts
	ShowTileOutline      bool
	TileOLThickness      float32
	ShowTileLines     []bool
	TileLineColors    []color.Color
	TileLineThickness []float32
}

*/

func (igd *IntegerGridManager) GetIMatDisplayOptions(colors []color.Color, ShowLines []bool, lineThicknesses []float32) IntMatrix_EbtnDisplay_Options {

	temp := IntMatrix_EbtnDisplay_Options{
		TileSize:      CoordInts{igd.Tile_Size.X, igd.Tile_Size.Y},
		TileSpacing:   CoordInts{igd.Margin.X, igd.Margin.Y},
		BoardMargin:   CoordInts{igd.BoardMargin.X, igd.BoardMargin.Y},
		BoardPosition: CoordInts{igd.BoardPosition.X, igd.BoardPosition.Y},

		TileLineColors:    colors,
		TileLineThickness: lineThicknesses,
		ShowTileLines:     ShowLines,
		// ShowTileOutline:      ShowLines[0],
		// ShowTileDLine0:       ShowLines[1],
		// ShowTileDLine1:       ShowLines[2],
		// TileOLThickness:      lineThicknesses[0],
		// TileDLine0_Thickness: lineThicknesses[1],
		// TileDLine1_Thickness: lineThicknesses[2],
	}

	if len(ShowLines) < 3 {
		temp.ShowTileLines = []bool{false, false, false}
	}
	if len(lineThicknesses) < 3 {
		temp.TileLineThickness = []float32{0.0, 0.0, 0.0}
	}
	if len(colors) < 3 {
		temp.TileLineColors = []color.Color{color.Black, color.Black, color.Black}
	}
	return temp
}
func (igd *IntegerGridManager) GetIMatDisplayOptions_ptr(colors []color.Color, ShowLines []bool, lineThicknesses []float32) *IntMatrix_EbtnDisplay_Options {

	temp := igd.GetIMatDisplayOptions(colors, ShowLines, lineThicknesses)
	return &temp
}
