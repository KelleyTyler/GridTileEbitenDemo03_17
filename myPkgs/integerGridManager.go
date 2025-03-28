package mypkgs

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type IntegerGridManager struct {
	Imat             IntMatrix
	Coords           CoordList
	Tile_Size        CoordInts
	Margin           CoordInts
	BoardMargin      CoordInts //internal margin; the margin within IMG where things are drawn.
	Position         CoordInts
	CycleStart       int
	CycleEnd         int
	Colors           []color.Color
	FullColors       bool
	LastPoint        CoordInts //the last point clicked
	Fails            int
	AlgorithmRunning bool
	FailsMax         int
	PFinder          Pathfinding
	//---------------------------------
	PFinderEndSelect   bool
	PFinderStartSelect bool
	//--------------------------------
	SelectPoints bool
	//--------------
	Img   *ebiten.Image
	Scale float64
}

func (igd *IntegerGridManager) Init(N_TilesX, N_TilesY int, TSizeX, TSizeY int, PosX, PosY int, MargX, MargY, iMargeX, iMargeY int) {
	igd.Margin = CoordInts{X: MargX, Y: MargY}
	igd.Position = CoordInts{X: PosX, Y: PosY}
	igd.Tile_Size = CoordInts{X: TSizeX, Y: TSizeY}
	igd.Colors = []color.Color{color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 255, 255, 255}, color.RGBA{255, 255, 0, 255}, color.RGBA{255, 255, 255, 255}, color.RGBA{75, 75, 75, 255}}
	igd.Imat = igd.Imat.MakeIntMatrix(N_TilesX, N_TilesY)
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
	igd.Img = ebiten.NewImage(iX, iY)
	igd.Img.Fill(color.Black)

	igd.BoardMargin = CoordInts{X: iMargeX, Y: iMargeY}
	igd.Imat.DrawGridTiles(igd.Img, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)
	igd.Scale = 1.0
}

func (igd *IntegerGridManager) Draw(screen *ebiten.Image) {
	if igd.PFinder.IsStartInit {
		igd.Imat[igd.PFinder.StartPos.Y][igd.PFinder.StartPos.X] = 5
	}
	if igd.PFinder.IsEndInit {
		igd.Imat[igd.PFinder.EndPos.Y][igd.PFinder.EndPos.X] = 6
	}
	// go igd.Imat.DrawGridTiles(igd.Img, 4, 4, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(igd.Position.X-igd.BoardMargin.X), float64(igd.Position.Y-igd.BoardMargin.Y))
	ops.GeoM.Scale(igd.Scale, igd.Scale)
	screen.DrawImage(igd.Img, &ops)

	//igd.Imat.DrawGridTiles(screen, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)
	if igd.PFinder.IsStartInit {
		igd.Imat.DrawAGridTile(screen, igd.PFinder.StartPos, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{250, 250, 250, 255}, true)
	}
	if igd.PFinder.IsEndInit {
		igd.Imat.DrawAGridTile(screen, igd.PFinder.EndPos, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 50, 50, 255}, true)
	}
	if igd.PFinder.IsFullyInitialized {
		//igd.Imat.DrawAGridTile(screen, igd.PFinder.Cursor.Position, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 140, 50, 255}, false)

		if igd.PFinder.HasFalsePos {
			// igd.Imat.DrawAGridTile(screen, igd.PFinder.FalsePos, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{140, 50, 50, 255}, false)

			for _, j := range igd.PFinder.FalsePos {
				igd.Imat.DrawAGridTile(screen, j, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{140, 50, 50, 255}, false)
			}
			for _, x := range igd.PFinder.Moves {
				igd.Imat.DrawAGridTile(screen, x, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 125, 125, 255}, false)

			}

		}
		igd.DrawCursor(screen)
		// igd.Imat.DrawAGridTile_With_Line(screen, igd.PFinder.Cursor.Position, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{200, 150, 0, 255}, color.Black, 2.0, false)
		// igd.Imat.DrawAGridTile(screen, igd.PFinder.Cursor.Position, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{255, 15, 0, 1}, false) //{255,15,0,1}
	}
}

