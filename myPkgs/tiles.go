package mypkgs

import (
	"fmt"
	"math/rand"

	//"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
func (imat IntMatrix) DrawAGridTile(screen *ebiten.Image, coord CoordInts, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int, clr color.Color, aa bool) {
	// test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	// test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	vector.DrawFilledRect(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), float32(tileW), float32(tileH), clr, aa)
	vector.StrokeRect(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), float32(tileW), float32(tileH), 2.0, color.Black, aa)
}
func (imat IntMatrix) DrawGridTile(screen *ebiten.Image, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int, colors []color.Color) {
	// colorA := color.RGBA{255, 0, 0, 255}
	// colorB := color.RGBA{0, 0, 255, 255}
	// colorC := color.RGBA{0, 255, 0, 255}
	// colorD := color.RGBA{0, 255, 255, 255}
	// colorE := color.RGBA{255, 255, 0, 255}
	test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	for y, _ := range imat {
		for x, b := range imat[y] {
			vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colors[b], false)
			// switch b {
			// case 0:
			// 	vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorA, false)

			// case 1:
			// 	vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorB, false)
			// case 2:
			// 	vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorC, false)
			// case 3:
			// 	vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorD, false)
			// case 4:
			// 	vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorE, false)
			// }
			vector.StrokeRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), 2.0, color.Black, false)
		}
	}
	//vector.StrokeRect(screen, float32(OffsetX-0), float32(OffsetY-0), float32(test1X-0), float32(test1Y+0), 2.0, color.RGBA{210, 153, 100, 255}, true) //0, 179, 100, 255
	vector.StrokeRect(screen, float32(OffsetX-3), float32(OffsetY-3), float32(test1X-OffsetX-GapX+6), float32(test1Y-OffsetY-GapY+6), 4.0, color.RGBA{0, 50, 50, 255}, true)
	//vector.StrokeRect(screen, float32(OffsetX-0), float32(OffsetY-0), float32(test1X-OffsetX), float32(test1Y-OffsetY), 2.0, color.RGBA{0, 253, 100, 255}, true)
}

func (imat IntMatrix) ChangeValOnMouseEvent(Raw_Mouse_X int, Raw_Mouse_Y int, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int, cycleStart int, cycleEnd int, makeChange bool) (int, int) {

	test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	var mXi = -1
	var mYi = -1
	if (Raw_Mouse_X > OffsetX && Raw_Mouse_X < test1X-GapX) && (Raw_Mouse_Y > OffsetY && Raw_Mouse_Y < test1Y-GapY) {

		var mX = float32(Raw_Mouse_X-OffsetX) / float32(tileW+GapX) //float32(test1X)
		var mY = float32(Raw_Mouse_Y-OffsetY) / float32(tileH+GapY) //float32(test1Y)
		mXi = int(mX)
		mYi = int(mY)
		// fmt.Printf("IN THE BOX\nRM_: %d,%d\n", Raw_Mouse_X, Raw_Mouse_Y)
		// fmt.Printf("test %d,%d\n", test1Y, test1X)
		//----
		// fmt.Printf("RMO: %d, %d\nm_f: %f %f\n", Raw_Mouse_X-OffsetX, Raw_Mouse_Y-OffsetY, mX, mY)
		// fmt.Printf("M_I: %d, %d\nm_i: %d,%d\n", (Raw_Mouse_X-OffsetX)/(tileH+GapX), (Raw_Mouse_Y-OffsetY)/(tileH+GapY), mXi, mYi)

		// fmt.Printf("A %d,%d\n", (tileW*mXi)+(mXi*GapX), (tileH*mYi)+(mYi*GapY))                           //inner bounds of each rectangle
		// fmt.Printf("B %d,%d\n", (tileW*mXi)+(mXi*GapX)+tileW, (tileH*mYi)+(mYi*GapY)+tileH)               //outer bounds of each rectangle;
		// fmt.Printf("C %d,%d\n", (tileW*mXi)+(mXi*GapX)+(tileW+GapX), (tileH*mYi)+(mYi*GapY)+(tileH+GapY)) //gaps ending\
		//-----------------------------------------------
		mXo, mYo := (Raw_Mouse_X - OffsetX), (Raw_Mouse_Y - OffsetY)
		mXi_01 := (tileW * mXi) + (mXi * GapX)
		mYi_01 := (tileH * mYi) + (mYi * GapY)

		mXi_02 := (tileW * mXi) + (mXi * GapX) + tileW
		mYi_02 := (tileH * mYi) + (mYi * GapY) + tileH
		//--------------------------------------------------------------------
		// fmt.Printf("m_o: %d,%d\n", mXo, mYo)
		// fmt.Printf("LENS: %d %d\n", len(imat), len(imat[0]))
		// fmt.Printf("TESTS:: %d %d \n", test1Y, test1X)
		// fmt.Printf("A: %d,%d\n", mXi_01, mYi_01)                                                               //inner bounds of each rectangle
		// fmt.Printf("B: %d,%d\n", mYi_02, mYi_02)                                                               //outer bounds of each rectangle;
		// fmt.Printf("C: %d,%d\n\n\n", (tileW*mXi)+(mXi*GapX)+(tileW+GapX), (tileH*mYi)+(mYi*GapY)+(tileH+GapY)) //gaps ending\
		if (((mXo) > (mXi_01)) && (mXo) < mXi_02) && ((mYo) > (mYi_01) && mYo < mYi_02) && makeChange {
			//change the coord;
			if imat[mYi][mXi] < cycleEnd {
				imat[mYi][mXi] += 1
			} else {
				imat[mYi][mXi] = cycleStart
			}
		}
		// if (((Raw_Mouse_X - OffsetX) > ((tileW * mXi) + (mXi * GapX))) && (Raw_Mouse_X-OffsetX) < mXi_02) && ((Raw_Mouse_Y-OffsetY) > ((tileH*mYi)+(mYi*GapY)) && (Raw_Mouse_Y-OffsetY) < mYi_02) {
		// 	//change the coord;
		// 	if imat[mYi][mXi] < cycleEnd {
		// 		imat[mYi][mXi] += 1
		// 	} else {
		// 		imat[mYi][mXi] = cycleStart
		// 	}
		// }
	}
	// else {
	// 	fmt.Printf("NOT IN THE BOX\nRM_: %d,%d\n", Raw_Mouse_X, Raw_Mouse_Y)
	// 	fmt.Printf("test %d,%d\n", test1X, test1Y)
	// }

	// var mX = Raw_Mouse_X - OffsetX
	// var mY = Raw_Mouse_Y - OffsetY
	return mXi, mYi
}
func (imat IntMatrix) GetCoordVal(cord CoordInts) int {
	//fmt.Printf("GET COORD VAL %d %d\n\n", cord.X, cord.Y)
	return imat[cord.Y][cord.X]

}

