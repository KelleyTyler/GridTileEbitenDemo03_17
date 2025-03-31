package mypkgs

import (
	"fmt"
)

type IntMatrix [][]int

// func (iMat *IntMatrix) Init(numColumns int, numRows int) {
// 	fmt.Printf("HEy")
// 	// TempMatrix := make(IntMatrix, r)
// 	for row := range *iMat {
// 		iMat[row] = make([]int, x)
// 	}
// 	//return TempMatrix

// }

func (iMat *IntMatrix) MakeIntMatrix(numColumns int, numRows int) IntMatrix {

	TempMatrix := make(IntMatrix, numRows)
	for row := range TempMatrix {
		TempMatrix[row] = make([]int, numColumns)
	}
	return TempMatrix
}

func (iMat IntMatrix) PrintMatrix() {
	fmt.Printf("PRINTING INTMATRIX:\n\n")
	fmt.Printf("SIZE:%d\n", len(iMat))
	for i, _ := range iMat {
		for _, c := range iMat[i] {
			fmt.Printf("[%2d]", c)
		}
		fmt.Printf("\n")
	}
}

func (iMat IntMatrix) InitBlankMatrix(numColumns int, numRows int, fillWith int) {
	for r, _ := range iMat {
		for c, _ := range iMat[r] {
			//iMat[r] = iMat[r].append(iMat[r], fillWith)
			iMat[r][c] = fillWith
		}
	}
}

func (iMat IntMatrix) GetDimensions() (int, int) {
	sizeY := len(iMat)
	sizeX := len(iMat[0])
	return sizeX, sizeY
}

func (iMat IntMatrix) GetNeighbors4(coordPoint CoordInts, buffer [4]int) (CoordList, [4]int, int) {
	outList := make(CoordList, 4)
	var outAr [4]int
	//North
	outList[0] = CoordInts{coordPoint.X, coordPoint.Y - 1}
	if coordPoint.Y < buffer[0] {
		outAr[0] = -1
	} else {
		outAr[0] = iMat.GetCoordVal(outList[0])
	}
	//east
	outList[1] = CoordInts{coordPoint.X + 1, coordPoint.Y}
	if coordPoint.X > len(iMat[0])-buffer[1] {
		outAr[1] = -1
	} else {
		outAr[1] = iMat.GetCoordVal(outList[1])
	}
	//south
	outList[2] = CoordInts{coordPoint.X, coordPoint.Y + 1}
	if coordPoint.Y > len(iMat)-buffer[2] {
		outAr[2] = -1
	} else {
		outAr[2] = iMat.GetCoordVal(outList[2])
	}
	//west
	outList[3] = CoordInts{coordPoint.X - 1, coordPoint.Y}
	if coordPoint.X < buffer[3] {
		outAr[3] = -1
	} else {
		outAr[3] = iMat.GetCoordVal(outList[3])
	}
	//----
	return outList, outAr, iMat.GetCoordVal(coordPoint)
}
func (iMat IntMatrix) GetNeighbors8(coordPoint CoordInts, buffer [4]int) (CoordList, [8]int, int) {
	outList := make(CoordList, 8)
	var outAr [8]int
	//North
	outList[0] = CoordInts{coordPoint.X, coordPoint.Y - 1}
	if coordPoint.Y < buffer[0] {
		outAr[0] = -1
	} else {
		outAr[0] = iMat.GetCoordVal(outList[0])
	}
	//northeast
	outList[1] = CoordInts{coordPoint.X + 1, coordPoint.Y - 1}
	if coordPoint.X > len(iMat[0])-buffer[1] || coordPoint.Y < buffer[0] {
		outAr[1] = -1
	} else {
		outAr[1] = iMat.GetCoordVal(outList[1])
	}
	//east
	outList[2] = CoordInts{coordPoint.X + 1, coordPoint.Y}
	if coordPoint.X > len(iMat[0])-buffer[1] {
		outAr[2] = -1
	} else {
		outAr[2] = iMat.GetCoordVal(outList[2])
	}
	//south east
	outList[3] = CoordInts{coordPoint.X + 1, coordPoint.Y + 1}
	if coordPoint.X > len(iMat[0])-buffer[1] || coordPoint.Y > len(iMat)-buffer[2] {
		outAr[3] = -1
	} else {
		outAr[3] = iMat.GetCoordVal(outList[3])
	}

	//south
	outList[4] = CoordInts{coordPoint.X, coordPoint.Y + 1}
	if coordPoint.Y > len(iMat)-buffer[2] {
		outAr[4] = -1
	} else {
		outAr[4] = iMat.GetCoordVal(outList[4])
	}
	//southwest
	outList[5] = CoordInts{coordPoint.X - 1, coordPoint.Y + 1}
	if coordPoint.X < buffer[3] || coordPoint.Y > len(iMat)-buffer[2] {
		outAr[5] = -1
	} else {
		outAr[5] = iMat.GetCoordVal(outList[5])
	}
	//west
	outList[6] = CoordInts{coordPoint.X - 1, coordPoint.Y}
	if coordPoint.X < buffer[3] {
		outAr[6] = -1
	} else {
		outAr[6] = iMat.GetCoordVal(outList[6])
	}
	//northwest
	outList[7] = CoordInts{coordPoint.X - 1, coordPoint.Y - 1}
	if coordPoint.X < buffer[3] || coordPoint.Y < buffer[0] {
		outAr[7] = -1
	} else {
		outAr[7] = iMat.GetCoordVal(outList[7])
	}
	//----
	return outList, outAr, iMat.GetCoordVal(coordPoint)
}
func (iMat IntMatrix) GetNeighborsButFiltered(coordPoint CoordInts, FilterFor []int, buffer [4]int) (CoordList, [4]int, int) {
	outList, numList, num := iMat.GetNeighbors4(coordPoint, buffer) //[4]int{1, 2, 2, 1}
	outList2 := make(CoordList, len(outList))
	copy(outList2, outList)

	for i, c := range numList {
		for _, x := range FilterFor {
			if c == x {
				numList[i] = -1
				outList2, _ = outList2.RemoveCoordFromList(outList[i])
			}
		}
	}
	return outList, numList, num
}

