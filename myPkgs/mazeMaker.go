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
	fmt.Printf("MAZEMaker INitalizing\n")
	mazeM.Imat = imat
	mazeM.Cords0 = make(CoordList, 0)
	mazeM.Cursor.Init0(CoordInts{X: -1, Y: -1})
	mazeM.Fails = 0
	mazeM.FailLimit = failLimit + 5
	mazeM.Cords0_IsVisible = false
	fmt.Printf("MAZEMaker INitalized\n")
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
	fmt.Printf("THIS IS DOING SOMETHING\n")
	// mazeM.Cords0 = mazeM.Cords0.PushToReturn(CoordInts{X: xx, Y: yy})
	mazeM.Cords0 = append(mazeM.Cords0, CoordInts{X: xx, Y: yy})

}

func (mazeM *MazeMaker) BasicDecayWrapper(Imat IntMatrix, filterFor []int, buffer [4]int, nSteps int) {

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
		templist, tempar, _ := mazeM.Imat.GetNeighborsButFiltered(c, []int{1, 2, 3, 4}, [4]int{1, 2, 2, 1})
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
func (imat *IntMatrix) BasicDecayProcess(CordsIn CoordList, fails int, filterFor []int, buffer [4]int) (CoordList, int) {
	temp := make(CoordList, len(CordsIn))
	copy(temp, CordsIn)
	fails_out := fails
	// var frustration bool = true
	for _, c := range CordsIn {
		imat.SetValAtCoord(c, 1)
		frustration := true
		templist, tempar, _ := imat.GetNeighborsButFiltered(c, []int{1, 2, 3, 4}, [4]int{5, 6, 6, 5})
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

func (mazeM *MazeMaker) MoreAdvancedDecay(filterFor []int, buffer [4]int) {
	temp := make(CoordList, len(mazeM.Cords0))
	copy(temp, mazeM.Cords0)

	// var frustration bool = true
	// for _, c := range mazeM.Cords0
	if len(mazeM.Cords0) > 0 {
		randInt := rand.Intn(len(mazeM.Cords0))
		c := mazeM.Cords0[randInt]
		frustration := true
		if mazeM.DiagonalChecking(c, filterFor, [4]int{2, 3, 3, 2}) {

			templist, tempar, _ := mazeM.Imat.GetNeighbors8(c, [4]int{2, 3, 3, 2})
			nEBool := tempar[1] != -1 && tempar[1] != 1
			sEBool := tempar[3] != -1 && tempar[3] != 1
			sWBool := tempar[5] != -1 && tempar[5] != 1
			nWBool := tempar[7] != -1 && tempar[7] != 1

			if tempar[0] != -1 && tempar[0] != 1 && nEBool && nWBool {
				temp = temp.PushToReturn(templist[0])

				frustration = false
			} else if tempar[2] != -1 && tempar[2] != 1 && nEBool && sEBool {
				temp = temp.PushToReturn(templist[2])

				frustration = false
			}
			if tempar[4] != -1 && tempar[4] != 1 && sEBool && sWBool {
				temp = temp.PushToReturn(templist[4])
				// temp, _ = temp.RemoveCoordFromList(c)
				frustration = false
			}
			if tempar[6] != -1 && tempar[6] != 1 && nWBool && sWBool {
				temp = temp.PushToReturn(templist[6])
				// temp, _ = temp.RemoveCoordFromList(c)
				frustration = false
			}
			temp, _ = temp.RemoveCoordFromList(c)
		} else {
			// mazeM.Imat.SetValAtCoord(c, 2)
		}

		if !frustration {
			temp, _ = temp.RemoveCoordFromList(c)
			mazeM.Imat.SetValAtCoord(c, 1)
			mazeM.Fails = 0
		} else {
			// temp, _ = temp.RemoveCoordFromList(c)
			mazeM.Fails++
			if mazeM.Fails > mazeM.FailLimit {
				temp, _ = temp.RemoveCoordFromList(c)
			}
		}
	}
	temp = temp.RemoveDuplicates()
	mazeM.Cords0 = temp
}

/*
this returns a number if the point being checked has
*/
func (mazeM *MazeMaker) DiagonalChecking(cord CoordInts, filter []int, buffer [4]int) bool {
	_, tempAr, _ := mazeM.Imat.GetNeighbors8(cord, buffer)
	num0 := 0
	num1 := 0
	// nEBool := tempAr[1] != -1 && tempAr[1] != 1
	// sEBool := tempAr[3] != -1 && tempAr[3] != 1
	// sWBool := tempAr[5] != -1 && tempAr[5] != 1
	// nWBool := tempAr[7] != -1 && tempAr[7] != 1
	if tempAr[0] != -1 && tempAr[0] != 1 {
		num1++
	}
	if tempAr[1] != -1 && tempAr[1] != 1 {
		num0++
	}
	if tempAr[2] != -1 && tempAr[2] != 1 {
		num1++
	}
	if tempAr[3] != -1 && tempAr[3] != 1 {
		num0++
	}
	if tempAr[4] != -1 && tempAr[4] != 1 {
		num1++
	}
	if tempAr[5] != -1 && tempAr[5] != 1 {
		num0++
	}
	if tempAr[6] != -1 && tempAr[6] != 1 {
		num1++
	}
	if tempAr[7] != -1 && tempAr[7] != 1 {
		num0++
	}
	d0 := false
	d1 := false
	if (tempAr[1] != -1 && tempAr[1] != 1) && (tempAr[5] != -1 && tempAr[5] != 1) {
		d0 = true
	}
	if (tempAr[3] != -1 && tempAr[3] != 1) && (tempAr[7] != -1 && tempAr[7] != 1) {
		d1 = true
	}
	h0 := false
	h1 := false
	h2 := false
	if (tempAr[7] != -1 && tempAr[7] != 1) && (tempAr[1] != -1 && tempAr[1] != 1) {
		h0 = true
	}
	if (tempAr[3] != -1 && tempAr[3] != 1) && (tempAr[5] != -1 && tempAr[5] != 1) {
		h1 = true
	}
	if (tempAr[2] != -1 && tempAr[2] != 1) && (tempAr[6] != -1 && tempAr[6] != 1) {
		h2 = true
	}
	v0 := false
	v1 := false
	v2 := false
	if (tempAr[7] != -1 && tempAr[7] != 1) && (tempAr[5] != -1 && tempAr[5] != 1) {
		v0 = true
	}
	if (tempAr[0] != -1 && tempAr[0] != 1) && (tempAr[4] != -1 && tempAr[4] != 1) {
		v1 = true
	}
	if (tempAr[3] != -1 && tempAr[3] != 1) && (tempAr[1] != -1 && tempAr[1] != 1) {
		v2 = true
	}

	if num0 > 1 {
		// if h0 && v1 && (h1 || h2) && (v0 || v2) && !(h1 && h2) && !(v0 && v2) { //&& (!h3 || !v3)
		// 	fmt.Printf("TEST-------010\n")
		// 	// if d0 && d1 {
		// 	// 	return true
		// 	// }
		// 	return true
		// }
		// if h0 && v0 && !h1 && !v2 { // && !v1
		// 	fmt.Printf("TEST000\n")
		// 	// if d0 && !d1 {
		// 	// 	return true
		// 	// }
		// 	// return false
		// 	return true
		// }
		// if h1 && v2 && !h0 && !v0 { //&& (!h3 || !v3)
		// 	fmt.Printf("TEST001\n")
		// 	// if d0 && d1 {
		// 	// 	return true
		// 	// }
		// 	// return false
		// 	return true
		// }
		// if h1 && v0 && !h0 && !v2 { //&& (!h3 || !v3)
		// 	fmt.Printf("TEST002\n")
		// 	// if d0 && d1 {
		// 	// 	return true
		// 	// }
		// 	// return false
		// 	return true
		// }
		// if h0 && v2 && !h1 && !v0 { //&& (!h3 || !v3)
		// 	fmt.Printf("TEST003\n")
		// 	// if d0 && d1 {
		// 	// 	return true
		// 	// }
		// 	return true
		// }
		//====================================
		if v2 && !v0 && !v1 { //&& (!h3 || !v3)
			fmt.Printf("TEST006\n")
			return true
			// if d0 && d1 {
			// 	fmt.Printf("TEST006\n")
			// 	return true
			// }
			// return false
		}
		if v0 && v1 && !v2 { //&& (!h3 || !v3)
			fmt.Printf("TEST004\n")
			// if d0 && d1 && (!h1 || !v1) {

			// 	return true
			// }
			return true
		}
		// if v1 && v2 && !v0 { //&& (!h3 || !v3)
		// 	fmt.Printf("TEST005\n")
		// 	// if d0 && d1 {
		// 	// 	return true
		// 	// }
		// 	return true
		// }

		if h2 && !h0 && !h1 { //&& (!h3 || !v3)
			fmt.Printf("TEST009\n")
			return true
			// if d0 && d1 {

			// }
			// return false
		}
		if h0 && h1 && !h2 { //&& (!h3 || !v3)
			fmt.Printf("TEST007\n")
			// if d0 && d1 && (!h1 || !v1) {
			// 	fmt.Printf("TEST007\n")
			// 	return true
			// }
			return true
		}
		// if h1 && h2 && !h0 { //&& (!h3 || !v3)
		// 	fmt.Printf("TEST008\n")
		// 	// if d0 && d1 {
		// 	// 	return true
		// 	// }
		// 	return true
		// }

		//
		if d0 && d1 {
			return true
		}
		return false
	}
	return true
	// return num, d1, d1
}

func (mazeM *MazeMaker) Process3() {

}
