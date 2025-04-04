package mypkgs

import (
	"fmt"
	"image/color"

	// "golang.org/x/exp/shiny/screen"
	"github.com/hajimehoshi/ebiten/v2"
)

type Pathfinding struct {
	StartPos           CoordInts
	EndPos             CoordInts
	IsActive           bool
	IsStartInit        bool
	IsEndInit          bool
	IsFullyInitialized bool
	Nodes              CoordList
	Color              color.Color
	SpriteDim          CoordInts
	Cursor             Cell
	FalsePos           CoordList //open list
	ClosedList         CoordList //closed list
	Moves              CoordList
	HasFalsePos        bool
}

// func (igd *Pathfinding) Tick_DownFalseposlane() {

// }

func (pFndr *Pathfinding) ToString() string {
	outstrng := "PATHFINDING:\n"
	outstrng += fmt.Sprintf("\n %6s %3d, %3d %t\n %6s %3d, %3d %t\n", "START:", pFndr.StartPos.X, pFndr.StartPos.Y, pFndr.IsStartInit, "END:", pFndr.EndPos.X, pFndr.EndPos.Y, pFndr.IsEndInit)
	outstrng += fmt.Sprintf("FULLY INIT: %t\n", pFndr.IsFullyInitialized)
	if pFndr.HasFalsePos {
		outstrng += fmt.Sprintf("%7s: %3d,%3d\n", "CURSOR", pFndr.Cursor.Position.X, pFndr.Cursor.Position.Y)
		outstrng += fmt.Sprintf("%7s: %3d,%3d\n", "END", pFndr.EndPos.X, pFndr.EndPos.Y)
		xx, yy := pFndr.Cursor.Position.GetDifferenceInInts(pFndr.EndPos)
		outstrng += fmt.Sprintf("%7s: %3d,%3d %d\n", "DIF.", xx, yy, GetDiffer(pFndr.Cursor.Position, pFndr.EndPos))
	}
	outstrng += fmt.Sprintf("FalsePos: %t ,len: %d\n", pFndr.HasFalsePos, len(pFndr.FalsePos))
	outstrng += fmt.Sprintf("MOVES:%d\n", len(pFndr.Moves))
	return outstrng
}
func (pFndr *Pathfinding) PrintString() {
	fmt.Printf("%s", pFndr.ToString())
}
func (igd *IntegerGridManager) DrawPathfinder(screen *ebiten.Image) {

}
func (igd *IntegerGridManager) RESETPathfinder() {
	if igd.PFinder.IsFullyInitialized {
		igd.Imat[igd.PFinder.EndPos.Y][igd.PFinder.EndPos.X] = 1
		igd.Imat[igd.PFinder.StartPos.Y][igd.PFinder.StartPos.X] = 1
		igd.PFinder.IsEndInit = false
		igd.PFinder.IsStartInit = false
		igd.PFinder.FalsePos = make(CoordList, 0)
		igd.PFinder.Moves = make(CoordList, 0)
		igd.PFinder.HasFalsePos = false
		igd.PFinder.IsFullyInitialized = false
		igd.PFinder.EndPos = CoordInts{X: -1, Y: -1}
		igd.PFinder.StartPos = CoordInts{X: -1, Y: -1}
	}
}

// func (igd *IntegerGridManager) SelectPathfinderStart() {

// }
// func (igd *IntegerGridManager) SelectPathfinderEnd() {

// }

