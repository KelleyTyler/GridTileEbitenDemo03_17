package mypkgs

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	// "github.com/hajimehoshi/ebiten/v2/vector"
)

type IntegerGridManager struct {
	Imat             IntMatrix
	Coords           CoordList
	Tile_Size        CoordInts
	Margin           CoordInts
	BoardMargin      CoordInts //internal margin; the margin within IMG where things are drawn.
	Position         CoordInts
	BoardPosition    CoordInts
	CycleStart       int
	CycleEnd         int
	Colors           []color.Color
	FullColors       bool
	LastPoint        CoordInts //the last point clicked
	Fails            int
	AlgorithmRunning bool
	FailsMax         int
	//----
	MazeM MazeMaker
	//--------------------------------
	PFinder Pathfinding
	//---------------------------------
	PFinderEndSelect   bool
	PFinderStartSelect bool
	//--------------------------------
	SelectPoints bool
	//--------------
	Img *ebiten.Image

	Scale                          float64
	ScreenTicker                   int
	ScreenTicker_max               int
	BoardBuffer, BoardOverlayLayer *ebiten.Image
}

/* Muted Colors:
[]color.Color{color.RGBA{55, 55, 75, 255}, color.RGBA{125, 125, 150, 255}, color.RGBA{80, 180, 80, 255},
		color.RGBA{0, 150, 150, 255}, color.RGBA{65, 85, 85, 255}, color.RGBA{255, 255, 255, 255}, color.RGBA{75, 75, 75, 255}}


		[]color.Color{color.RGBA{55, 55, 75, 255}, color.RGBA{125, 125, 150, 255}, color.RGBA{80, 180, 80, 255},
		color.RGBA{0, 150, 150, 255}, color.RGBA{65, 105, 105, 255}, color.RGBA{255, 255, 255, 255}, color.RGBA{75, 75, 75, 255}}
*/

/*
//Vibrant Colors:
//

	[]color.Color{color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 255, 255, 255}, color.RGBA{255, 255, 0, 255}, color.RGBA{255, 255, 255, 255}, color.RGBA{75, 75, 75, 255}}
*/
func (igd *IntegerGridManager) Init(N_TilesX, N_TilesY int, TSizeX, TSizeY int, PosX, PosY int, MargX, MargY, iMargeX, iMargeY int) {
	igd.Margin = CoordInts{X: MargX, Y: MargY}
	igd.Position = CoordInts{X: PosX, Y: PosY}
	igd.Tile_Size = CoordInts{X: TSizeX, Y: TSizeY}
	igd.Colors = []color.Color{color.RGBA{55, 55, 75, 255}, color.RGBA{125, 125, 150, 255}, color.RGBA{80, 180, 80, 255},
		color.RGBA{0, 150, 150, 255}, color.RGBA{55, 65, 95, 255}, color.RGBA{255, 255, 255, 255}, color.RGBA{75, 75, 75, 255}}

	igd.Imat = igd.Imat.MakeIntMatrix(N_TilesX, N_TilesY) //color.RGBA{255, 255, 0, 255}
	igd.Imat.InitBlankMatrix(N_TilesX, N_TilesY, 0)
	igd.FullColors = false
	igd.CycleStart = 0
	igd.CycleEnd = 4
	igd.LastPoint = CoordInts{-1, -1}
	igd.Fails = 0
	igd.FailsMax = 30
	igd.AlgorithmRunning = false
	igd.PFinder = Pathfinding{IsActive: false, IsFullyInitialized: false, IsEndInit: false, HasFalsePos: false}
	//--------
	iX, iY := igd.Imat.GetCursorBounds(iMargeX+iMargeX-MargX, iMargeY+iMargeY-MargY, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)
	fmt.Printf("SCREEN SIZE: %d, %d\n", iX, iY)
	igd.Img = ebiten.NewImage(644, 644)
	igd.BoardBuffer = ebiten.NewImage(iX, iY)
	igd.BoardOverlayLayer = ebiten.NewImage(iX, iY)
	igd.Img.Fill(color.Black)

	igd.BoardMargin = CoordInts{X: iMargeX, Y: iMargeY}
	igd.BoardPosition = CoordInts{X: iMargeX, Y: iMargeY}
	igd.Imat.DrawGridTiles(igd.Img, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)
	igd.Scale = 1
	fmt.Printf("MAZEM\n")
	igd.MazeM.Init(&igd.Imat, 10)
	igd.ScreenTicker_max = 6
	igd.ScreenTicker = 0
}