func (iMat IntMatrix) SetValAtCoord(coord CoordInts, newVal int) {
	iMat[coord.Y][coord.X] = newVal
}
func (iMat IntMatrix) CycleValAtCoord(coord CoordInts, min, max, iterator int, circleBack bool) {
	curr := iMat.GetCoordVal(coord)
	if iterator > 0 {
		if curr+iterator > max {
			if circleBack {
				iMat.SetValAtCoord(coord, min)
			}
		} else {
			iMat.SetValAtCoord(coord, curr+iterator)
		}

	} else if iterator < 0 {
		if curr+iterator < min {
			if circleBack {
				iMat.SetValAtCoord(coord, max)
			}
		} else {
			iMat.SetValAtCoord(coord, curr+iterator)
		}
	}
}

//	func (iMat IntMatrix) MakeIntMatrixRowBlank(row int, numColumns int, fillWith int) {
//		var temp [40]int
//		for i := 0; i < numColumns; i++ {
//			temp[i] = 0
//		}
//		for j := 0; j < numColumns; j++ {
//			iMat[row][j] = temp[j]
//		}
//	}

/*
	Function: ValidatePointAgainstArrayOfValues

-This looks for the point described by 'coord' on the IntMatrix;
-If that point is found it compares the value of that point to the values in the array 'num'
-- if there is a match 'true' is returned; if there is no match false is returned;
-If that value is not found/the point is out of the bounds of the Intmatrix it returs false;
*/
func (imat *IntMatrix) IsCoordValueInArrayOfValues(coord CoordInts, num []int) bool {
	ret := false
	if imat.IsValid(coord) {
		x := imat.GetCoordVal(coord)
		ret = IntArrayContains(num, x)
	}
	return ret
}

