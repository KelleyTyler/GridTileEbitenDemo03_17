package mypkgs

import (
	"fmt"
	"math"
)

/*	This returns only the points before the line (c1-c2) encounters an obstacle (a value in imat from nums)
 */
func BresenhamLine_CullAfterImpact(c1, c2 CoordInts, imat IntMatrix, nums []int) (CoordList, bool) {
	var temp CoordList
	temp0 := BresenhamLine(c1, c2)
	EndedShort := false
	for _, a := range temp0 {
		if imat.IsValid(a) {
			if IntArrayContains(nums, imat.GetCoordVal(a)) {
				EndedShort = true
				break
			} else {
				temp = append(temp, a)
			}
		}
	}
	return temp, EndedShort
}

func BresenhamLine(c1 CoordInts, c2 CoordInts) CoordList {
	//outList := make(CoordList, 0)
	var outList CoordList
	if math.Abs(float64(c2.Y)-float64(c1.Y)) < math.Abs(float64(c2.X)-float64(c1.X)) {
		if c1.X > c2.X {
			fmt.Printf("BRESENHAM:%16s \n", "Low Inverted")
			outList = BresenhamLine_Low(c2, c1)
			outList = outList.FlipOrder()
		} else {
			fmt.Printf("BRESENHAM:%16s \n", "Low Regular")
			outList = BresenhamLine_Low(c1, c2)
		}
	} else {
		if c1.Y > c2.Y {
			fmt.Printf("BRESENHAM:%16s \n", "High Inverted")
			outList = BresenHamLine_High(c2, c1)
			outList = outList.FlipOrder()
		} else {
			fmt.Printf("BRESENHAM:%16s \n", "High Regular")
			outList = BresenHamLine_High(c1, c2)
		}
	}
	if !outList[0].IsEqualTo(c1) {
		outList = outList.PushToFrontThenReturn(c1)
	}
	if !outList[len(outList)-1].IsEqualTo(c2) {
		outList = outList.PushToReturn(c2)
	}
	return outList
}
func BresenhamLine_Low(c1 CoordInts, c2 CoordInts) CoordList {
	outList := make(CoordList, 0)
	dx := c2.X - c1.X
	dy := c2.Y - c1.Y
	yi := 1
	if dy < 0 {
		yi = -1
		dy = -dy
	}
	D := (2 * dy) - dx
	y := c1.Y
	//fmt.Printf("\tdx,dy:%3d,%3d\n\tyi:%3d y:%d\n\tInitial D:%d\n", dx, dy, yi, y, D)
	//need a conditional here;
	for x := c1.X; x < c2.X; x++ {
		//add to array here

		if D > 0 {
			y = y + yi
			D = D + (2 * (dy - dx))
		} else {
			D = D + (2 * dy)
		}
		//fmt.Printf("\n\t x:%3d y:%3d D:%3d\n", x, y, D)
		outList = append(outList, CoordInts{X: x, Y: y})
	}
	outList = append(outList, c2)
	return outList
}
func BresenHamLine_High(c1 CoordInts, c2 CoordInts) CoordList {
	outList := make(CoordList, 0)
	dx := c2.X - c1.X
	dy := c2.Y - c1.Y
	xi := 1
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	D := (2 * dx) - dy
	x := c1.X

	for y := c1.Y; y < c2.Y; y++ {
		//Add to array here;
		if D > 0 {
			x = x + xi
			D = D + (2 * (dx - dy))
		} else {
			D = D + (2 * dx)
		}
		outList = append(outList, CoordInts{X: x, Y: y})
	}
	outList = append(outList, c2)
	return outList
}

func ManhattanDistanceCulling(c1, c2 CoordInts, Y_Axis_First bool, imat IntMatrix, nums []int) (CoordList, bool) {
	var Outlist CoordList
	temp0 := ManhattanDistance_Basic(c1, c2, Y_Axis_First)
	EndedShort := false
	for _, a := range temp0 {
		if imat.IsValid(a) {
			if IntArrayContains(nums, imat.GetCoordVal(a)) {
				EndedShort = true
				break
			} else {
				Outlist = append(Outlist, a)
			}
		}
	}
	return Outlist, EndedShort
}

func ManhattanDistance_Basic(c1, c2 CoordInts, YFirst bool) CoordList {
	var outList CoordList
	if YFirst {
		temp := GetAllAlongAxis(c1, c2, true)
		outList = append(outList, temp...)
		//Outlist.PrintCordArray()
		temp = GetAllAlongAxis(outList[len(outList)-1], c2, false)
		outList = append(outList, temp...)
	} else {
		temp := GetAllAlongAxis(c1, c2, false)
		outList = append(outList, temp...)
		temp = GetAllAlongAxis(outList[len(outList)-1], c2, true)
		outList = append(outList, temp...)
	}
	if !outList[0].IsEqualTo(c1) {
		outList = outList.PushToFrontThenReturn(c1)
	}
	if !outList[len(outList)-1].IsEqualTo(c2) {
		outList = outList.PushToReturn(c2)
	}
	return outList
}
func GetAllAlongAxis(c1, c2 CoordInts, isYAxis bool) CoordList {
	var Outlist CoordList
	vX, vY := c1.GetDifferenceInInts(c2)
	var tempCoord = c1
	if isYAxis {
		if c2.Y > c1.Y {
			for i := 0; i < vY; i++ {
				tempCoord.Y++
				Outlist = append(Outlist, tempCoord)
			}
		} else {
			for i := 0; i > vY; i-- {
				tempCoord.Y--
				Outlist = append(Outlist, tempCoord)
			}
		}
	} else {
		if c2.X > c1.X {
			for i := 0; i < vX; i++ {
				tempCoord.X++
				Outlist = append(Outlist, tempCoord)
			}
		} else {
			for i := 0; i > vX; i-- {
				tempCoord.X--
				Outlist = append(Outlist, tempCoord)
			}
		}
	}
	return Outlist
}