func (igd *IntegerGridManager) Rescale(TSizeX, TSizeY, margX, margY int) {
	igd.Tile_Size = CoordInts{X: TSizeX, Y: TSizeY}
	igd.Margin = CoordInts{X: margX, Y: margY}
	iX, iY := igd.Imat.GetCursorBounds(igd.BoardMargin.X+(igd.BoardMargin.X-igd.Margin.X), igd.BoardMargin.Y+(igd.BoardMargin.Y-igd.Margin.Y), igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)
	igd.Img = ebiten.NewImage(644, 644)
	igd.BoardBuffer = ebiten.NewImage(iX, iY)
	igd.BoardOverlayLayer = ebiten.NewImage(iX, iY)
	igd.Img.Fill(color.Black)
}

func (igd *IntegerGridManager) Draw(screen *ebiten.Image) {
	// if igd.PFinder.IsStartInit {
	// 	igd.Imat[igd.PFinder.StartPos.Y][igd.PFinder.StartPos.X] = 5
	// }
	// if igd.PFinder.IsEndInit {
	// 	igd.Imat[igd.PFinder.EndPos.Y][igd.PFinder.EndPos.X] = 6
	// }
	if igd.ScreenTicker > igd.ScreenTicker_max {
		igd.RedrawBoard()
		igd.ScreenTicker = 0
	} else {
		igd.ScreenTicker++
	}

	// igd.Img.DrawImage(igd.BoardOverlayLayer, nil)
	// go igd.Imat.DrawGridTiles(igd.Img, 4, 4, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)

	// ops.GeoM.Reset()
	// go igd.BoardOverlayLayer.Clear()
	//igd.BoardOverlayLayer.Fill(color.RGBA{0, 150, 150, 50})
	// vector.DrawFilledRect(igd.BoardOverlayLayer, float32(-igd.BoardPosition.X+64), float32(-igd.BoardPosition.Y+64), 64.0, 64.0, color.RGBA{150, 0, 150, 255}, true)
	// ops.GeoM.Translate(float64(igd.Position.X-igd.BoardMargin.X), float64(igd.Position.Y-igd.BoardMargin.Y))//igd.BoardPosition.X, igd.BoardPosition.Y
	// ops.GeoM.Translate(float64(igd.BoardPosition.X-igd.BoardMargin.X), float64(igd.BoardPosition.Y-igd.BoardMargin.X))
	// ops.GeoM.Scale(igd.Scale, igd.Scale)
	// igd.Img.DrawImage(igd.BoardOverlayLayer, &ops)
	// ops.GeoM.Reset()

	// xx, yy := igd.BoardPosition.X-igd.BoardMargin.X, igd.BoardPosition.Y-igd.BoardMargin.Y //adjusted positions for the buffered area;
	//igd.Imat.DrawGridTiles(screen, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)
	if igd.PFinder.IsStartInit {
		// igd.Imat.DrawAGridTile(screen, igd.PFinder.StartPos, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{250, 250, 250, 255}, true)
		igd.Imat.DrawAGridTile(igd.Img, igd.PFinder.StartPos, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{250, 250, 250, 255}, true)
	}
	if igd.PFinder.IsEndInit {
		// igd.Imat.DrawAGridTile(screen, igd.PFinder.EndPos, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 50, 50, 255}, true)
		igd.Imat.DrawAGridTile(igd.Img, igd.PFinder.EndPos, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 50, 50, 255}, true)
		// igd.Imat.DrawAGridTile(igd.BoardOverlayLayer, igd.PFinder.EndPos, 0, 0, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 50, 50, 255}, true)

	}
	if igd.MazeM.Cords0_IsVisible {
		// igd.MazeM.DrawCoordLinesFromIGD(*igd, color.RGBA{100, 200, 200, 255}) igd.BoardPosition.X, igd.BoardPosition.Y,
		// igd.MazeM.Draw_CoordLines_raw(screen, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{150, 200, 150, 255})
		igd.MazeM.Draw_CoordLines_raw(igd.Img, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{150, 200, 150, 255})
		// igd.MazeM.Draw_CoordLines_raw(igd.BoardOverlayLayer, 0, 0, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{150, 200, 150, 255})
	}

	if igd.PFinder.IsFullyInitialized {
		//igd.Imat.DrawAGridTile(screen, igd.PFinder.Cursor.Position, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 140, 50, 255}, false)

		if igd.PFinder.HasFalsePos {
			// igd.Imat.DrawAGridTile(screen, igd.PFinder.FalsePos, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{140, 50, 50, 255}, false)

			for _, j := range igd.PFinder.FalsePos { //igd.BoardPosition.X, igd.BoardPosition.Y
				// igd.Imat.DrawAGridTile(screen, j, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{140, 50, 50, 255}, false)
				igd.Imat.DrawAGridTile(igd.Img, j, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{140, 50, 50, 255}, false)

			}
			for _, x := range igd.PFinder.Moves {
				// igd.Imat.DrawAGridTile(screen, x, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 125, 125, 255}, false)
				igd.Imat.DrawAGridTile(igd.Img, x, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 125, 125, 255}, false)

			}

		}
		igd.DrawCursor(igd.Img)
		// igd.Imat.DrawAGridTile_With_Line(screen, igd.PFinder.Cursor.Position, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{200, 150, 0, 255}, color.Black, 2.0, false)
		// igd.Imat.DrawAGridTile(screen, igd.PFinder.Cursor.Position, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{255, 15, 0, 1}, false) //{255,15,0,1}
	}
	// ops.GeoM.Translate(float64(igd.BoardPosition.X), float64(igd.BoardPosition.Y))
	// ops.GeoM.Scale(igd.Scale, igd.Scale)
	// igd.Img.DrawImage(igd.BoardOverlayLayer, nil)
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(igd.Position.X-igd.BoardMargin.X), float64(igd.Position.Y-igd.BoardMargin.Y))
	ops.GeoM.Scale(igd.Scale, igd.Scale)
	screen.DrawImage(igd.Img, &ops)
}
func (igd *IntegerGridManager) RedrawBoard() { //color.RGBA{20, 20, 20, 255} //color.RGBA{50, 50, 50, 255}
	igd.Img.Fill(color.RGBA{20, 20, 20, 255})
	igd.Imat.DrawGridTiles(igd.Img, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)
}