func (imat *IntMatrix) IsCoordValueInArrayOfValues_What_Exists(coord CoordInts, num []int) (bool, bool, int) {
	exists := false
	ret := false
	y := -2
	if imat.IsValid(coord) {
		x := imat.GetCoordVal(coord)
		exists = true
		ret, y = IntArrayContains_giveMeWhat(num, x)
	}
	return exists, ret, y
}

/*
IntMatrix.BasicDecayStep()
This function takes a List of CoordInts, which form the 'frontier' of this cascading effect; it will move outwards towards
*/
func (imat *IntMatrix) BasicDecayStep(Value_To_Set_As, fails int, FrontCoordList CoordList, filterFor []int, buffer [4]int) (CoordList, int) {
	temp := make(CoordList, len(FrontCoordList))
	copy(temp, FrontCoordList)
	fails_out := fails
	// var frustration bool = true
	for _, c := range FrontCoordList {
		imat.SetValAtCoord(c, Value_To_Set_As)
		frustration := true
		templist, tempar, _ := imat.GetNeighborsButFiltered(c, filterFor, buffer) //[]int{1, 2, 3, 4}, [4]int{1, 2, 2, 1}
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
			fails_out++
		}
	}
	temp = temp.RemoveDuplicates()

	return temp, fails_out
}
func (imat *IntMatrix) BasicDecayProcess(nsteps, Value_To_Set_As, current_fails, max_fails int, FrontCoordList CoordList, filterFor []int, buffer [4]int) (CoordList, int, bool) {
	temp := make(CoordList, len(FrontCoordList))
	copy(temp, FrontCoordList)
	IsComplete := false
	fails_out := current_fails
	for range nsteps {
		temp, fails_out = imat.BasicDecayStep(Value_To_Set_As, fails_out, temp, filterFor, buffer)
		if fails_out > max_fails {
			IsComplete = true //For better or worse this is the only way I think I can really test to see if this is complete;
		}
	}
	temp = temp.RemoveDuplicates()
	return temp, fails_out, IsComplete
}

//ValFloorAs, ValWallAs,

func (imat *IntMatrix) PrimLike_Maze_Algorithm_Step(FCLNum, fails, max_fails int, FrontCoordList CoordList, filterFor, filter1 []int, buffer [4]int, cullDiags bool) (CoordList, int) {
	temp := make(CoordList, len(FrontCoordList))
	copy(temp, FrontCoordList)
	fails_out := fails
	if len(temp) > 0 {
		//randInt := rand.Intn(len(mazeM.Cords0))
		// c := mazeM.Cords0[randInt]
		c := temp[FCLNum]
		frustration := true //mazeM.DiagonalChecking(c, filterFor, [4]int{2, 3, 3, 2})
		templist, tempar, _ := imat.GetNeighbors8(c, [4]int{2, 3, 3, 2})
		if imat.PrimMazeGenCell_CheckingRules(c, filterFor, buffer) {

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
			imat.SetValAtCoord(c, 4)
			//fmt.Printf("\nFAILURE AT %d, %d \n", c.X, c.Y)
			temp, _ = temp.RemoveCoordFromList(c)
		}

		if !frustration {
			temp, _ = temp.RemoveCoordFromList(c)
			imat.SetValAtCoord(c, 1)
			fails_out = 0

		} else {
			// temp, _ = temp.RemoveCoordFromList(templist[1])
			// temp, _ = temp.RemoveCoordFromList(templist[3])
			// temp, _ = temp.RemoveCoordFromList(templist[5])
			// temp, _ = temp.RemoveCoordFromList(templist[7])
			// temp, _ = temp.RemoveCoordFromList(c)
			fails_out++
			if fails_out > max_fails {
				temp, _ = temp.RemoveCoordFromList(c)
				// mazeM.Imat.SetValAtCoord(c, 4)
			}
		}
	}
	temp = temp.RemoveDuplicates()
	return temp, fails_out
}
func (imat *IntMatrix) PrimMazeGenCell_CheckingRules(cord CoordInts, filter []int, buffer [4]int) bool {
	_, tempAr, _ := imat.GetNeighbors8(cord, buffer)
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