func (imat IntMatrix) IsValid(cord CoordInts) bool {
	if (cord.X > -1 && cord.X < len(imat[0])) && (cord.Y > -1 && cord.Y < len(imat)) {
		return true
	}
	return false
}
func (imat IntMatrix) IsValid_With_Constant_Buffer(cord CoordInts, buffer int) bool {
	if (cord.X > -1+buffer && cord.X < len(imat[0])-buffer) && (cord.Y > -1+buffer && cord.Y < len(imat)-buffer) {
		return true
	}
	return false
}
func (imat IntMatrix) IsValid_WithDir_Buffer(cord CoordInts, buffer [4]int) bool {
	if (cord.X > -1+buffer[3] && cord.X < len(imat[0])-buffer[1]) && (cord.Y > -1+buffer[0] && cord.Y < len(imat)-buffer[2]) {
		return true
	}
	return false
}

type CoordList []CoordInts

func (cord CoordList) ToString() string {
	retStrng := "COORDAR:\n"
	retStrng += fmt.Sprintf("--SIZE: %d\n", len(cord))
	return retStrng
}

func (cord CoordList) PushToReturn(coord CoordInts) CoordList {
	temp := append(cord, coord)
	return temp
}
func (cord CoordList) PopFromFront() (CoordInts, CoordList) {
	temp := cord[0]
	temp2 := cord.RemovePointFromList(0)
	return temp, temp2
}
func (cord CoordList) PopFromBack() (CoordInts, CoordList) {
	temp := cord[len(cord)-1]
	temp2 := cord.RemovePointFromList(len(cord) - 1)
	return temp, temp2
}

func (coord CoordList) ToCoordArray() []CoordInts {
	outAr := make([]CoordInts, len(coord))
	copy(outAr, coord)
	return outAr
}
func (coord CoordList) FromCoordArray(c []CoordInts) CoordList {
	outList := make(CoordList, len(c))
	copy(outList, c)
	return outList
}
func (cord CoordList) RemoveCoordFromList(coord CoordInts) (CoordList, bool) {
	temp := make(CoordList, 0)
	isThere := false
	for _, c := range cord {
		if !c.IsEqualTo(coord) {
			temp = append(temp, c)
		}
	}
	return temp, isThere
}
func (cord CoordList) RemovePointFromList(num int) CoordList {
	temp := make(CoordList, 0)
	for i, _ := range cord {
		if i != num {
			temp = append(temp, cord[i])
		}
	}
	return temp
}
func (cord CoordList) CountInstances(coord CoordInts) int {
	temp := 0
	for _, c := range cord {
		if c.IsEqualTo(coord) {
			temp++
		}
	}
	return temp
}