func (igd *IntegerGridManager) ResetCoordPosition() {
	igd.BoardPosition.X = igd.BoardMargin.X
	igd.BoardPosition.Y = igd.BoardMargin.Y
}
func IsCursorInBounds(PosX int, PosY int, Width, Height int) bool {
	mX, mY := ebiten.CursorPosition()
	return mX > PosX && mX < PosX+Width && mY > PosY && mY < PosY+Height
}
func IsCursorInBounds_02(mX, mY, PosX int, PosY int, Width, Height int) bool {
	return mX > PosX && mX < PosX+Width && mY > PosY && mY < PosY+Height
}
func (igd *IntegerGridManager) UpdateOnMouseEvent() {
	Raw_Mouse_X, Raw_Mouse_Y := ebiten.CursorPosition()
	tempX, tempY := -1, -1
	// temp2X, temp2Y := -1, -1
	isOnTile := false
	// isOnTile2 := false
	xx, yy := (igd.BoardPosition.X - igd.BoardMargin.X), (igd.BoardPosition.Y - igd.BoardMargin.Y)
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton(0)) && IsCursorInBounds_02(Raw_Mouse_X, Raw_Mouse_Y, igd.Position.X, igd.Position.Y, 644, 644) {
		//go igd.RedrawBoard()
		//temp0X, temp0Y, isOnTile = igd.Imat.GetCoordOfMouseEvent(Raw_Mouse_X-xx, Raw_Mouse_Y-yy, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)
		tempX, tempY, isOnTile = igd.Imat.GetCoordOfMouseEvent(Raw_Mouse_X-xx, Raw_Mouse_Y-yy, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)

		// temp2X, temp2Y, isOnTile = igd.Imat.GetCoordOfMouseEvent_Scalable(Raw_Mouse_X-igd.BoardPosition.X, Raw_Mouse_Y-igd.BoardPosition.Y, igd.Scale, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)
		// temp2X, temp2Y, isOnTile2 = igd.Imat.GetCoordOfMouseEvent_Scalable(Raw_Mouse_X-xx, Raw_Mouse_Y-yy, igd.Scale, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)

		if !igd.PFinderEndSelect && !igd.PFinderStartSelect && !igd.SelectPoints {
			if igd.FullColors {

				_, _ = igd.Imat.ChangeValOnMouseEvent(Raw_Mouse_X-xx, Raw_Mouse_Y-yy, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.CycleStart, len(igd.Colors)-1, !(igd.PFinderStartSelect || igd.PFinderEndSelect))

			} else {
				_, _ = igd.Imat.ChangeValOnMouseEvent(Raw_Mouse_X-xx, Raw_Mouse_Y-yy, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.CycleStart, igd.CycleEnd, !(igd.PFinderStartSelect || igd.PFinderEndSelect))
			}
		}
		//
		// // igd.MazeM.PrintString()
		// igd.AddToCoords(tempX, tempY)
		// igd.MazeM.Cords0 = igd.MazeM.Cords0.PushToReturn(igd.LastPoint)
		// fmt.Printf("ADDED\n")
		// fmt.Printf("%d %d %t\n", temp2X, temp2Y, isOnTile)
		// fmt.Printf("%d %d %t\n", temp2X, temp2Y, isOnTile2)
		if tempX != -1 && tempY != -1 {
			if !igd.PFinderEndSelect && !igd.PFinderStartSelect && igd.SelectPoints && isOnTile {
				igd.LastPoint = CoordInts{tempX, tempY}
				if (!igd.LastPoint.IsEqualTo(CoordInts{-1, -1})) {
					igd.MazeM.AddToCoords(tempX, tempY)
					//igd.Coords = igd.Coords.PushToReturn(igd.LastPoint)
				}
			} else if igd.PFinderEndSelect {
				igd.PFinder.EndPos = CoordInts{tempX, tempY}
				// igd.Imat[tempY][tempX] = 5
				igd.PFinder.IsEndInit = true
				igd.PFinderEndSelect = false
			} else if igd.PFinderStartSelect {
				igd.PFinder.StartPos = CoordInts{tempX, tempY}
				// igd.Imat[tempY][tempX] = 6
				igd.PFinder.IsStartInit = true
				igd.PFinderStartSelect = false
			}
		}
	}
}