func (igd *IntegerGridManager) PathfindingProcess() {
	fmt.Printf("\n\nPATHFINDING PROCESS\n\n")
	if igd.PFinder.IsStartInit && igd.PFinder.IsEndInit && !igd.PFinder.IsFullyInitialized {
		igd.PFinder.IsFullyInitialized = true
		fmt.Printf("\nINITIALIZED AND READY\n")
		igd.PFinder.Cursor.InitP(igd.PFinder.StartPos, igd.PFinder.EndPos, igd.Imat)
	}

	if igd.PFinder.IsFullyInitialized {
		//get vector to the target area;
		//igd.SLOPEMOVE(2, []int{0, 2, 3, 4})'
		// igd.PFinder.FalsePos = BresenhamLine(igd.PFinder.StartPos, igd.PFinder.EndPos)
		// igd.PFinder.HasFalsePos = true

	}
}
func (igd *IntegerGridManager) PFindr_DrawSlope() {
	if igd.PFinder.IsFullyInitialized {
		igd.SLOPEMOVE(1, []int{0, 2, 3, 4})
		igd.PFinder.HasFalsePos = true
	}
}
func (igd *IntegerGridManager) PFindr_DrawManhattan2(nums []int) {
	if igd.PFinder.IsFullyInitialized {
		var temper bool
		if !igd.PFinder.HasFalsePos {
			igd.PFinder.FalsePos, temper = ManhattanDistanceCulling(igd.PFinder.StartPos, igd.PFinder.EndPos, true, igd.Imat, nums)
			igd.PFinder.HasFalsePos = true
			if temper {
				fmt.Printf("HAH HIT A WALL\n")
			}
		} else {
			// igd.PFinder.Cursor.Position = igd.PFinder.FalsePos[igd.PFinder.Cursor.ticker]
			temp := igd.PFinder.FalsePos[igd.PFinder.Cursor.ticker]
			if a, c := IntArrayContains_giveMeWhat(nums, igd.Imat.GetCoordVal(temp)); a {
				fmt.Printf("HAS WALL: %d at %d , %d\n", c, temp.X, temp.Y)
			} else {
				igd.PFinder.Cursor.Position = temp
				igd.PFinder.Moves = append(igd.PFinder.Moves, temp)
				if igd.PFinder.Cursor.ticker < len(igd.PFinder.FalsePos)-1 {
					igd.PFinder.Cursor.ticker++
				}
			}

		}
	}

}
func (igd *IntegerGridManager) PFindr_DrawManhattan() {
	if igd.PFinder.IsFullyInitialized {
		yFirst := true
		// yLock := false
		// xLock := false
		vX, vY := igd.PFinder.Cursor.Position.GetDifferenceInInts(igd.PFinder.EndPos)
		var dirs = CoordInts{0, 0}
		var vectrex = CoordInts{X: 0, Y: 0}
		if vY == 0 {
			//-----
			fmt.Printf("Y== 0\n")
			dirs.X = 0
		} else if vY > 0 {
			vectrex.Y = vY
			dirs.Y = 2
		} else {
			vectrex.Y = vY * -1
			dirs.Y = 0
		}
		if vX == 0 {
			//-----
			dirs.X = 0
		} else if vX > 0 {
			vectrex.X = vX
			dirs.X = 1
		} else {
			vectrex.X = vX * -1
			dirs.X = 3
		}

		if yFirst {
			if igd.PFinder.Cursor.Position.Y != igd.PFinder.EndPos.Y {
				igd.MoveCursorSteps(CoordInts{dirs.Y, 1}, 2, []int{0, 2, 3, 4})
			} else {

				fmt.Printf("DELTA Y== 0\n")
				if igd.PFinder.Cursor.Position.X != igd.PFinder.EndPos.X {

					igd.MoveCursorSteps(CoordInts{dirs.X, 1}, 2, []int{0, 2, 3, 4})
				} else {
					fmt.Printf("DELTA X== 0\n")
				}
			}

		}

		// tempPos := MoveModifierCoords(igd.PFinder.Cursor.Position, Vect.X, 0)
		// if len(igd.PFinder.Nodes) > 0 {
		// 	if igd.PFinder.Nodes.CountInstances(tempPos) == 0 {
		// 		igd.PFinder.Nodes = append(igd.PFinder.Nodes, tempPos)
		// 	}
		// }
	}
}