func (cord CoordList) PrintCordArray() {
	fmt.Print("\n\n------------------------\n")
	for i, c := range cord {
		fmt.Printf("%2d: {%3d %3d}", i, c.X, c.Y)
		// if i%1 == 0 {
		// 	fmt.Print("\n")
		// } else {
		// 	fmt.Print("\t")
		// }
		fmt.Print("\n")

	}
	fmt.Print("\n------------------------\n")
}

func (cord CoordList) SortDescOnX() CoordList {
	temp := make([]CoordInts, len(cord))
	copy(temp, cord)
	var tempcord CoordInts
	if len(temp) > 1 {
		// var halfTemp = (len(temp) / 2)
		for range temp {
			for i := 1; i < (len(temp)); i++ {
				// q := halfTemp + i
				if temp[i].X > temp[i-1].X {
					tempcord = temp[i]
					temp[i] = temp[i-1]
					temp[i-1] = tempcord
				} else if temp[i].X == temp[i-1].X {
					if temp[i].Y > temp[i-1].Y {
						tempcord = temp[i]
						temp[i] = temp[i-1]
						temp[i-1] = tempcord
					}
				}
			}
		}
	}
	return temp
}

/*
CoordList.RemoveDuplicates
this should be done;
this will remove duplicates;
*/
func (cord CoordList) RemoveDuplicates() CoordList {
	temp := make(CoordList, len(cord))
	copy(temp, cord)
	temp = temp.SortDescOnX()
	for i := 1; i < len(temp); i++ {
		if temp[i].IsEqualTo(temp[i-1]) {
			temp[i-1] = CoordInts{X: -1, Y: -1}
		}
	}
	temp, _ = temp.RemoveCoordFromList(CoordInts{X: -1, Y: -1})
	// temp2 := make(CoordList, 0)
	// for _, c := range cord {

	// }
	return temp
}

type IntegerGridManager struct {
	Imat               IntMatrix
	Coords             CoordList
	Tile_Size          CoordInts
	Margin             CoordInts
	Position           CoordInts
	CycleStart         int
	CycleEnd           int
	Colors             []color.Color
	FullColors         bool
	LastPoint          CoordInts //the last point clicked
	Fails              int
	AlgorithmRunning   bool
	FailsMax           int
	PFinder            Pathfinding
	PFinderEndSelect   bool
	PFinderStartSelect bool
}

func (igd *IntegerGridManager) Draw(screen *ebiten.Image) {
	if igd.PFinder.IsStartInit {
		igd.Imat[igd.PFinder.StartPos.Y][igd.PFinder.StartPos.X] = 5
	}
	if igd.PFinder.IsEndInit {
		igd.Imat[igd.PFinder.EndPos.Y][igd.PFinder.EndPos.X] = 6
	}
	igd.Imat.DrawGridTile(screen, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, igd.Colors)
	if igd.PFinder.IsFullyInitialized {
		igd.Imat.DrawAGridTile(screen, igd.PFinder.Cursor.Position, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 140, 50, 255}, false)
		if igd.PFinder.HasFalsePos {
			// igd.Imat.DrawAGridTile(screen, igd.PFinder.FalsePos, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{140, 50, 50, 255}, false)

			for _, j := range igd.PFinder.FalsePos {
				igd.Imat.DrawAGridTile(screen, j, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{140, 50, 50, 255}, false)
			}
			for _, x := range igd.PFinder.Moves {
				igd.Imat.DrawAGridTile(screen, x, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{50, 125, 125, 255}, false)

			}
		}
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
		if !igd.PFinderEndSelect && !igd.PFinderStartSelect {
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
func (igd *IntegerGridManager) Init(N_TilesX, N_TilesY int, TSizeX, TSizeY int, PosX, PosY int, MargX, MargY int) {
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
func (igd *IntegerGridManager) Process3b(maxTicks int, lims int, lim2 int, nums []int) {

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
				igd.CullCoords(2, false, []int{0, 2})
			} else if i%5 == 0 {
				igd.CullCoords(4, false, []int{0, 2})
			}
		}
		// if igd.Fails > igd.FailsMax {
		// 	fmt.Printf("DEAD\n")
		// }
	}

}

func (igd *IntegerGridManager) Process3c(maxTicks int, lims int, lim2 int, nums []int) {
	if igd.AlgorithmRunning {
		igd.Process3b(maxTicks, lims, lim2, nums)
		if igd.Fails > igd.FailsMax {
			fmt.Printf("FINISHED\n")
			igd.CullCoords(8, true, nums)
			igd.AlgorithmRunning = false
		}
		// igd.CullCoords(2, false, []int{0, 2})
		// igd.CullCoords(2, true, []int{0, 2})
		// igd.CullCoords(2, false, []int{0, 2})
		// igd.CullCoords(2, false, []int{0, 2})
	}
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
		temp, _ = temp.RemoveCoordFromList(c)
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
