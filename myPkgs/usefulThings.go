package mypkgs

/*
	The purpose of this file is to provide some "USEFUL" but misc things
*/
import (
	"fmt"
)

type CoordInts struct {
	X, Y int
}
type CoordFloat64 struct {
	X, Y float64
}
type CoordFloat32 struct {
}

func TestFunc() {
	fmt.Printf("HEY THERE BABE!")
}

func (coord *CoordInts) IsEqualTo(coord2 CoordInts) bool {
	if coord.X == coord2.X && coord.Y == coord2.Y {
		return true
	}
	return false
}

// func (coord CoordInts) Copy() CoordInts {
// 	return
// }

func IntArrayContains(s []int, c int) bool {
	for _, a := range s {
		if a == c {
			return true
		}
	}
	return false
}
func IntArrayContains_giveMeWhat(s []int, c int) (bool, int) {
	for _, a := range s {
		if a == c {
			return true, a
		}
	}
	return false, -1
}
func (coord1 CoordInts) GetDifferenceInInts(coord2 CoordInts) (int, int) {
	return (coord2.X - coord1.X), (coord2.Y - coord1.Y)
}

func (coord1 CoordInts) GetDistance(coord2 CoordInts) float64 {
	//rise over run y2-y1 over x2-x1)
	var x0, y0 int = -1, -1

	if coord1.Y > coord2.Y {
		x0, y0 = coord1.GetDifferenceInInts(coord2)
	} else {
		x0, y0 = coord2.GetDifferenceInInts(coord2)
	}
	return (float64(y0) / float64(x0))
}

func (coord1 CoordInts) AddCoords(coord2 CoordInts) CoordInts {
	return CoordInts{X: coord1.X + coord2.X, Y: coord1.Y + coord2.Y}
}