func (igd *IntegerGridManager) UpdateOnMouseEvent(Raw_Mouse_X, Raw_Mouse_Y int) {
	tempX, tempY := -1, -1
	if igd.FullColors {
		tempX, tempY = igd.Imat.ChangeValOnMouseEvent(Raw_Mouse_X, Raw_Mouse_Y, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.CycleStart, len(igd.Colors)-1, !(igd.PFinderStartSelect || igd.PFinderEndSelect))
	} else {
		tempX, tempY = igd.Imat.ChangeValOnMouseEvent(Raw_Mouse_X, Raw_Mouse_Y, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.CycleStart, igd.CycleEnd, !(igd.PFinderStartSelect || igd.PFinderEndSelect))
	}
	if tempX != -1 && tempY != -1 {
		if !igd.PFinderEndSelect && !igd.PFinderStartSelect && igd.SelectPoints {
			igd.LastPoint = CoordInts{tempX, tempY}
			if (!igd.LastPoint.IsEqualTo(CoordInts{-1, -1})) {
				igd.Coords = igd.Coords.PushToReturn(igd.LastPoint)
			}
		} else if igd.PFinderEndSelect {
			igd.PFinder.EndPos = CoordInts{tempX, tempY}
			//igd.Imat[tempY][tempX] = 5
			igd.PFinder.IsEndInit = true
			igd.PFinderEndSelect = false
		} else if igd.PFinderStartSelect {
			igd.PFinder.StartPos = CoordInts{tempX, tempY}
			//igd.Imat[tempY][tempX] = 6
			igd.PFinder.IsStartInit = true
			igd.PFinderStartSelect = false
		}
	}
}

func (igd *IntegerGridManager) UpdateOnMouseEvent2() {
	Raw_Mouse_X, Raw_Mouse_Y := ebiten.CursorPosition()
	tempX, tempY := -1, -1
	isOnTile := false
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton(0)) {
		go igd.Imat.DrawGridTiles(igd.Img, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)
		tempX, tempY, isOnTile = igd.Imat.GetCoordOfMouseEvent(Raw_Mouse_X, Raw_Mouse_Y, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)
		if igd.FullColors {

			_, _ = igd.Imat.ChangeValOnMouseEvent(Raw_Mouse_X, Raw_Mouse_Y, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.CycleStart, len(igd.Colors)-1, !(igd.PFinderStartSelect || igd.PFinderEndSelect))

		} else {
			_, _ = igd.Imat.ChangeValOnMouseEvent(Raw_Mouse_X, Raw_Mouse_Y, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.CycleStart, igd.CycleEnd, !(igd.PFinderStartSelect || igd.PFinderEndSelect))
		}
		if tempX != -1 && tempY != -1 {
			if !igd.PFinderEndSelect && !igd.PFinderStartSelect && igd.SelectPoints && isOnTile {
				igd.LastPoint = CoordInts{tempX, tempY}
				if (!igd.LastPoint.IsEqualTo(CoordInts{-1, -1})) {
					igd.Coords = igd.Coords.PushToReturn(igd.LastPoint)
				}
			} else if igd.PFinderEndSelect {
				igd.PFinder.EndPos = CoordInts{tempX, tempY}
				igd.Imat[tempY][tempX] = 5
				igd.PFinder.IsEndInit = true
				igd.PFinderEndSelect = false
			} else if igd.PFinderStartSelect {
				igd.PFinder.StartPos = CoordInts{tempX, tempY}
				igd.Imat[tempY][tempX] = 6
				igd.PFinder.IsStartInit = true
				igd.PFinderStartSelect = false
			}
		}
	}
}

