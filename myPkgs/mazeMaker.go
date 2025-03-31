package mypkgs

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// why a type who knows; think it might be a good way to bundle all the params I'd want into a single bundle of things;
type MazeMaker struct {
	Imat             *IntMatrix //possibly??
	ProcessStarted   bool
	ProcessOngoing   bool
	ProcessEnd       bool
	Cords0           CoordList
	Cords0_IsVisible bool
	Cursor           Cell
	Fails            int
	FailLimit        int
}

func (mazeM *MazeMaker) Init(imat *IntMatrix, failLimit int) {
	// fmt.Printf("MAZEMaker Initalizing\n")
	mazeM.Imat = imat
	mazeM.Cords0 = make(CoordList, 0)
	mazeM.Cursor.Init0(CoordInts{X: -1, Y: -1})
	mazeM.Fails = 0
	mazeM.FailLimit = failLimit + 5
	mazeM.Cords0_IsVisible = false
	// fmt.Printf("MAZEMaker Initialized\n")
}

func (mazeM *MazeMaker) Update() {
	if mazeM.Fails > mazeM.FailLimit {
		mazeM.ProcessEnd = true
		mazeM.ProcessOngoing = false
		mazeM.ProcessStarted = false
	}
}
func (mazeM *MazeMaker) Draw_CoordLines_raw(screen *ebiten.Image, offsetX, offsetY, tileW, tileH, gapX, gapY int, clr0 color.Color) {
	mazeM.Imat.DrawListAsTiles_withLines(screen, mazeM.Cords0, offsetX, offsetY, tileW, tileH, gapX, gapY, clr0, color.Black, color.Black, 2.0, true)
}
func (mazeM *MazeMaker) ClearCords0() {
	mazeM.Cords0 = make(CoordList, 0)
}

func (mazeM *MazeMaker) DrawCoordLinesFromIGD(igd IntegerGridManager, clr0 color.Color) {
	mazeM.Imat.DrawListAsTiles_withLines(igd.Img, mazeM.Cords0, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, clr0, color.Black, color.Black, 2.0, true)
}
func (mazeM *MazeMaker) ToString() string {
	strng := "MAZE MAKER:\n"
	strng += fmt.Sprintf("Cords0 size: %3d\n", len(mazeM.Cords0))
	strng += fmt.Sprintf("PROCESS:\n %8s: %5t\n %8s: %5t\n %8s: %5t\n", "Started", mazeM.ProcessStarted, "ongoing", mazeM.ProcessOngoing, "ended", mazeM.ProcessEnd)
	strng += fmt.Sprintf(" Fails: %3d : %3d\n", mazeM.Fails, mazeM.FailLimit)
	return strng
}
func (mazeM *MazeMaker) PrintString() {
	fmt.Printf("%s\n", mazeM.ToString())
}

func (mazeM *MazeMaker) ProcessStart() {

}
func (mazeM *MazeMaker) AddToCoords(xx, yy int) {
	// fmt.Printf("THIS IS DOING SOMETHING\n")
	// mazeM.Cords0 = mazeM.Cords0.PushToReturn(CoordInts{X: xx, Y: yy})
	mazeM.Cords0 = append(mazeM.Cords0, CoordInts{X: xx, Y: yy})

}

func (mazeM *MazeMaker) BasicDecayWrapper(Imat IntMatrix, filterFor []int, buffer [4]int, nSteps int) {
	if mazeM.ProcessOngoing {
		isDone := false

		if len(mazeM.Cords0) > 0 {
			mazeM.Cords0, mazeM.Fails, isDone = mazeM.Imat.BasicDecayProcess(nSteps, 1, mazeM.Fails, mazeM.FailLimit, mazeM.Cords0, filterFor, buffer)
		}
		if isDone {
			mazeM.ProcessEnd = true
			mazeM.ProcessOngoing = false
		}
	}
}
func (mazeM *MazeMaker) BasicDecayWrapper00(Imat IntMatrix, filterFor []int, buffer [4]int, nSteps int) {
	if mazeM.ProcessOngoing {
		if len(mazeM.Cords0) > 0 {
			for range nSteps {
				mazeM.BasicDecayProcess(filterFor, buffer)
			}
		}
	}
}