func (igd *IntegerGridManager) DrawACircleOnClick(Raw_Mouse_X, Raw_Mouse_Y int, Radius int, valueIs int) {
	var center CoordInts
	var is_OnPoint bool
	xx, yy := (igd.BoardPosition.X - igd.BoardMargin.X), (igd.BoardPosition.Y - igd.BoardMargin.Y)
	center.X, center.Y, is_OnPoint = igd.Imat.GetCoordOfMouseEvent(Raw_Mouse_X-xx, Raw_Mouse_Y-yy, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)

	if is_OnPoint && IsCursorInBounds_02(Raw_Mouse_X, Raw_Mouse_Y, igd.Position.X, igd.Position.Y, 644, 644) {
		fmt.Printf("MIDPOINT CIRCLE\n\tCENTER (x,y): %3d,%3d\n\tRadius:%3d\n", center.X, center.Y, Radius)
		//temp := center
		P := 1 - Radius
		x := Radius
		y := 0
		for x > y {
			y++
			if P <= 0 {
				fmt.Printf("P is less than Or Equal to zero\n")
				P = P + 2*y + 1
			} else {
				fmt.Printf("P is Greater than zero\n")
				x--
				P = P + 2*y - 2*x + 1
			}
			//output here
			igd.circleDrawWSub(x, y, valueIs, center)
			if x < y {
				break
			}
		}

		temp_01A := center
		temp_01A.X += Radius
		// igd.Imat[temp_01A.Y][temp_01A.X] = valueIs
		if igd.Imat.IsValid(temp_01A) {
			igd.Imat[temp_01A.Y][temp_01A.X] = valueIs
		}
		temp_01B := center
		temp_01B.X -= Radius
		if igd.Imat.IsValid(temp_01B) {
			igd.Imat[temp_01B.Y][temp_01B.X] = valueIs
		}
		temp_01C := center
		temp_01C.Y -= Radius
		if igd.Imat.IsValid(temp_01C) {
			igd.Imat[temp_01C.Y][temp_01C.X] = valueIs
		}
		temp_01D := center
		temp_01D.Y += Radius
		if igd.Imat.IsValid(temp_01D) {
			igd.Imat[temp_01D.Y][temp_01D.X] = valueIs
		}

		// fmt.Printf("0: %d %d VALUEIS %d\n", center.X, center.Y, valueIs)
		// fmt.Printf("A: %d %d\n", temp_01A.X, temp_01A.Y)
		// fmt.Printf("B: %d %d\n", temp_01B.X, temp_01B.Y)
		// fmt.Printf("C: %d %d\n", temp_01C.X, temp_01C.Y)
		// fmt.Printf("D: %d %d\n", temp_01D.X, temp_01D.Y)
	}
}

