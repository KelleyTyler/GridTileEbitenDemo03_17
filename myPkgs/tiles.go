package mypkgs

import (
	"fmt"
	"math/rand"

	//"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Tile struct {
	passable bool
	position CoordInts
	img      *ebiten.Image
	value    int
	name     string
}
type TileGrid struct {
}

func (t *Tile) ToString() string {
	return fmt.Sprintf("name: %s\nPOS:%3d,%3d\nPassable:%t VALUE: %2d", t.name, t.position.X, t.position.Y, t.passable, t.value)
}
func (t *Tile) PrintTile() {
	fmt.Printf("%s", t.ToString())
}

// func (t *Tile)GetTILE() {}

func (imat IntMatrix) DrawGridTile(screen *ebiten.Image, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int) {
	colorA := color.RGBA{255, 0, 0, 255}
	colorB := color.RGBA{0, 0, 255, 255}
	colorC := color.RGBA{0, 255, 0, 255}
	colorD := color.RGBA{0, 255, 255, 255}
	colorE := color.RGBA{255, 255, 0, 255}
	test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	for y, _ := range imat {
		for x, b := range imat[y] {

			switch b {
			case 0:
				vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorA, false)

			case 1:
				vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorB, false)
			case 2:
				vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorC, false)
			case 3:
				vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorD, false)
			case 4:
				vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colorE, false)
			}
			vector.StrokeRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), 2.0, color.Black, false)
		}
	}
	//vector.StrokeRect(screen, float32(OffsetX-0), float32(OffsetY-0), float32(test1X-0), float32(test1Y+0), 2.0, color.RGBA{210, 153, 100, 255}, true) //0, 179, 100, 255
	vector.StrokeRect(screen, float32(OffsetX-3), float32(OffsetY-3), float32(test1X-OffsetX-GapX+6), float32(test1Y-OffsetY-GapY+6), 4.0, color.RGBA{0, 50, 50, 255}, true)
	//vector.StrokeRect(screen, float32(OffsetX-0), float32(OffsetY-0), float32(test1X-OffsetX), float32(test1Y-OffsetY), 2.0, color.RGBA{0, 253, 100, 255}, true)
}

