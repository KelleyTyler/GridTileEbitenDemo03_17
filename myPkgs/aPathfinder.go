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
	Cursor             PF_Cursor
	FalsePos           CoordList
	Moves              CoordList
	HasFalsePos        bool
}

// func (igd *Pathfinding) Tick_DownFalseposlane() {

// }

func (pFndr *Pathfinding) ToString() string {
	outstrng := "PATHFINDING:\n"
	outstrng += fmt.Sprintf("\n START: %d, %d\n END: %d %d\n", pFndr.StartPos.X, pFndr.StartPos.Y, pFndr.EndPos.X, pFndr.EndPos.Y)
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
		igd.PFinder.Cursor.InitP(igd.PFinder.StartPos, igd.Imat)
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

func (igd *IntegerGridManager) PFindr_DrawManhattan() {
	if igd.PFinder.IsFullyInitialized {
		yFirst := true
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
				igd.MoveCursorSteps(CoordInts{dirs.X, 1}, 2, []int{0, 2, 3, 4})
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
func (igd *IntegerGridManager) PFindr_DrawBresenHamLine() {
	if igd.PFinder.IsFullyInitialized {
		if !igd.PFinder.HasFalsePos {
			igd.PFinder.FalsePos = BresenhamLine(igd.PFinder.StartPos, igd.PFinder.EndPos)
			igd.PFinder.HasFalsePos = true
		} else {
			// igd.PFinder.Cursor.Position = igd.PFinder.FalsePos[igd.PFinder.Cursor.ticker]
			temp := igd.PFinder.FalsePos[igd.PFinder.Cursor.ticker]
			if a, c := IntArrayContains_giveMeWhat([]int{0, 2, 3, 4}, igd.Imat.GetCoordVal(temp)); a {
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

type PF_Cursor struct {
	Position        CoordInts
	Neighbors       [4]CoordInts
	Neighbor_Values [4]int
	ticker          int
	// Previous        *PF_Cursor
	// Next            *PF_Cursor
	// Number          int
}

func (pfCurs *PF_Cursor) InitP(StartPos CoordInts, imat IntMatrix) {
	// limit_X, limit_Y := imat.GetDimensions()

	pfCurs.Position = StartPos
	temp, temp2, _ := imat.GetNeighbors(StartPos)
	pfCurs.Neighbor_Values = temp2
	pfCurs.Neighbors = [4]CoordInts(temp)
	pfCurs.ticker = 0
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

func (igd *IntegerGridManager) MoveCursorSteps(Vect CoordInts, steps int, walls []int) (int, bool) {
	var tempCurs PF_Cursor = igd.PFinder.Cursor
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

func (igd *IntegerGridManager) MoveCursorAround(Vect CoordInts, walls []int) (int, bool) {
	var tempCurs PF_Cursor = igd.PFinder.Cursor
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
	var tempCurs PF_Cursor = igd.PFinder.Cursor
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
