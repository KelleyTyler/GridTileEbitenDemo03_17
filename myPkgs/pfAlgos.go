package mypkgs

import (
	"fmt"
	"math"
)

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