func (igd *IntegerGridManager) MoveCursorSteps(Vect CoordInts, steps int, walls []int) (int, bool) {
	var tempCurs Cell = igd.PFinder.Cursor
	tempPos := MoveModifierCoords(tempCurs.Position, Vect.X, 0)
	var valer int

	for i := 0; (i < len(igd.Imat)) && (i < steps); i++ {
		tempPos = MoveModifierCoords(tempCurs.Position, Vect.X, i)
		fmt.Printf("MOVE: %d - POINT: %d %d Val Here: %d\n", i, tempPos.X, tempPos.Y, igd.Imat.GetCoordVal(tempPos))
		if a, c := IntArrayContains_giveMeWhat(walls, igd.Imat.GetCoordVal(tempPos)); a {
			valer = i
			fmt.Printf("HAS WALL: %d at %d , %d\n", c, tempPos.X, tempPos.Y)
			igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, tempPos)
			tempPos = MoveModifierCoords(tempCurs.Position, Vect.X, i-1)

			break
		} else {
			igd.PFinder.Moves = append(igd.PFinder.Moves, tempPos)
		}
	}
	igd.PFinder.HasFalsePos = true
	igd.PFinder.Cursor.Position = tempPos
	fmt.Printf(" MOVE CURSOR AROUND: Dir: %d  Mag %d  Actual: %d \n\n", Vect.X, Vect.Y, valer)
	return valer, true
}

func (igd *IntegerGridManager) PFindr_DrawBresenHamLine(nums []int) {
	if igd.PFinder.IsFullyInitialized {
		var temper bool
		if !igd.PFinder.HasFalsePos {
			igd.PFinder.FalsePos, temper = BresenhamLine_CullAfterImpact(igd.PFinder.StartPos, igd.PFinder.EndPos, igd.Imat, nums)
			igd.PFinder.HasFalsePos = true
			if temper {
				fmt.Printf("HAH HIT A WALL\n")
			}
		} else {
			// igd.PFinder.Cursor.Position = igd.PFinder.FalsePos[igd.PFinder.Cursor.ticker]
			temp := igd.PFinder.FalsePos[igd.PFinder.Cursor.ticker]
			if a, c := IntArrayContains_giveMeWhat(nums, igd.Imat.GetCoordVal(temp)); a {
				fmt.Printf("HAS WALL: %d at %d , %d\n", c, temp.X, temp.Y)
			} else {
				igd.PFinder.Cursor.Position = temp
				igd.PFinder.Moves = append(igd.PFinder.Moves, temp)
				if igd.PFinder.Cursor.ticker < len(igd.PFinder.FalsePos)-1 {
					igd.PFinder.Cursor.ticker++
				}
			}

		}
	}
}

