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

func (iMat IntMatrix) GetNeighbors(coordPoint CoordInts) (CoordList, [4]int, int) {
	outList := make(CoordList, 4)
	var outAr [4]int
	//North
	outList[0] = CoordInts{coordPoint.X, coordPoint.Y - 1}
	if coordPoint.Y < 1 {
		outAr[0] = -1
	} else {
		outAr[0] = iMat.GetCoordVal(outList[0])
	}
	//east
	outList[1] = CoordInts{coordPoint.X + 1, coordPoint.Y}
	if coordPoint.X > len(iMat[0])-2 {
		outAr[1] = -1
	} else {
		outAr[1] = iMat.GetCoordVal(outList[1])
	}
	//south
	outList[2] = CoordInts{coordPoint.X, coordPoint.Y + 1}
	if coordPoint.Y > len(iMat)-2 {
		outAr[2] = -1
	} else {
		outAr[2] = iMat.GetCoordVal(outList[2])
	}
	//west
	outList[3] = CoordInts{coordPoint.X - 1, coordPoint.Y}
	if coordPoint.X < 1 {
		outAr[3] = -1
	} else {
		outAr[3] = iMat.GetCoordVal(outList[3])
	}
	//----
	return outList, outAr, iMat.GetCoordVal(coordPoint)
}

func (iMat IntMatrix) GetNeighborsButFiltered(coordPoint CoordInts, FilterFor []int) (CoordList, [4]int, int) {
	outList, numList, num := iMat.GetNeighbors(coordPoint)
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

// func (iMat *IntMatrix)

// func (iMat IntMatrix) MakeIntMatrixRowBlank(row int, numColumns int, fillWith int) {
// 	var temp [40]int
// 	for i := 0; i < numColumns; i++ {
// 		temp[i] = 0
// 	}
// 	for j := 0; j < numColumns; j++ {
// 		iMat[row][j] = temp[j]
// 	}
// }