func (igd *IntegerGridManager) DrawACircleOnClick(Raw_Mouse_X, Raw_Mouse_Y int, Radius int, valueIs int) {
	var center CoordInts
	var is_OnPoint bool
	center.X, center.Y, is_OnPoint = igd.Imat.GetCoordOfMouseEvent(Raw_Mouse_X, Raw_Mouse_Y, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)

	if is_OnPoint {
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

		fmt.Printf("0: %d %d VALUEIS %d\n", center.X, center.Y, valueIs)
		fmt.Printf("A: %d %d\n", temp_01A.X, temp_01A.Y)
		fmt.Printf("B: %d %d\n", temp_01B.X, temp_01B.Y)
		fmt.Printf("C: %d %d\n", temp_01C.X, temp_01C.Y)
		fmt.Printf("D: %d %d\n", temp_01D.X, temp_01D.Y)
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

func (igd *IntegerGridManager) DrawCoordsOnImat() {
	//igd.ClearImat()
	for _, c := range igd.Coords {
		if igd.Imat[c.Y][c.X] != 4 {
			igd.Imat[c.Y][c.X] = 2
		}
	}
}

func (igd *IntegerGridManager) Process() {
	//igd.ClearImat()
	igd.Coords = igd.Coords.RemoveDuplicates()
	temp := make(CoordList, len(igd.Coords))
	copy(temp, igd.Coords)
	var frustration bool = true
	for _, c := range igd.Coords {

		nn := CoordInts{X: c.X, Y: c.Y - 1}
		//ne := CoordInts{X: c.X + 1, Y: c.Y - 1}
		ee := CoordInts{X: c.X + 1, Y: c.Y}
		// se := CoordInts{X: c.X + 1, Y: c.Y + 1}
		ss := CoordInts{X: c.X, Y: c.Y + 1}
		// sw := CoordInts{X: c.X - 1, Y: c.Y + 1}
		ww := CoordInts{X: c.X - 1, Y: c.Y}
		// nw := CoordInts{X: c.X - 1, Y: c.Y - 1}
		// if igd.Imat.GetCoordVal(c) != 1 {

		// }
		buffer := 2
		if igd.Imat.IsValid_With_Constant_Buffer(nn, buffer) && igd.Imat.GetCoordVal(nn) != 1 { // && igd.Imat.GetCoordVal(ne) != 1 && igd.Imat.GetCoordVal(nw) != 1
			igd.Imat[c.Y][c.X] = 1
			if igd.Imat.IsValid(nn) {
				temp = temp.PushToReturn(nn)
			}
			temp, _ = temp.RemoveCoordFromList(c)
			frustration = false
		}

		if igd.Imat.IsValid_With_Constant_Buffer(ee, buffer) && igd.Imat.GetCoordVal(ee) != 1 { //&& igd.Imat.GetCoordVal(ne) != 1 && igd.Imat.GetCoordVal(se) != 1
			igd.Imat[c.Y][c.X] = 1
			temp = temp.PushToReturn(ee)
			temp, _ = temp.RemoveCoordFromList(c)
			frustration = false

		}
		if igd.Imat.IsValid_With_Constant_Buffer(ss, buffer) && igd.Imat.GetCoordVal(ss) != 1 { //&& igd.Imat.GetCoordVal(sw) != 1&& igd.Imat.GetCoordVal(se) != 1
			igd.Imat[c.Y][c.X] = 1
			temp = temp.PushToReturn(ss)
			temp, _ = temp.RemoveCoordFromList(c)
			frustration = false

		}
		if igd.Imat.IsValid_With_Constant_Buffer(ww, buffer) && igd.Imat.GetCoordVal(ww) != 1 { //&& igd.Imat.GetCoordVal(sw) != 1 && igd.Imat.GetCoordVal(nw) != 1
			igd.Imat[c.Y][c.X] = 1
			temp = temp.PushToReturn(ww)
			temp, _ = temp.RemoveCoordFromList(c)
			frustration = false

		}

		if frustration {
			igd.Fails++
		}
		//igd.DrawCoordsOnImat()
		//fmt.Printf("C:%2d,%2d\t D:%2d,%2d\t E:%2d,%2d\n-------\n", c.X, c.Y, d.X, d.Y, e.X, e.Y)
	}
	igd.Coords = temp
	//copy(igd.Coords, temp)
	igd.DrawCoordsOnImat()
}

func (igd *IntegerGridManager) Process2(setBools bool) {
	//igd.ClearImat()
	igd.Coords = igd.Coords.RemoveDuplicates()
	temp := make(CoordList, len(igd.Coords))
	copy(temp, igd.Coords)
	var frustration bool = true
	randInt := rand.Intn(len(igd.Coords))
	c := igd.Coords[randInt]
	nn := CoordInts{X: c.X, Y: c.Y - 1}
	ne := CoordInts{X: c.X + 1, Y: c.Y - 1}
	ee := CoordInts{X: c.X + 1, Y: c.Y}
	se := CoordInts{X: c.X + 1, Y: c.Y + 1}
	ss := CoordInts{X: c.X, Y: c.Y + 1}
	sw := CoordInts{X: c.X - 1, Y: c.Y + 1}
	ww := CoordInts{X: c.X - 1, Y: c.Y}
	nw := CoordInts{X: c.X - 1, Y: c.Y - 1}
	// if igd.Imat.GetCoordVal(c) != 1 {
	nEBool := false
	nWBool := false
	sEBool := false
	sWBool := false

	if setBools {
		sWBool = true
		nWBool = true
		nEBool = true
		nWBool = true
	} else {
		sWBool = igd.ValidatePointAgainstValue(sw, 1)
		sEBool = igd.ValidatePointAgainstValue(se, 1)
		nEBool = igd.ValidatePointAgainstValue(ne, 1)
		nWBool = igd.ValidatePointAgainstValue(nw, 1)
	}
	// }
	// randInt = rand.Intn(3)
	// switch(randInt){
	// case 0:
	// case 1:
	// case 2:
	// case 3:
	// default:
	// }
	if igd.Imat.IsValid(nn) && igd.Imat.GetCoordVal(nn) != 1 && nEBool && nWBool { // igd.Imat.GetCoordVal(ne) != 1 && igd.Imat.GetCoordVal(nw) != 1
		igd.Imat[c.Y][c.X] = 1
		if igd.Imat.IsValid(nn) {
			temp = temp.PushToReturn(nn)
		}
		temp, _ = temp.RemoveCoordFromList(c)
		frustration = false
	}

	if igd.Imat.IsValid(ee) && igd.Imat.GetCoordVal(ee) != 1 && nEBool && sEBool { //&& igd.Imat.GetCoordVal(ne) != 1 && igd.Imat.GetCoordVal(se) != 1
		igd.Imat[c.Y][c.X] = 1
		temp = temp.PushToReturn(ee)
		temp, _ = temp.RemoveCoordFromList(c)
		frustration = false

	}
	if igd.Imat.IsValid(ss) && igd.Imat.GetCoordVal(ss) != 1 && sEBool && sWBool { //&& igd.Imat.GetCoordVal(sw) != 1&& igd.Imat.GetCoordVal(se) != 1
		igd.Imat[c.Y][c.X] = 1
		temp = temp.PushToReturn(ss)
		temp, _ = temp.RemoveCoordFromList(c)
		frustration = false

	}
	if igd.Imat.IsValid(ww) && igd.Imat.GetCoordVal(ww) != 1 && sWBool && nWBool { //&& igd.Imat.GetCoordVal(sw) != 1 && igd.Imat.GetCoordVal(nw) != 1
		igd.Imat[c.Y][c.X] = 1
		temp = temp.PushToReturn(ww)
		temp, _ = temp.RemoveCoordFromList(c)
		frustration = false

	}

	if frustration {
		igd.Fails++
	}
	//igd.DrawCoordsOnImat()
	//fmt.Printf("C:%2d,%2d\t D:%2d,%2d\t E:%2d,%2d\n-------\n", c.X, c.Y, d.X, d.Y, e.X, e.Y)

	igd.Coords = temp
	//copy(igd.Coords, temp)
	igd.DrawCoordsOnImat()
}
func (igd *IntegerGridManager) ValidatePointAgainstValue(coord CoordInts, num int) bool {
	if igd.Imat.IsValid(coord) {
		return igd.Imat.GetCoordVal(coord) != num
	}
	return false
}
func (igd *IntegerGridManager) ValidatePointAgainstArrayOfValues(coord CoordInts, num []int) bool {
	ret := true
	if igd.Imat.IsValid(coord) {
		for _, x := range num {
			if igd.Imat.GetCoordVal(coord) == x {
				ret = false
			}
		}
	}
	return ret
}
func (igd *IntegerGridManager) ValidatePointsAgainstValuesDiag(coord CoordInts, num []int) (bool, bool, bool, bool) {
	// nn := CoordInts{X: coord.X, Y: coord.Y - 1}
	nePoint := CoordInts{X: coord.X + 1, Y: coord.Y - 1}
	// ee := CoordInts{X: coord.X + 1, Y: coord.Y}
	sePoint := CoordInts{X: coord.X + 1, Y: coord.Y + 1}
	// ss := CoordInts{X: coord.X, Y: coord.Y + 1}
	swPoint := CoordInts{X: coord.X - 1, Y: coord.Y + 1}
	// ww := CoordInts{X: coord.X - 1, Y: coord.Y}
	nwPoint := CoordInts{X: coord.X - 1, Y: coord.Y - 1}
	ne := igd.ValidatePointAgainstArrayOfValues(nePoint, num)
	se := igd.ValidatePointAgainstArrayOfValues(sePoint, num)
	sw := igd.ValidatePointAgainstArrayOfValues(swPoint, num)
	nw := igd.ValidatePointAgainstArrayOfValues(nwPoint, num)
	return ne, se, sw, nw
}
func (igd *IntegerGridManager) Process2b(maxTicks int) {
	if len(igd.Coords) > 0 {
		for range maxTicks {
			igd.Process2(false)
			// igd.Process3()
		}
	}
}
func (igd *IntegerGridManager) Process3b(maxTicks int, lims int, lim2 int, nums []int) { //, lim4, lim5 int
	// ,lim3 int, lim4, lim5 int,
	for i := range maxTicks {
		if len(igd.Coords) > 0 {
			igd.Process3(lims, lim2, nums)
			// igd.Process3()
			if igd.Fails > igd.FailsMax {
				break
			}
		}
		if i != 0 {
			if i%2 == 0 {
				igd.CullCoords(2, false, nums) //2 //lim4
			} else if i%5 == 0 {
				igd.CullCoords(4, false, nums) //4 //lim5
			}
		}
		// if igd.Fails > igd.FailsMax {
		// 	fmt.Printf("DEAD\n")
		// }
	}

}

func (igd *IntegerGridManager) Process3c(maxTicks int, lims int, lim2 int, nums []int) { //,lim4, lim5 int //, lim3, lim4, lim5 int
	min_Y, max_Y := 2, len(igd.Imat)-2
	min_X, max_X := 2, len(igd.Imat[0])-2
	//fmt.Printf("TEST\n")
	if igd.AlgorithmRunning {
		if len(igd.Coords) < 1 {
			//fmt.Printf("STARTING\n")
			rX, rY := rand.Intn(max_X-min_X)+min_X, rand.Intn(max_Y-min_Y)+min_Y
			igd.Coords = append(igd.Coords, CoordInts{X: rX, Y: rY})
			//igd.AlgorithmRunning = true
		}
		igd.Process3b(maxTicks, lims, lim2, nums) //lim4, lim5
		if igd.Fails > igd.FailsMax {
			fmt.Printf("FINISHED\n")
			igd.CullCoords(8, true, nums) //8
			go igd.Imat.DrawGridTiles(igd.Img, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)
			igd.AlgorithmRunning = false
		}
		// igd.CullCoords(2, false, []int{0, 2})
		// igd.CullCoords(2, true, []int{0, 2})
		// igd.CullCoords(2, false, []int{0, 2})
		// igd.CullCoords(2, false, []int{0, 2})
	}
	// else {
	// 	fmt.Printf("STARTING\n")

	// }
}

func (igd *IntegerGridManager) Process3(lims int, lim2 int, nums []int) {
	var randInt = 0
	igd.Coords = igd.Coords.RemoveDuplicates()
	if len(igd.Coords) > 0 {
		randInt = rand.Intn(len(igd.Coords))
	}
	temp := make(CoordList, len(igd.Coords))
	copy(temp, igd.Coords)
	var frustration bool = true
	var canDiag = false
	c := igd.Coords[randInt]
	if igd.SamplePoint(c, lims) {
		nn := CoordInts{X: c.X, Y: c.Y - 1}
		ne := CoordInts{X: c.X + 1, Y: c.Y - 1}
		ee := CoordInts{X: c.X + 1, Y: c.Y}
		se := CoordInts{X: c.X + 1, Y: c.Y + 1}
		ss := CoordInts{X: c.X, Y: c.Y + 1}
		sw := CoordInts{X: c.X - 1, Y: c.Y + 1}
		ww := CoordInts{X: c.X - 1, Y: c.Y}
		nw := CoordInts{X: c.X - 1, Y: c.Y - 1}
		// if igd.Imat.GetCoordVal(c) != 1 {
		nEBool := igd.ValidatePointAgainstValue(sw, 1)
		nWBool := igd.ValidatePointAgainstValue(se, 1)
		sEBool := igd.ValidatePointAgainstValue(ne, 1)
		sWBool := igd.ValidatePointAgainstValue(nw, 1)

		// }
		// randInt = rand.Intn(3)
		// switch(randInt){
		// case 0:
		// case 1:
		// case 2:
		// case 3:
		// default:
		// }
		if igd.Imat.IsValid(nn) && (igd.Imat.GetCoordVal(nn) != 1) && (igd.Imat.GetCoordVal(nn) != 4) && nEBool && nWBool { // igd.Imat.GetCoordVal(ne) != 1 && igd.Imat.GetCoordVal(nw) != 1
			igd.Imat[c.Y][c.X] = 1
			temp = temp.PushToReturn(nn)
			temp = append(temp, igd.AddAllToCoords(c, canDiag, nums)...)
			// temp = temp.RemoveDuplicates()
			temp, _ = temp.RemoveCoordFromList(c)
			frustration = false
		}

		if igd.Imat.IsValid(ee) && igd.Imat.GetCoordVal(ee) != 1 && (igd.Imat.GetCoordVal(ee) != 4) && nEBool && sEBool { //&& igd.Imat.GetCoordVal(ne) != 1 && igd.Imat.GetCoordVal(se) != 1
			igd.Imat[c.Y][c.X] = 1
			temp = temp.PushToReturn(ee)
			temp, _ = temp.RemoveCoordFromList(c)
			temp = append(temp, igd.AddAllToCoords(c, canDiag, nums)...)
			//temp = temp.RemoveDuplicates()
			frustration = false

		}
		if igd.Imat.IsValid(ss) && igd.Imat.GetCoordVal(ss) != 1 && (igd.Imat.GetCoordVal(ss) != 4) && sEBool && sWBool { //&& igd.Imat.GetCoordVal(sw) != 1&& igd.Imat.GetCoordVal(se) != 1
			igd.Imat[c.Y][c.X] = 1
			temp = temp.PushToReturn(ss)
			temp, _ = temp.RemoveCoordFromList(c)
			temp = append(temp, igd.AddAllToCoords(c, canDiag, nums)...)
			// temp = temp.RemoveDuplicates()
			frustration = false

		}
		if igd.Imat.IsValid(ww) && igd.Imat.GetCoordVal(ww) != 1 && (igd.Imat.GetCoordVal(ww) != 4) && sWBool && nWBool { //&& igd.Imat.GetCoordVal(sw) != 1 && igd.Imat.GetCoordVal(nw) != 1
			igd.Imat[c.Y][c.X] = 1
			temp = temp.PushToReturn(ww)
			temp, _ = temp.RemoveCoordFromList(c)
			// temp2 := igd.AddAllToCoords(c)
			// for _, c := range temp2 {
			// 	temp = append(temp, c)
			// }
			temp = append(temp, igd.AddAllToCoords(c, canDiag, nums)...)
			// temp = temp.RemoveDuplicates()
			frustration = false

		}
		temp = temp.RemoveDuplicates()

		igd.Coords = temp
	} else {
		//temp, _ = temp.RemoveCoordFromList(c) <----- MIGHT BE AN EDGE CASE BE AWARE THIS IS NOT TO BE ERASED
		// igd.Imat[c.Y][c.X] = 4
		igd.CullCoords(lim2, canDiag, nums)
		//igd.CullCoords(lim2, true, nums)
		//temp = igd.Coords
	}
	if frustration {
		igd.Fails++
	} else {
		igd.Fails = 0

	}

	//igd.DrawCoordsOnImat()
	//fmt.Printf("C:%2d,%2d\t D:%2d,%2d\t E:%2d,%2d\n-------\n", c.X, c.Y, d.X, d.Y, e.X, e.Y)
	// temp = temp.RemoveDuplicates()

	// igd.Coords = temp
	//copy(igd.Coords, temp)
	igd.DrawCoordsOnImat()
	// fmt.Printf("SIZE: %3d SIZE %3d\n", len(igd.Coords), len(temp))
}

/* should cull igd.Coords;
 */
// func (igd *IntegerGridManager) SamplePoint(cord CoordInts, min int) bool {
// 	num := igd.GetNumberOfValidNeighbors(cord, true, []int{0, 2})
// 	// if  {
// 	// 	igd.Coords.RemoveCoordFromList(cord)
// 	// }
// 	// fmt.Printf(" NUM : %d\n", num)
// 	return (8 - num) < min
// }

func (igd *IntegerGridManager) SamplePoint(cord CoordInts, min int) bool {
	num := igd.GetNumberOfValidNeighbors(cord, true, []int{1, 4})
	// if  {
	// 	igd.Coords.RemoveCoordFromList(cord)
	// }
	// fmt.Printf(" NUM : %d\n", num)
	return (num) < min
}

func (igd *IntegerGridManager) CullCoords(limit int, canDiag bool, nums []int) {
	temp := make(CoordList, len(igd.Coords))
	copy(temp, igd.Coords)
	for _, c := range igd.Coords {
		//b := igd.GetNumberOfValidNeighbors(c, []int{1, 4})
		q := igd.GetNumberOfValidNeighbors(c, !canDiag, nums)
		if q < limit {
			temp, _ = temp.RemoveCoordFromList(c)
			igd.Imat[c.Y][c.X] = 4
			//fmt.Printf("LIMIT COORDCULLING %d %d %d %d\n", b, q, c.X, c.Y)
		}
		//additional culling:
		// if !canDiag {
		// 	// nn := igd.Imat.GetCoordVal(CoordInts{X: c.X, Y: c.Y - 1})

		// 	// ee := igd.Imat.GetCoordVal(CoordInts{X: c.X + 1, Y: c.Y})

		// 	// ss := igd.Imat.GetCoordVal(CoordInts{X: c.X, Y: c.Y + 1})

		// 	// ww := igd.Imat.GetCoordVal(CoordInts{X: c.X - 1, Y: c.Y})
		// 	// if(nn==1 && ss==1 ||)
		// 	//if(igd.IMat.)
		// }
		//fmt.Printf("LIMIT COORDCULLING BEE %d %d %d %d\n", b, q, c.X, c.Y)
	}
	//fmt.Printf("\n\n")
	igd.Coords = temp
}
func (igd *IntegerGridManager) AddAllToCoords(coord CoordInts, canDiag bool, nums []int) CoordList {
	nn, ne, ee, se, ss, sw, ww, nw := igd.ValidateAllNeighbors(coord, nums) //[]int{0, 2}
	temp := make(CoordList, 0)
	if !nn {
		temp = append(temp, CoordInts{X: coord.X, Y: coord.Y - 1})
	} else {
		//fmt.Printf("ADD ALL COORDS: --- nn: %d %d %d\n", coord.X, coord.Y-1, igd.Imat[coord.Y-1][coord.X])
	}
	if !ne && canDiag {
		temp = append(temp, CoordInts{X: coord.X + 1, Y: coord.Y - 1})
	} else {
		//fmt.Printf("ADD ALL COORDS: --- ne: %d %d %d\n", coord.X+1, coord.Y-1, igd.Imat[coord.Y-1][coord.X+1])
	}

	if !ee {
		temp = append(temp, CoordInts{X: coord.X + 1, Y: coord.Y})
	} else {
		//fmt.Printf("ADD ALL COORDS: --- ee: %d %d %d\n", coord.X+1, coord.Y, igd.Imat[coord.Y][coord.X+1])
	}
	if !se && canDiag {
		temp = append(temp, CoordInts{X: coord.X + 1, Y: coord.Y + 1})
	} else {
		//fmt.Printf("ADD ALL COORDS: --- se: %d %d %d\n", coord.X+1, coord.Y+1, igd.Imat[coord.Y+1][coord.X+1])
	}
	if !ss {
		temp = append(temp, CoordInts{X: coord.X, Y: coord.Y + 1})
	} else {
		//fmt.Printf("ADD ALL COORDS: --- ss: %d %d %d\n", coord.X, coord.Y+1, igd.Imat[coord.Y+1][coord.X])
	}
	if !sw && canDiag {
		temp = append(temp, CoordInts{X: coord.X - 1, Y: coord.Y + 1})
	} else {
		//fmt.Printf("ADD ALL COORDS: --- sW: %d %d %d\n", coord.X-1, coord.Y+1, igd.Imat[coord.Y+1][coord.X-1])
	}
	if !ww {
		temp = append(temp, CoordInts{X: coord.X - 1, Y: coord.Y})
	} else {
		//fmt.Printf("ADD ALL COORDS: --- WW: %d %d %d\n", coord.X-1, coord.Y, igd.Imat[coord.Y][coord.X-1])
	}
	if !nw && canDiag {
		temp = append(temp, CoordInts{X: coord.X - 1, Y: coord.Y - 1})
	} else {
		//fmt.Printf("ADD ALL COORDS: --- NW: %d %d %d\n", coord.X-1, coord.Y-1, igd.Imat[coord.Y-1][coord.X-1])
	}
	//fmt.Printf("\n==============================\n")
	return temp
}

func (igd *IntegerGridManager) ValidateAllNeighbors(coord CoordInts, numVals []int) (nn bool, ne bool, ee bool, se bool, ss bool, sw bool, ww bool, nw bool) {
	// nnPoint :=CoordInts{X: coord.X, Y: coord.Y - 1}
	// nePoint := CoordInts{X: coord.X + 1, Y: coord.Y - 1}
	// eePoint := CoordInts{X: coord.X + 1, Y: coord.Y}
	// sePoint := CoordInts{X: coord.X + 1, Y: coord.Y + 1}
	// ssPoint := CoordInts{X: coord.X, Y: coord.Y + 1}
	// swPoint := CoordInts{X: coord.X - 1, Y: coord.Y + 1}
	// wwPoint := CoordInts{X: coord.X - 1, Y: coord.Y}
	// nwPoint := CoordInts{X: coord.X - 1, Y: coord.Y - 1}
	nn = igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X, Y: coord.Y - 1}, numVals)
	ne = igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X + 1, Y: coord.Y - 1}, numVals)
	ee = igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X + 1, Y: coord.Y}, numVals)
	se = igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X + 1, Y: coord.Y + 1}, numVals)
	ss = igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X, Y: coord.Y + 1}, numVals)
	sw = igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X - 1, Y: coord.Y + 1}, numVals)
	ww = igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X - 1, Y: coord.Y}, numVals)
	nw = igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X - 1, Y: coord.Y - 1}, numVals)
	return nn, ne, ee, se, ss, sw, ww, nw
}
func (igd *IntegerGridManager) GetNumberOfValidNeighbors(coord CoordInts, canDiag bool, numVals []int) int {
	var retInt = 0
	if (!igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X, Y: coord.Y - 1}, numVals)) {
		// fmt.Printf("GNofVN: HAS North\n")
		retInt++
	}
	if (!igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X + 1, Y: coord.Y - 1}, numVals)) && canDiag {
		// fmt.Printf("GNofVN: HAS NorthEAST\n")
		retInt++
	}
	if (!igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X + 1, Y: coord.Y}, numVals)) {
		// fmt.Printf("GNofVN: HAS EAST\n")
		retInt++
	}
	if (!igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X + 1, Y: coord.Y + 1}, numVals)) && canDiag {
		// fmt.Printf("GNofVN: HAS SouthEast\n")
		retInt++
	}
	if (!igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X, Y: coord.Y + 1}, numVals)) {
		// fmt.Printf("GNofVN: HAS south\n")

		retInt++
	}
	if (!igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X - 1, Y: coord.Y + 1}, numVals)) && canDiag {
		// fmt.Printf("GNofVN: HAS southwest\n")

		retInt++
	}
	if (!igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X - 1, Y: coord.Y}, numVals)) {
		// fmt.Printf("GNofVN: HAS west\n")

		retInt++
	}
	if (!igd.ValidatePointAgainstArrayOfValues(CoordInts{X: coord.X - 1, Y: coord.Y - 1}, numVals)) && canDiag {
		// fmt.Printf("GNofVN: HAS Northwest\n")

		retInt++
	}
	return retInt
}