func (igd *IntegerGridManager) circleDrawWSub(x, y, valueIs int, center CoordInts) {

	/*
			cout << "(" << x + x_centre << ", " << y + y_centre << ") ";
		        cout << "(" << -x + x_centre << ", " << y + y_centre << ") ";
		        cout << "(" << x + x_centre << ", " << -y + y_centre << ") ";
		        cout << "(" << -x + x_centre << ", " << -y + y_centre << ")\n";

	*/

	temp_01A := center
	// temp_01A.X += r
	// temp_01A.Y += r
	temp_01A.X += x
	temp_01A.Y += y

	temp_01B := center
	// temp_01B.X += r
	// temp_01B.Y += r
	temp_01B.X -= x
	temp_01B.Y += y

	temp_02A := center
	temp_02B := center
	// temp_02A.Y += r
	// temp_02A.X += r
	temp_02A.X += x
	temp_02A.Y -= y
	// temp_02B.Y += r
	// temp_02B.X += r
	temp_02B.X -= x
	temp_02B.Y -= y
	//temp.X += Radius
	// rsqure := math.Pow(float64(Radius), 2.0)
	//x^2 +y^2 = r^2
	if igd.Imat.IsValid(center) {
		igd.Imat[center.Y][center.X] = valueIs
	}
	if igd.Imat.IsValid(temp_01A) {
		igd.Imat[temp_01A.Y][temp_01A.X] = valueIs
	}
	if igd.Imat.IsValid(temp_01B) {
		igd.Imat[temp_01B.Y][temp_01B.X] = valueIs
	}
	if igd.Imat.IsValid(temp_02A) {
		igd.Imat[temp_02A.Y][temp_02A.X] = valueIs
	}
	if igd.Imat.IsValid(temp_02B) {
		igd.Imat[temp_02B.Y][temp_02B.X] = valueIs
	}

	if x != y {
		if igd.Imat.IsValid(CoordInts{center.X + y, center.Y + x}) {
			igd.Imat[center.Y+x][center.X+y] = valueIs
		}
		if igd.Imat.IsValid(CoordInts{center.X - y, center.Y + x}) {
			igd.Imat[center.Y+x][center.X-y] = valueIs
		}
		if igd.Imat.IsValid(CoordInts{center.X + y, center.Y - x}) {
			igd.Imat[center.Y-x][center.X+y] = valueIs
		}
		if igd.Imat.IsValid(CoordInts{center.X - y, center.Y - x}) {
			igd.Imat[center.Y-x][center.X-y] = valueIs
		}
	}
}

func (igd *IntegerGridManager) ToString() string {
	strng := "INTEGER GRID MANAGER:\n"
	strng += fmt.Sprintf("DIM %3d,%3d\n", len(igd.Imat), len(igd.Imat[0]))
	strng += fmt.Sprintf("Tiles: %3d,%3d\n", igd.Tile_Size.X, igd.Tile_Size.Y)
	strng += igd.Coords.ToString()
	strng += fmt.Sprintf("Last Point: %d,%d\nfails:%d\n", igd.LastPoint.X, igd.LastPoint.Y, igd.Fails)
	strng += "-------------\n"
	strng += igd.MazeM.ToString()
	return strng
}

func (igd *IntegerGridManager) DEMO_COORDS_00(a, x, y int) {
	switch a {
	case 0:
		igd.Coords = igd.Coords.PushToReturn(CoordInts{x, y})
	case 1:
		// igd.Coords, _ = igd.Coords.RemoveCoordFromList(CoordInts{x, y})
		igd.Imat.PrintMatrix()
	case 2:
		igd.Coords = igd.Coords.RemovePointFromList(x)
	case 3:
		igd.Coords, _ = igd.Coords.RemoveCoordFromList(igd.LastPoint)
	case 4:
		igd.Coords.PrintCordArray()
	case 5:
		igd.Coords = igd.Coords.SortDescOnX()
	case 6:
		igd.Coords = igd.Coords.RemoveDuplicates()
	case 7:

	default:
		fmt.Printf("DEFAULT DEMO COORDS_00")
	}

}

func (igd *IntegerGridManager) ClearImat() {
	for i, a := range igd.Imat {
		for j, _ := range a {
			igd.Imat[i][j] = 0
		}
	}
}

// func (igd *IntegerGridManager) DrawCoordsOnImat() {
// 	//igd.ClearImat()
// 	for _, c := range igd.Coords {
// 		if igd.Imat[c.Y][c.X] != 4 {
// 			igd.Imat[c.Y][c.X] = 2
// 		}
// 	}
// }