// func (mazeM *MazeMaker) BasicDecayProcess(Imat IntMatrix, filterFor []int, buffer [4]int) {
func (mazeM *MazeMaker) BasicDecayProcess(filterFor []int, buffer [4]int) {
	temp := make(CoordList, len(mazeM.Cords0))
	copy(temp, mazeM.Cords0)

	// var frustration bool = true
	for _, c := range mazeM.Cords0 {
		mazeM.Imat.SetValAtCoord(c, 1)
		frustration := true
		templist, tempar, _ := mazeM.Imat.GetNeighborsButFiltered(c, filterFor, buffer) //[]int{1, 2, 3, 4}, [4]int{1, 2, 2, 1}
		if tempar[0] != -1 && tempar[0] != 1 {
			temp = temp.PushToReturn(templist[0])
			// temp, _ = temp.RemoveCoordFromList(c)
			frustration = false
		}
		if tempar[1] != -1 && tempar[1] != 1 {
			temp = temp.PushToReturn(templist[1])
			// temp, _ = temp.RemoveCoordFromList(c)
			frustration = false
		}
		if tempar[2] != -1 && tempar[2] != 1 {
			temp = temp.PushToReturn(templist[2])
			// temp, _ = temp.RemoveCoordFromList(c)
			frustration = false
		}
		if tempar[3] != -1 && tempar[3] != 1 {
			temp = temp.PushToReturn(templist[3])
			// temp, _ = temp.RemoveCoordFromList(c)
			frustration = false
		}
		if !frustration {
			temp, _ = temp.RemoveCoordFromList(c)
		} else {
			mazeM.Fails++
		}
	}
	temp = temp.RemoveDuplicates()
	mazeM.Cords0 = temp
}

func (mazeM *MazeMaker) PrimLike_Maze_Algorithm00(filterFor []int, filter2 []int, buffer [4]int, cullDiags bool) {

	temp := make(CoordList, len(mazeM.Cords0))
	copy(temp, mazeM.Cords0)
	var failsOut = mazeM.Fails

	if len(mazeM.Cords0) > 0 {
		randInt := rand.Intn(len(mazeM.Cords0))
		temp, failsOut = mazeM.Imat.PrimLike_Maze_Algorithm_Step(randInt, failsOut, mazeM.FailLimit, temp, filterFor, filter2, buffer, cullDiags)
	}

	mazeM.Cords0 = temp
	mazeM.Fails = failsOut
}
func (mazeM *MazeMaker) PrimLike_Maze_Algorithm00_Looper(filterFor []int, filter2 []int, buffer [4]int, cullDiags bool) {
	temp := make(CoordList, len(mazeM.Cords0))
	copy(temp, mazeM.Cords0)
	var failsOut = mazeM.Fails
	for i := 0; i < len(temp); i++ {
		temp, failsOut = mazeM.Imat.PrimLike_Maze_Algorithm_Step(i, failsOut, mazeM.FailLimit, temp, filterFor, filter2, buffer, cullDiags)
	}
	mazeM.Cords0 = temp
	mazeM.Fails = failsOut
}
func (mazeM *MazeMaker) PrimLike_Maze_Algorithm03(filterFor []int, filter2 []int, buffer [4]int, cullDiags bool) {
	temp := make(CoordList, len(mazeM.Cords0))
	copy(temp, mazeM.Cords0)

	// var frustration bool = true
	// for _, c := range mazeM.Cords0
	if len(mazeM.Cords0) > 0 {
		randInt := rand.Intn(len(mazeM.Cords0))
		c := mazeM.Cords0[randInt]
		frustration := true //mazeM.DiagonalChecking(c, filterFor, [4]int{2, 3, 3, 2})
		templist, tempar, _ := mazeM.Imat.GetNeighbors8(c, [4]int{2, 3, 3, 2})
		if mazeM.PrimMazeGenCell_CheckingRules(c, filterFor, buffer) {

			nEBool := !IntArrayContains([]int{-1, 1, 4}, tempar[1])
			sEBool := !IntArrayContains([]int{-1, 1, 4}, tempar[3])
			sWBool := !IntArrayContains([]int{-1, 1, 4}, tempar[5])
			nWBool := !IntArrayContains([]int{-1, 1, 4}, tempar[7])

			if tempar[0] != -1 && tempar[0] != 1 && tempar[0] != 4 && nEBool && nWBool {
				temp = temp.PushToReturn(templist[0])

				frustration = false
			}
			if tempar[2] != -1 && tempar[2] != 1 && tempar[2] != 4 && nEBool && sEBool {
				temp = temp.PushToReturn(templist[2])

				frustration = false
			}
			if tempar[4] != -1 && tempar[4] != 1 && tempar[4] != 4 && sEBool && sWBool {
				temp = temp.PushToReturn(templist[4])
				// temp, _ = temp.RemoveCoordFromList(c)
				frustration = false
			}
			if tempar[6] != -1 && tempar[6] != 1 && tempar[6] != 4 && nWBool && sWBool {
				temp = temp.PushToReturn(templist[6])
				// temp, _ = temp.RemoveCoordFromList(c)
				frustration = false
			}
			if cullDiags {
				temp, _ = temp.RemoveCoordFromList(templist[1])
				temp, _ = temp.RemoveCoordFromList(templist[3])
				temp, _ = temp.RemoveCoordFromList(templist[5])
				temp, _ = temp.RemoveCoordFromList(templist[7])
			}

		} else {
			mazeM.Imat.SetValAtCoord(c, 4)
			//fmt.Printf("\nFAILURE AT %d, %d \n", c.X, c.Y)
			temp, _ = temp.RemoveCoordFromList(c)
		}

		if !frustration {
			temp, _ = temp.RemoveCoordFromList(c)
			mazeM.Imat.SetValAtCoord(c, 1)
			mazeM.Fails = 0

		} else {
			// temp, _ = temp.RemoveCoordFromList(templist[1])
			// temp, _ = temp.RemoveCoordFromList(templist[3])
			// temp, _ = temp.RemoveCoordFromList(templist[5])
			// temp, _ = temp.RemoveCoordFromList(templist[7])
			// temp, _ = temp.RemoveCoordFromList(c)
			mazeM.Fails++
			if mazeM.Fails > mazeM.FailLimit {
				temp, _ = temp.RemoveCoordFromList(c)
				// mazeM.Imat.SetValAtCoord(c, 4)
			}
		}
	}
	temp = temp.RemoveDuplicates()
	mazeM.Cords0 = temp
}

