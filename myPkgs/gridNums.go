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

	TempMatrix := make(IntMatrix, numColumns)
	for row := range TempMatrix {
		TempMatrix[row] = make([]int, numRows)
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