func (imat IntMatrix) ChangeValOnMouseEvent(Raw_Mouse_X int, Raw_Mouse_Y int, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int, cycleStart int, cycleEnd int) {

	//first figure out what coord the mouse button was in (if any);
	//var NewPoint image.Point = image.Point{0, 0}
	// test0X := OffsetX
	// test0Y := OffsetY
	// test1X := (len(imat[0]) * (tileH + GapX)) + OffsetX
	// test1Y := (len(imat) * (tileW + GapY)) + OffsetY
	// test1X := (len(imat[0]) * (tileW + GapX)) + OffsetX
	// test1Y := (len(imat) * (tileH + GapY)) + OffsetY
	test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	// test1X := ((len(imat[0]) * GapX) + (len(imat[0]) * tileW)) + OffsetX
	// test1Y := ((len(imat) * GapY) + (len(imat) * tileH)) + OffsetY
	if (Raw_Mouse_X > OffsetX && Raw_Mouse_X < test1X-GapX) && (Raw_Mouse_Y > OffsetY && Raw_Mouse_Y < test1Y-GapY) {

		var mX = float32(Raw_Mouse_X-OffsetX) / float32(tileW+GapX) //float32(test1X)
		var mY = float32(Raw_Mouse_Y-OffsetY) / float32(tileH+GapY) //float32(test1Y)
		var mXi = int(mX)
		var mYi = int(mY)
		// fmt.Printf("IN THE BOX\nRM_: %d,%d\n", Raw_Mouse_X, Raw_Mouse_Y)
		// fmt.Printf("test %d,%d\n", test1Y, test1X)
		//----
		// fmt.Printf("RMO: %d, %d\nm_f: %f %f\n", Raw_Mouse_X-OffsetX, Raw_Mouse_Y-OffsetY, mX, mY)
		// fmt.Printf("M_I: %d, %d\nm_i: %d,%d\n", (Raw_Mouse_X-OffsetX)/(tileH+GapX), (Raw_Mouse_Y-OffsetY)/(tileH+GapY), mXi, mYi)

		// fmt.Printf("A %d,%d\n", (tileW*mXi)+(mXi*GapX), (tileH*mYi)+(mYi*GapY))                           //inner bounds of each rectangle
		// fmt.Printf("B %d,%d\n", (tileW*mXi)+(mXi*GapX)+tileW, (tileH*mYi)+(mYi*GapY)+tileH)               //outer bounds of each rectangle;
		// fmt.Printf("C %d,%d\n", (tileW*mXi)+(mXi*GapX)+(tileW+GapX), (tileH*mYi)+(mYi*GapY)+(tileH+GapY)) //gaps ending\
		//-----------------------------------------------
		mXo, mYo := (Raw_Mouse_X - OffsetX), (Raw_Mouse_Y - OffsetY)
		mXi_01 := (tileW * mXi) + (mXi * GapX)
		mYi_01 := (tileH * mYi) + (mYi * GapY)

		mXi_02 := (tileW * mXi) + (mXi * GapX) + tileW
		mYi_02 := (tileH * mYi) + (mYi * GapY) + tileH
		//--------------------------------------------------------------------
		// fmt.Printf("m_o: %d,%d\n", mXo, mYo)
		// fmt.Printf("LENS: %d %d\n", len(imat), len(imat[0]))
		// fmt.Printf("TESTS:: %d %d \n", test1Y, test1X)
		// fmt.Printf("A: %d,%d\n", mXi_01, mYi_01)                                                               //inner bounds of each rectangle
		// fmt.Printf("B: %d,%d\n", mYi_02, mYi_02)                                                               //outer bounds of each rectangle;
		// fmt.Printf("C: %d,%d\n\n\n", (tileW*mXi)+(mXi*GapX)+(tileW+GapX), (tileH*mYi)+(mYi*GapY)+(tileH+GapY)) //gaps ending\
		if (((mXo) > (mXi_01)) && (mXo) < mXi_02) && ((mYo) > (mYi_01) && mYo < mYi_02) {
			//change the coord;
			if imat[mYi][mXi] < cycleEnd {
				imat[mYi][mXi] += 1
			} else {
				imat[mYi][mXi] = cycleStart
			}
		}
		// if (((Raw_Mouse_X - OffsetX) > ((tileW * mXi) + (mXi * GapX))) && (Raw_Mouse_X-OffsetX) < mXi_02) && ((Raw_Mouse_Y-OffsetY) > ((tileH*mYi)+(mYi*GapY)) && (Raw_Mouse_Y-OffsetY) < mYi_02) {
		// 	//change the coord;
		// 	if imat[mYi][mXi] < cycleEnd {
		// 		imat[mYi][mXi] += 1
		// 	} else {
		// 		imat[mYi][mXi] = cycleStart
		// 	}
		// }
	}
	// else {
	// 	fmt.Printf("NOT IN THE BOX\nRM_: %d,%d\n", Raw_Mouse_X, Raw_Mouse_Y)
	// 	fmt.Printf("test %d,%d\n", test1X, test1Y)
	// }

	// var mX = Raw_Mouse_X - OffsetX
	// var mY = Raw_Mouse_Y - OffsetY

}

type CoordArray []CoordInts

func (imat *IntMatrix) ActiveMagic(cord CoordArray) {

}
func (cord CoordArray) ToString() string {
	retStrng := fmt.Sprintf("COORDAR:\n")
	retStrng += fmt.Sprintf("--SIZE: %d\n", len(cord))
	return retStrng
}
func (imat IntMatrix) GetCoordVal(cord CoordInts) int {
	fmt.Printf("GET COORD VAL %d %d\n\n", cord.X, cord.Y)
	return imat[cord.Y][cord.X]

}

