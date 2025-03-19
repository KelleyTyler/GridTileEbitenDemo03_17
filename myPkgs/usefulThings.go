package mypkgs

/*
	The purpose of this file is to provide some "USEFUL" but misc things
*/
import "fmt"

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