func MoveModifier(startX, startY, dir, mag int) (int, int) {
	endX, endY := startX, startY
	switch dir {
	case 0:
		endY -= mag
	case 1:
		endX += mag
	case 2:
		endY += mag
	case 3:
		endX -= mag
	}
	return endX, endY
}
func MoveModifierCoords(start CoordInts, dir, mag int) CoordInts {
	ender := start
	switch dir {
	case 0:
		ender.Y -= mag
	case 1:
		ender.X += mag
	case 2:
		ender.Y += mag
	case 3:
		ender.X -= mag
	}
	return ender
}
func MoveModifierCoords8(start CoordInts, dir, mag int) CoordInts {
	ender := start
	switch dir {
	case 0:
		ender.Y -= mag
	case 1:
		ender.Y -= mag
		ender.X += mag
	case 2:
		ender.X += mag
	case 3:
		ender.X += mag
		ender.Y += mag
	case 4:
		ender.Y += mag
	case 5:
		ender.Y += mag
		ender.X -= mag
	case 6:
		ender.X -= mag
	case 7:
		ender.Y -= mag
		ender.X -= mag
	}
	return ender
}
func (igd *IntegerGridManager) MoveCursorFreely(dir int, speed int, walls []int) bool {
	if igd.PFinder.IsFullyInitialized {
		var tempCurs Cell = igd.PFinder.Cursor
		tempPos := MoveModifierCoords(tempCurs.Position, dir, speed)
		if igd.Imat.IsValid(tempPos) {
			if !IntArrayContains(walls, igd.Imat.GetCoordVal(tempPos)) {
				igd.PFinder.Cursor.Position = tempPos
				igd.UpdateCursor()

				return true

			} else {
				return false
			}
		}
	}
	return false
}
func (igd *IntegerGridManager) MoveCursorAroundPath(dir int, speed int, walls []int) {
	if igd.PFinder.IsFullyInitialized {
		var tempCurs Cell = igd.PFinder.Cursor
		tempPos := MoveModifierCoords(tempCurs.Position, dir, speed)
		if igd.Imat.IsValid(tempPos) {
			if !IntArrayContains(walls, igd.Imat.GetCoordVal(tempPos)) {
				igd.PFinder.Cursor.Position = tempPos
				igd.UpdateCursor()

			}
		}
	}

}
func (igd *IntegerGridManager) UpdateCursor() {
	temp, temp2, _ := igd.Imat.GetNeighbors8(igd.PFinder.Cursor.Position, [4]int{1, 2, 1, 2})
	igd.PFinder.Cursor.Neighbor_Values = temp2
	igd.PFinder.Cursor.Neighbors = [8]CoordInts(temp)
	temp3, temp4 := igd.PFinder.Cursor.GetCircle(igd.PFinder.Cursor.circRad, igd.Imat)
	igd.PFinder.Cursor.CirclePoints = temp3
	igd.PFinder.Cursor.CircleValues = temp4

	igd.FindPath(0)
	// igd.BoardChange = true
	igd.BoardOverlayChange = true
}
func (igd *IntegerGridManager) UpdateCursor2() {
	temp, temp2, _ := igd.Imat.GetNeighbors8(igd.PFinder.Cursor.Position, [4]int{1, 2, 1, 2})
	igd.PFinder.Cursor.Neighbor_Values = temp2
	igd.PFinder.Cursor.Neighbors = [8]CoordInts(temp)
	igd.FindPath(0)
}
func (igd *IntegerGridManager) DrawCursor_00(screen *ebiten.Image) {
	igd.Imat.DrawAGridTile_With_Line(screen, igd.PFinder.Cursor.Position, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{200, 150, 0, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)
	if igd.PFinder.Cursor.ShowNeighbors {
		for i, a := range igd.PFinder.Cursor.Neighbors {
			if igd.PFinder.Cursor.Neighbor_Values[i] == 1 {
				igd.Imat.DrawAGridTile_With_Line(screen, a, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{200, 0, 150, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)
			} else {
				igd.Imat.DrawAGridTile_With_Line(screen, a, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{0, 150, 200, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)
			}
		}
	}
}
func (igd *IntegerGridManager) DrawCursor(screen *ebiten.Image) {
	igd.Imat.DrawAGridTile_With_Line(screen, igd.PFinder.Cursor.Position, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{200, 150, 0, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)
	if igd.PFinder.Cursor.ShowNeighbors {
		for i, a := range igd.PFinder.Cursor.Neighbors {
			if igd.PFinder.Cursor.Neighbor_Values[i] == 1 {
				// igd.Imat.DrawAGridTile_With_Line(screen, a, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{200, 0, 150, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)
				igd.Imat.DrawAGridTile_With_Line(screen, a, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{200, 0, 150, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)

			} else {
				// igd.Imat.DrawAGridTile_With_Line(screen, a, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{0, 150, 200, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)
				igd.Imat.DrawAGridTile_With_Line(screen, a, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{0, 150, 200, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)

			}
		}
	}
	if igd.PFinder.Cursor.ShowCircle {
		// fmt.Printf("CURSOR %3d/%3d\n", 0, len(igd.PFinder.Cursor.CircleValues))
		for i, a := range igd.PFinder.Cursor.CirclePoints {
			if igd.PFinder.Cursor.CircleValues[i] == 1 {
				//fmt.Printf("CURSOR %3d/%3d\n", i, len(igd.PFinder.Cursor.CircleValues))
				// igd.Imat.DrawAGridTile_With_Line(screen, a, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{0, 200, 150, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)
				igd.Imat.DrawAGridTile_With_Line(screen, a, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{0, 200, 150, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)

			} else {
				// igd.Imat.DrawAGridTile_With_Line(screen, a, igd.BoardPosition.X, igd.BoardPosition.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{0, 150, 200, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)
				igd.Imat.DrawAGridTile_With_Line(screen, a, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{0, 200, 150, 255}, color.Black, color.Black, color.Black, 2.0, 2.0, 2.0, true, true, true, false)

			}
		}
	}
}

func (igd *IntegerGridManager) MoveCamToCursor() {

}

func (igd *IntegerGridManager) MoveCursorAround(Vect CoordInts, walls []int) (int, bool) {
	var tempCurs Cell = igd.PFinder.Cursor
	//step1;
	var valer int
	tempPos := MoveModifierCoords(tempCurs.Position, Vect.X, 0)
	//

	// tempPos := MoveModifierCoords(igd.PFinder.Cursor.Position, Vect.X, 0)
	// if len(igd.PFinder.Nodes) > 0 {
	// 	if igd.PFinder.Nodes.CountInstances(tempPos) == 0 {
	// 		igd.PFinder.Nodes = append(igd.PFinder.Nodes, tempPos)
	// 	}
	// }
	// if !IntArrayContains(walls, igd.Imat.GetCoordVal(tempPos)) {

	// 	for i := range Vect.Y {
	// 		tempPos = MoveModifierCoords(tempCurs.Position, Vect.X, i)
	// 		fmt.Printf("MOVE: %d - POINT: %d %d Val Here: %d\n", i, tempPos.X, tempPos.Y, igd.Imat.GetCoordVal(tempPos))
	// 		if IntArrayContains(walls, igd.Imat.GetCoordVal(tempPos)) {
	// 			valer = i

	// 			break
	// 		}
	// 	}
	// }
	// for i := range Vect.Y {
	for i := 0; i < len(igd.Imat); i++ {

		tempPos = MoveModifierCoords(tempCurs.Position, Vect.X, i)
		fmt.Printf("MOVE: %d - POINT: %d %d Val Here: %d\n", i, tempPos.X, tempPos.Y, igd.Imat.GetCoordVal(tempPos))
		if IntArrayContains(walls, igd.Imat.GetCoordVal(tempPos)) {
			valer = i

			igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, tempPos)
			tempPos = MoveModifierCoords(tempCurs.Position, Vect.X, i-1)
			break
		} else {
			igd.PFinder.Moves = append(igd.PFinder.Moves, tempPos)
		}
	}

	igd.PFinder.HasFalsePos = true
	fmt.Printf(" MOVE CURSOR AROUND: Dir: %d  Mag %d  Actual: %d \n\n", Vect.X, Vect.Y, valer)
	return valer, true
}

/*
goal is to find a slope
*/
func (cord CoordInts) MoveCursorAlongSlope(target CoordInts, ticks int) CoordInts {
	temp := cord
	var slope float64
	if temp.X == target.X || temp.Y == target.Y {

	} else if temp.X < target.X {
		slope = float64(target.Y-cord.Y) / float64(target.X-cord.X)
		temp = CoordInts{X: (temp.X + ticks), Y: int(float64(temp.Y) + (float64(ticks) * slope))}
	} else {
		// slope = float64(cord.Y-target.Y) / float64(cord.X-target.X)
		slope = float64(target.Y-cord.Y) / float64(target.X-cord.X)
		temp = CoordInts{X: (temp.X - ticks), Y: int(float64(temp.Y) - (float64(ticks) * slope))}
	}

	fmt.Printf("SLOPE is %f %d %d - TARGET:%d %d \n", slope, temp.X, temp.Y, target.X, target.Y)
	return temp
}
func (igd *IntegerGridManager) SLOPEMOVE(ticks int, walls []int) {
	var tempCurs Cell = igd.PFinder.Cursor
	tempPos := tempCurs.Position.MoveCursorAlongSlope(igd.PFinder.EndPos, ticks)
	igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, tempPos)

	for i := 0; (i < len(igd.Imat)) && (i < ticks); i++ {

		fmt.Printf("MOVE: %d - POINT: %d %d -to - %d %d - Val Here: %d\n", i, tempPos.X, tempPos.Y, tempCurs.Position.X, tempCurs.Position.Y, igd.Imat.GetCoordVal(tempPos))
		if IntArrayContains(walls, igd.Imat.GetCoordVal(tempPos)) {

			igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, tempPos)
			tempPos = tempCurs.Position.MoveCursorAlongSlope(igd.PFinder.EndPos, i-1)

			break
		} else {
			igd.PFinder.Moves = append(igd.PFinder.Moves, tempPos)
		}
	}
	igd.PFinder.Cursor.Position = tempPos

	// igd.PFinder.HasFalsePos = true
	igd.PFinder.HasFalsePos = true
}

// -------------------------------------------------------

func (igd *IntegerGridManager) GetACirclePointsOnClick(Raw_Mouse_X, Raw_Mouse_Y int, Radius int, valueIs int) CoordList {
	var center CoordInts
	var is_OnPoint bool
	center.X, center.Y, is_OnPoint = igd.Imat.GetCoordOfMouseEvent(Raw_Mouse_X, Raw_Mouse_Y, igd.Position.X, igd.Position.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y)
	var temp CoordList
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
			// x, y :=
			temp2 := igd.GetACirclePointsWSub(x, y, valueIs, center)
			temp = append(temp, temp2...)
			// for _, a := range temp2 {

			// }
			if x < y {
				break
			}
		}

		temp_01A := center
		temp_01A.X += Radius
		// igd.Imat[temp_01A.Y][temp_01A.X] = valueIs
		if igd.Imat.IsValid(temp_01A) {
			// igd.Imat[temp_01A.Y][temp_01A.X] = valueIs
			temp = append(temp, temp_01A)
		}
		temp_01B := center
		temp_01B.X -= Radius
		if igd.Imat.IsValid(temp_01B) {
			// igd.Imat[temp_01B.Y][temp_01B.X] = valueIs
			temp = append(temp, temp_01B)
		}
		temp_01C := center
		temp_01C.Y -= Radius
		if igd.Imat.IsValid(temp_01C) {
			// igd.Imat[temp_01C.Y][temp_01C.X] = valueIs
			temp = append(temp, temp_01C)
		}
		temp_01D := center
		temp_01D.Y += Radius
		if igd.Imat.IsValid(temp_01D) {
			// igd.Imat[temp_01D.Y][temp_01D.X]
			temp = append(temp, temp_01D)
		}
	}
	return temp
}
func (igd *IntegerGridManager) GetACirclePointsWSub(x, y, valueIs int, center CoordInts) CoordList {

	/*
			cout << "(" << x + x_centre << ", " << y + y_centre << ") ";
		        cout << "(" << -x + x_centre << ", " << y + y_centre << ") ";
		        cout << "(" << x + x_centre << ", " << -y + y_centre << ") ";
		        cout << "(" << -x + x_centre << ", " << -y + y_centre << ")\n";

	*/
	var temp CoordList
	temp_01A := center
	temp_01A.X += x
	temp_01A.Y += y

	temp_01B := center
	temp_01B.X -= x
	temp_01B.Y += y

	temp_02A := center
	temp_02B := center
	temp_02A.X += x
	temp_02A.Y -= y
	temp_02B.X -= x
	temp_02B.Y -= y
	if igd.Imat.IsValid(center) {
		igd.Imat[center.Y][center.X] = valueIs
	}
	if igd.Imat.IsValid(temp_01A) {
		// igd.Imat[temp_01A.Y][temp_01A.X] = valueIs
		igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, temp_01A)
	}
	if igd.Imat.IsValid(temp_01B) {
		// igd.Imat[temp_01B.Y][temp_01B.X] = valueIs
		igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, temp_01B)
	}
	if igd.Imat.IsValid(temp_02A) {
		// igd.Imat[temp_02A.Y][temp_02A.X] = valueIs
		igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, temp_02A)
	}
	if igd.Imat.IsValid(temp_02B) {
		// igd.Imat[temp_02B.Y][temp_02B.X] = valueIs
		igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, temp_02B)
	}

	if x != y {
		if igd.Imat.IsValid(CoordInts{center.X + y, center.Y + x}) {
			igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, CoordInts{X: center.X + y, Y: center.Y + x})
			// igd.Imat[center.Y+x][center.X+y] = valueIs
		}
		if igd.Imat.IsValid(CoordInts{center.X - y, center.Y + x}) {
			// igd.Imat[center.Y+x][center.X-y] = valueIs
			igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, CoordInts{X: center.Y + x, Y: center.X - x})
		}
		if igd.Imat.IsValid(CoordInts{center.X + y, center.Y - x}) {
			// igd.Imat[center.Y-x][center.X+y] = valueIs
			igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, CoordInts{X: center.Y - x, Y: center.X + x})
		}
		if igd.Imat.IsValid(CoordInts{center.X - y, center.Y - x}) {
			// igd.Imat[center.Y-x][center.X-y] = valueIs
			igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, CoordInts{X: center.X - x, Y: center.X - y})
		}
	}
	return temp
}