// func (imat IntMatrix) CanDo(cord CoordInts) bool {
// 	// if cord.X < 0 || cord.X > len(imat[0]) {
// 	// 	return false
// 	// } else {

//		// }
//		if imat[cord.Y][cord.X] != 0 {
//			return true
//		} else {
//			return false
//		}
//		//return false
//	}
func (imat IntMatrix) InField(cord CoordInts) bool {
	if cord.X > -1 && cord.X < len(imat[0])-3 && cord.Y > -1 && cord.Y < len(imat)-3 {
		return true
	} else {
		return false
	}
}

func (imat IntMatrix) Process(cord CoordArray, tickLimit, failureLimit int) CoordArray {
	// var fails int = 0
	// var tempAr CoordArray = make(CoordArray, 0)
	// tempAr.AddToListIfNotAlready(cord[0])
	for i := 0; i < tickLimit; i++ {
		temp := imat.RandomCoordFromList(cord)
		imat[temp.Y][temp.X] = 1
		cord = cord.RemoveCoordFromList(temp)
		cord = imat.AddToList(temp, cord)
		//temp := coord.AddToListIfNotAlready(CoordInts{0, 0})
		for i, c := range cord {
			for j, q := range cord {
				if c == q && i != j {
					fmt.Printf("List SIZE: %d\n", len(cord))
					cord = cord.RemoveCoordFromList(q)
					fmt.Printf("List SIZE: %d\n", len(cord))
				} else {
					imat[c.X][c.Y] = 2
				}
			}
		}
		//return temp
		//fmt.Printf("\n%d---%d, %d- LONG: %d\n", i, temp.X, temp.Y, len(cord))
		// if true {
		// 	fails++
		// }
	}
	imat.Hilight_FromCoord(cord)

	return cord
}

func (imat IntMatrix) CleanItUp(cord CoordArray) CoordArray {
	//var temp CoordArray
	temp := cord.AddToListIfNotAlready(CoordInts{0, 0})
	for i, c := range cord {
		for j, q := range cord {
			if c == q && i != j {
				fmt.Printf("List SIZE: %d\n", len(temp))
				temp = cord.RemoveCoordFromList(q)
				fmt.Printf("List SIZE: %d\n", len(temp))
			} else {
				imat[c.X][c.Y] = 2
			}
		}
	}
	return temp
}

func (imat IntMatrix) Hilight_FromCoord(coord CoordArray) {
	for _, c := range coord {
		imat[c.X][c.Y] = 2
	}
}
func (imat IntMatrix) AddToList(cordInt CoordInts, coList CoordArray) CoordArray {
	for i := 0; i < 8; i++ {
		temp, valid, value := imat.GetFourths(cordInt, i)
		fmt.Printf("-%d-%d-", i, len(coList))
		if valid {
			//fmt.Printf("-VALID-")
			if value != 1 && value != -1 {
				//fmt.Printf("IT DOESN'T----")
				if imat[temp.X][temp.Y] != 1 {
					coList = coList.AddToListIfNotAlready(temp)
					if imat[temp.X][temp.Y] == 0 {
						imat[temp.X][temp.Y] = 2
					} else if !coList.CheckIfAlready(temp) {
						imat[temp.X][temp.Y] = 3
					}
				} else {
					coList = coList.RemoveCoordFromList(temp)
				}
				//fmt.Printf("s: %d ", len(coList))
			} else {
				fmt.Printf("NO THANK YOU\n")
			}
		}
		// if i == 7 {
		// 	fmt.Printf("-%d, %d-", temp.X, temp.Y)
		// }
	}
	fmt.Printf("\n")
	return coList
}

// needs to have rules for adding to the list;
func (coList CoordArray) AddToListIfNotAlready(cordInt CoordInts) CoordArray {
	if !coList.CheckIfAlready(cordInt) {
		// coListCopy := make([]CoordInts, len(coList))
		coList = append(coList, cordInt)
	} else {
		fmt.Printf("WE'VE ALREADY GOT IT!\n\n")
		coList = coList.RemoveCoordFromList(cordInt)
	}

	return coList
}

