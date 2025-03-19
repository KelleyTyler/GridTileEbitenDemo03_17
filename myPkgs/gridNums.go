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
