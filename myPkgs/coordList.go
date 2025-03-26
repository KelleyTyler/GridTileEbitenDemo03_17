package mypkgs

import "fmt"

type CoordList []CoordInts

func (cord CoordList) ToString() string {
	retStrng := "COORDAR:\n"
	retStrng += fmt.Sprintf("--SIZE: %d\n", len(cord))
	return retStrng
}

func (cord CoordList) PushToReturn(coord CoordInts) CoordList {
	temp := append(cord, coord)
	return temp
}
func (cord CoordList) PushToFrontThenReturn(coord CoordInts) CoordList {
	temp := make(CoordList, 0)
	temp = append(temp, coord)
	temp = append(temp, cord...)
	return temp
}
func (cord CoordList) PopFromFront() (CoordInts, CoordList) {
	temp := cord[0]
	temp2 := cord.RemovePointFromList(0)
	return temp, temp2
}
func (cord CoordList) PopFromBack() (CoordInts, CoordList) {
	temp := cord[len(cord)-1]
	temp2 := cord.RemovePointFromList(len(cord) - 1)
	return temp, temp2
}

func (coord CoordList) ToCoordArray() []CoordInts {
	outAr := make([]CoordInts, len(coord))
	copy(outAr, coord)
	return outAr
}
func (coord CoordList) FromCoordArray(c []CoordInts) CoordList {
	outList := make(CoordList, len(c))
	copy(outList, c)
	return outList
}
func (cord CoordList) RemoveCoordFromList(coord CoordInts) (CoordList, bool) {
	temp := make(CoordList, 0)
	isThere := false
	for _, c := range cord {
		if !c.IsEqualTo(coord) {
			temp = append(temp, c)
		}
	}
	return temp, isThere
}
func (cord CoordList) RemovePointFromList(num int) CoordList {
	temp := make(CoordList, 0)
	for i, _ := range cord {
		if i != num {
			temp = append(temp, cord[i])
		}
	}
	return temp
}
func (cord CoordList) CountInstances(coord CoordInts) int {
	temp := 0
	for _, c := range cord {
		if c.IsEqualTo(coord) {
			temp++
		}
	}
	return temp
}

func (cord CoordList) PrintCordArray() {
	fmt.Print("\n\n------------------------\n")
	for i, c := range cord {
		fmt.Printf("%2d: {%3d %3d}", i, c.X, c.Y)
		// if i%1 == 0 {
		// 	fmt.Print("\n")
		// } else {
		// 	fmt.Print("\t")
		// }
		fmt.Print("\n")

	}
	fmt.Print("\n------------------------\n")
}

func (cord CoordList) SortDescOnX() CoordList {
	temp := make([]CoordInts, len(cord))
	copy(temp, cord)
	var tempcord CoordInts
	if len(temp) > 1 {
		// var halfTemp = (len(temp) / 2)
		for range temp {
			for i := 1; i < (len(temp)); i++ {
				// q := halfTemp + i
				if temp[i].X > temp[i-1].X {
					tempcord = temp[i]
					temp[i] = temp[i-1]
					temp[i-1] = tempcord
				} else if temp[i].X == temp[i-1].X {
					if temp[i].Y > temp[i-1].Y {
						tempcord = temp[i]
						temp[i] = temp[i-1]
						temp[i-1] = tempcord
					}
				}
			}
		}
	}
	return temp
}

/*
CoordList.RemoveDuplicates
this should be done;
this will remove duplicates;
*/
func (cord CoordList) RemoveDuplicates() CoordList {
	temp := make(CoordList, len(cord))
	copy(temp, cord)
	temp = temp.SortDescOnX()
	for i := 1; i < len(temp); i++ {
		if temp[i].IsEqualTo(temp[i-1]) {
			temp[i-1] = CoordInts{X: -1, Y: -1}
		}
	}
	temp, _ = temp.RemoveCoordFromList(CoordInts{X: -1, Y: -1})
	// temp2 := make(CoordList, 0)
	// for _, c := range cord {

	// }
	return temp
}

//	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//		s[i], s[j] = s[j], s[i]
//	}
func (cord CoordList) FlipOrder() CoordList {
	temp := make(CoordList, len(cord))
	copy(temp, cord)
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return temp
}