/*
GetFourths

	0= north
	1= northEast
	2= East
	3= SouthEast
	4= south
	5=southwest
	6=west
	7=northwest
	--------------------------------------
	1111414	Unrelated to above this is showing the grades; 4 is a solid wall;
	0000001 0 is an open path
	1111404 1 is an unchecked wall 2 is a wall with potential 3 is a maybe
*/
func (imat IntMatrix) GetFourths(cord CoordInts, num int) (CoordInts, bool, int) {
	var temp CoordInts = CoordInts{cord.X, cord.Y}
	switch num {
	case 0:
		temp.Y -= 1
	case 1:
		temp.Y -= 1
		temp.X += 1
	case 2:
		temp.X += 1
	case 3:
		temp.X += 1
		temp.Y += 1
	case 4:
		temp.Y += 1
	case 5:
		temp.Y += 1
		temp.X -= 1
	case 6:
		temp.X -= 1
	case 7:
		temp.Y -= 1
		temp.X -= 1
	}
	if imat.InField(temp) {
		return temp, true, imat.GetCoordVal(temp)
	}
	return temp, false, -1
}

func (imat IntMatrix) RandomCoord() {
	var cord CoordInts = CoordInts{X: 0, Y: 0}
	max := len(imat[0])
	may := len(imat)
	for {
		//randomRange := rand.Intn(max-min+1) + min
		cord.X = rand.Intn(max-0+1) + 0
		cord.Y = rand.Intn(may-0+1) + 0
		if imat.GetCoordVal(cord) == 0 {

			fmt.Printf("MAGICAL MYSTERY 242/n")
			break
		} else {
			fmt.Printf("MAGICAL MYSTERY TOUR/n")
		}
	}
}
func (imat IntMatrix) SETEVERYTHINGTOGREEN(coord CoordArray) CoordArray {
	for i, _ := range imat {
		for j, x := range imat[i] {
			if x > 1 {
				// fmt.Printf("WE GOT EM")
				imat[i][j] = 1
			}
		}
	}
	temp := coord.AddToListIfNotAlready(CoordInts{0, 0})
	for i, c := range temp {
		for j, q := range temp {
			if c == q && i != j {
				fmt.Printf("WE GOT IT %d %d %d %d\n\n", c.X, c.Y, q.X, q.Y)
				temp = temp.RemoveCoordFromList(q)
			} else {
				imat[c.X][c.Y] = 2
			}
		}
		imat[c.X][c.Y] = 3
	}
	return temp
}

func (coList CoordArray) CheckIfAlready(coord2 CoordInts) bool {
	//cX_Size := len(coList)
	for _, c := range coList {
		if c.IsEqualTo(coord2) {
			return true
		}
	}
	return false
}

func (imat IntMatrix) RandomCoordFromList(coList CoordArray) CoordInts {
	loner := len(coList) - 1
	var ranNum int = 0
	if loner < 2 {
		ranNum = 0
	} else {
		ranNum = rand.Intn(loner-0+1) + 0

	}
	mee := coList[ranNum]
	return mee
}
func (coList CoordArray) RemovePointFromList(num int) {
	var temp CoordArray = make(CoordArray, 0)
	for i, c := range coList {
		if i != num {
			temp = append(temp, c)
		}
	}
	// temp.
	//coListCopy := make([]CoordInts, len(coList))
	copy(coList, temp)
}
func (coList CoordArray) RemoveCoordFromList(coo CoordInts) CoordArray {
	var temp CoordArray = make(CoordArray, 0)
	for _, c := range coList {
		if !coo.IsEqualTo(c) {
			temp = append(temp, c)
		}
	}
	// temp.
	//coListCopy := make([]CoordInts, len(coList))
	return temp
}

// func (imat IntMatrix) EatShitR(){

// }