/*
Prim's Maze Generation Algorithm;
>have a grid of cells
>have an array/list/thing for "cells in Maze"
>have another list "frontier"--->these are cells that are not in the maze but are adjacent to the cells that are in the maze
->select cell at random for the maze
->add adjacent cells to frontier;
->select a cell randomly from frontier; add it to the maze; remove it from frontier
*/
// func (mazeM *MazeMaker) PrimsMazeAlgorithm() {

// }

/*
-'Additions to maze' that will be banned: (center is where cord is)
- 1 = "in MazeList"/floor
- 0 = not in maze list/wall
- ? = does not matter (regardless if it's in maze list or not)
(Ex_Row0)-------------------------------------------------------------------------
----0-------1-------2-------3-------4-------5-------6-------7-------8-------9----
| 1 0 1 | 1 0 1 | 1 1 1 | 1 1 1 | 1 1 1 | 1 0 1 | 1 0 0 | 1 0 1 | 1 0 1 | 0 0 0 |
| 0 0 1 | 1 0 0 | 1 0 0 | 0 0 1 | 1 0 0 | 0 0 1 | 0 0 1 | 1 0 0 | 0 0 0 | 0 0 0 |
| 1 1 1 | 1 1 1 | 1 0 1 | 1 0 1 | 0 0 1 | 0 1 1 | 1 1 1 | 1 1 0 | 1 0 1 | 0 0 0 |
(Ex_Row1)-------------------------------------------------------------------------
| 1 0 1 | 1 0 1 | 1 1 1 | 1 0 1 | 0 0 0 | 0 0 0 | 0 0 0 | 0 0 0 | 0 0 0 | 0 0 0 |
| 0 0 0 | 1 0 0 | 0 0 0 | 0 0 1 | 0 0 0 | 0 0 0 | 0 0 0 | 0 0 0 | 0 0 0 | 0 0 0 |
| 1 1 1 | 1 0 1 | 1 0 1 | 1 0 1 | 0 0 0 | 0 0 0 | 0 0 0 | 0 0 0 | 0 0 0 | 0 0 0 |
---------------------------------------------------------------------------------
---------------------------------------------------------------------------------
>Written longform it looks something like this;
----   {NN,NE,EE,SE,SS,SW,WW,NW}
- 0,0: {tt,ff,ff,ff,ff,ff,tt,ff}
- 0,1: {tt,ff,tt,ff,ff,ff,ff,ff}
- 0,2: {ff,ff,tt,ff,tt,ff,ff,ff}
- 0,3: {ff,ff,ff,ff,tt,ff,tt,ff}
- 0,4: {ff,ff,tt,ff,tt,tt,ff,ff}
- 0,5: {tt,ff,ff,ff,ff,tt,tt,ff}
- 0,6: {tt,tt,ff,ff,ff,ff,tt,ff}
- 0,7: {tt,ff,tt,tt,ff,ff,ff,ff}
- 0,8: {tt,ff,tt,ff,tt,ff,tt,ff}

- 1,0: {tt,ff,tt,ff,ff,ff,tt,ff}
- 1,1: {tt,ff,tt,ff,tt,ff,ff,ff}
- 1,2: {ff,ff,tt,ff,tt,ff,tt,ff}
- 1,3: {tt,ff,ff,ff,tt,ff,tt,ff}
---------------------------------------------------------------------------------

so solutions:

---------------------------------------------------------------------------------
*/
func (mazeM *MazeMaker) PrimMazeGenCell_CheckingRules(cord CoordInts, filter []int, buffer [4]int) bool {
	_, tempAr, _ := mazeM.Imat.GetNeighbors8(cord, buffer)
	//fmt.Printf("AT %2d, %2d --->", cord.X, cord.Y)
	//---------------------
	nn := !IntArrayContains([]int{-1, 1, 4}, tempAr[0])
	ne := !IntArrayContains([]int{-1, 1, 4}, tempAr[1])
	ee := !IntArrayContains([]int{-1, 1, 4}, tempAr[2])
	se := !IntArrayContains([]int{-1, 1, 4}, tempAr[3])
	ss := !IntArrayContains([]int{-1, 1, 4}, tempAr[4])
	sw := !IntArrayContains([]int{-1, 1, 4}, tempAr[5])
	ww := !IntArrayContains([]int{-1, 1, 4}, tempAr[6])
	nw := !IntArrayContains([]int{-1, 1, 4}, tempAr[7])
	// if !ww && !ee && !nn && !ss {
	// 	fmt.Printf("no Cardinals_ \n")
	// 	return true
	// }
	// if (!ww && !ee && nn && ss) || (ww && ee && !nn && !ss) {
	// 	fmt.Printf("not enough Cardinals_ \n")
	// 	return false
	// }
	// if !nw && !sw && !se && !ne {
	// 	fmt.Printf("no intercardinals \n")
	// 	return true
	// }
	// if (!nw && !sw && !se && ne) || (nw && !sw && !se && !ne) || (!nw && sw && !se && !ne) || (!nw && !sw && se && !ne) {
	// 	fmt.Printf("not enough intercardinals \n")
	// 	return false
	// }
	if (!nn) && ((se && !sw) || (!se && sw)) {
		//fmt.Printf("VeryStarange 2 N\n")
		return false
	}
	if (!ee) && ((!sw && nw) || (sw && !nw)) {
		//fmt.Printf("VeryStarange 2 E\n")
		return false
	}
	if (!ww) && ((!se && ne) || (!ne && se)) {
		//fmt.Printf("VeryStarange 2 W\n")
		return false
	}
	if (!ss) && ((!ne && nw) || (ne && !nw)) { //(nn && ee && ww && !ss) && ((!ne && se && sw && nw) || (ne && se && sw && !nw))
		//fmt.Printf("VeryStarange 2 S\n")
		return false
	}
	//fmt.Printf("\n")
	return true
}

/*
	Idea is to make a stippling pattern:
	before:    after:
	|0 0 0 0 0|0 0 0 0 0| the point selected
	|0 0 0 0 0|0 1 0 1 0| here is (1,1)
	|0 0 0 0 0|0 0 0 0 0| with a cell size of 3
	|0 0 0 0 0|0 1 0 1 0|
	|0 0 0 0 0|0 0 0 0 0|
*/

// func (mazeM *MazeMaker) Stippling(cell_size CoordInts) {
// 	if !mazeM.ProcessOngoing {
// 		if len(mazeM.Cords0) > 1 {
// 			mazeM.Cords0
// 		}
// 	}
// }
