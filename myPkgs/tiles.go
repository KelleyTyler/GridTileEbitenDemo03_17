package mypkgs

import (
	"fmt"

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
func (imat IntMatrix) DrawAGridTile(screen *ebiten.Image, coord CoordInts, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int, clr color.Color, aa bool) {
	vector.DrawFilledRect(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), float32(tileW), float32(tileH), clr, aa)

	vector.StrokeRect(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), float32(tileW), float32(tileH), 2.0, color.Black, aa)
	vector.StrokeLine(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), float32((tileW*coord.X)+(GapX*coord.X)+OffsetX+tileW), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY+tileH), 2.0, color.Black, aa)

}

func (imat IntMatrix) DrawAGridTile_With_Line(screen *ebiten.Image, coord CoordInts, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int, clr0, clr1, clr2 color.Color, lineThick float32, aa bool) {
	vector.DrawFilledRect(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), float32(tileW), float32(tileH), clr0, aa)
	vector.StrokeRect(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), float32(tileW), float32(tileH), 2.0, color.Black, aa)
	vector.StrokeLine(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), float32((tileW*coord.X)+(GapX*coord.X)+OffsetX+tileW), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY+tileH), lineThick, clr1, aa)
	vector.StrokeLine(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY+tileH), float32((tileW*coord.X)+(GapX*coord.X)+OffsetX+tileW), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), lineThick, clr2, aa)
	//vector.StrokeRect(screen, float32((tileW*coord.X)+(GapX*coord.X)+OffsetX), float32((tileH*coord.Y)+(GapY*coord.Y)+OffsetY), float32(tileW), float32(tileH), 2.0, color.Black, aa)
}

func (imat IntMatrix) DrawGridTiles(screen *ebiten.Image, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int, colors []color.Color) {
	test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	for y, _ := range imat {
		for x, b := range imat[y] {
			vector.DrawFilledRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), colors[b], false)

			color0 := color.RGBA{12, 12, 12, 100} //color.Black
			vector.StrokeRect(screen, float32((tileW*x)+(GapX*x)+OffsetX), float32((tileH*y)+(GapY*y)+OffsetY), float32(tileW), float32(tileH), 0, color0, false)
		}
	}
	//vector.StrokeRect(screen, float32(OffsetX-0), float32(OffsetY-0), float32(test1X-0), float32(test1Y+0), 2.0, color.RGBA{210, 153, 100, 255}, true) //0, 179, 100, 255
	vector.StrokeRect(screen, float32(OffsetX-3), float32(OffsetY-3), float32(test1X-OffsetX-GapX+6), float32(test1Y-OffsetY-GapY+6), 4.0, color.RGBA{0, 50, 50, 255}, true)
	//vector.StrokeRect(screen, float32(OffsetX-0), float32(OffsetY-0), float32(test1X-OffsetX), float32(test1Y-OffsetY), 2.0, color.RGBA{0, 253, 100, 255}, true)
}
func (imat IntMatrix) GetCoordOfMouseEvent(Raw_Mouse_X int, Raw_Mouse_Y int, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int) (int, int, bool) {
	test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	var mXi = -1
	var mYi = -1
	var isOnTile = false
	if (Raw_Mouse_X > OffsetX && Raw_Mouse_X < test1X-GapX) && (Raw_Mouse_Y > OffsetY && Raw_Mouse_Y < test1Y-GapY) {
		var mX = float32(Raw_Mouse_X-OffsetX) / float32(tileW+GapX) //float32(test1X)
		var mY = float32(Raw_Mouse_Y-OffsetY) / float32(tileH+GapY) //float32(test1Y)
		mXi = int(mX)
		mYi = int(mY)
		mXo, mYo := (Raw_Mouse_X - OffsetX), (Raw_Mouse_Y - OffsetY)
		mXi_01 := (tileW * mXi) + (mXi * GapX)
		mYi_01 := (tileH * mYi) + (mYi * GapY)

		mXi_02 := (tileW * mXi) + (mXi * GapX) + tileW
		mYi_02 := (tileH * mYi) + (mYi * GapY) + tileH
		//fmt.Printf("0-->%6d,%6d\t%6d,%6d\t rec start: %6d %6d\t %6d %6d \n", Raw_Mouse_X, Raw_Mouse_Y, mXo, mYo, mXi_01, mYi_01, mXi_02, mYi_02)
		if (((mXo) > (mXi_01)) && (mXo) < mXi_02) && ((mYo) > (mYi_01) && mYo < mYi_02) {
			isOnTile = true
		}
	}
	return mXi, mYi, isOnTile
}
func (imat IntMatrix) GetCoordOfMouseEvent_Scalable(Raw_Mouse_X int, Raw_Mouse_Y int, scale float64, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int) (int, int, bool) {
	// test0X := scale * float64(OffsetY)
	// test0Y := scale * float64(OffsetY)
	test0X := scale * float64(OffsetY)
	test0Y := scale * float64(OffsetY)
	test1Xa := scale * float64(((len(imat[0])*tileW)+(len(imat[0])*GapX))+OffsetX)
	test1Ya := scale * float64(((len(imat)*tileH)+(len(imat)*GapY))+OffsetY)
	// test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	// test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	var mXi = -1
	var mYi = -1
	var isOnTile = false
	if (float64(Raw_Mouse_X) > test0X && float64(Raw_Mouse_X) < test1Xa-float64(GapX)) && (float64(Raw_Mouse_Y) > test0Y && float64(Raw_Mouse_Y) < test1Ya-float64(GapX)) {
		// var mX = float32(Raw_Mouse_X-OffsetX) / float32(tileW+GapX) //float32(test1X)
		// var mY = float32(Raw_Mouse_Y-OffsetY) / float32(tileH+GapY) //float32(test1Y)
		var mX = (-float64(Raw_Mouse_X) + test1Xa) / (float64(tileW+GapX) * -scale) //float32(test1X)
		var mY = (-float64(Raw_Mouse_Y) + test1Ya) / (float64(tileH+GapY) * -scale) //float32(test1Y)

		mXi = int(mX) + len(imat[0]) - 1
		mYi = int(mY) + len(imat) - 1
		mXo, mYo := int((Raw_Mouse_X-OffsetX))/int(scale*scale), int((Raw_Mouse_Y-OffsetY))/int(scale*scale)
		mXi_01 := int((tileW*mXi)+(mXi*GapX)) / int(scale)
		mYi_01 := int((tileH*mYi)+(mYi*GapY)) / int(scale)
		// fmt.Printf("--->%d, %d %f,%f\n", mXo, mYo, float64(mXo)-(float64(OffsetX)*scale), float64(mYo)-(float64(OffsetY)*scale))
		//The scale is off; the cursor occassionally gets the wrong point;
		// at the very least it's not working when it's
		mXi_02 := ((tileW*mXi)*int(scale) + (mXi*GapX)*int(scale) + tileW*int(scale)) / int(scale)
		mYi_02 := ((tileH*mYi)*int(scale) + (mYi*GapY)*int(scale) + tileH*int(scale)) / int(scale)
		fmt.Printf("--->%6d,%6d\t%6.1d,%6.1d\t rec start: %6.d,%6.d \t %6d,%6d--\n", Raw_Mouse_X/int(scale), Raw_Mouse_Y/int(scale), mXo, mYo, mXi_01, mYi_01, mXi_02, mYi_02)
		if (((mXo) > int(mXi_01)) && (mXo) < int(mXi_02)) && ((mYo) > int(mYi_01) && mYo < int(mYi_02)) { //float64(
			isOnTile = true
		}
	}
	return mXi, mYi, isOnTile
}
func (imat IntMatrix) IsCursorInBounds(OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int) bool {
	Raw_Mouse_X, Raw_Mouse_Y := ebiten.CursorPosition()
	test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	return ((Raw_Mouse_X > OffsetX && Raw_Mouse_X < test1X-GapX) && (Raw_Mouse_Y > OffsetY && Raw_Mouse_Y < test1Y-GapY))
}

func (imat IntMatrix) GetCursorBounds(OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int) (int, int) {
	test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	return test1X, test1Y
}

func (imat IntMatrix) ChangeValOnMouseEvent(Raw_Mouse_X int, Raw_Mouse_Y int, OffsetX int, OffsetY int, tileW int, tileH int, GapX int, GapY int, cycleStart int, cycleEnd int, makeChange bool) (int, int) {

	test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	var mXi = -1
	var mYi = -1
	if (Raw_Mouse_X > OffsetX && Raw_Mouse_X < test1X-GapX) && (Raw_Mouse_Y > OffsetY && Raw_Mouse_Y < test1Y-GapY) {

		var mX = float32(Raw_Mouse_X-OffsetX) / float32(tileW+GapX) //float32(test1X)
		var mY = float32(Raw_Mouse_Y-OffsetY) / float32(tileH+GapY) //float32(test1Y)
		mXi = int(mX)
		mYi = int(mY)
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
		if (((mXo) > (mXi_01)) && (mXo) < mXi_02) && ((mYo) > (mYi_01) && mYo < mYi_02) && makeChange {
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
	return mXi, mYi
}
func (imat IntMatrix) GetCoordVal(cord CoordInts) int {
	//fmt.Printf("GET COORD VAL %d %d\n\n", cord.X, cord.Y)
	return imat[cord.Y][cord.X]

}

func (imat IntMatrix) IsValid(cord CoordInts) bool {
	if (cord.X > -1 && cord.X < len(imat[0])) && (cord.Y > -1 && cord.Y < len(imat)) {
		return true
	}
	return false
}
func (imat IntMatrix) IsValid_With_Constant_Buffer(cord CoordInts, buffer int) bool {
	if (cord.X > -1+buffer && cord.X < len(imat[0])-buffer) && (cord.Y > -1+buffer && cord.Y < len(imat)-buffer) {
		return true
	}
	return false
}
func (imat IntMatrix) IsValid_WithDir_Buffer(cord CoordInts, buffer [4]int) bool {
	if (cord.X > -1+buffer[3] && cord.X < len(imat[0])-buffer[1]) && (cord.Y > -1+buffer[0] && cord.Y < len(imat)-buffer[2]) {
		return true
	}
	return false
}

// func (imat IntMatrix) ClearAnArea(c1_X, c1_Y, c2_X, c2_Y, val int) {

// 	if c1_Y < len(imat) && c1_Y > -1 {
// 		if c2_Y > -1 && c2_Y < len(imat) {
// 			for i := c1_Y; i < c2_Y; i++ {
// 				for j := c1_X; j < c2_X; j++ {
// 					imat[i][j] = val
// 				}
// 			}
// 		}
// 	}
// }

func (imat IntMatrix) ClearAnArea(c0_X, c0_Y, c1_X, c1_Y, val int) {

	if c0_Y < len(imat) && c0_Y > -1 {
		if c1_Y > -1 && c1_Y < len(imat) {
			for i := c0_Y; i < c1_Y; i++ {
				for j := c0_X; j < c1_X; j++ {
					imat[i][j] = val
				}
			}
		}
	}
}

func (imat IntMatrix) DrawListAsTiles(screen *ebiten.Image, cord CoordList, offsetX, offsetY int, tileW, tileH int, gapX, gapY int, clr0 color.Color, aa bool) {
	for _, a := range cord {
		imat.DrawAGridTile(screen, a, offsetX, offsetY, tileW, tileH, gapX, gapY, clr0, aa)
	}
}
func (imat IntMatrix) DrawListAsTiles_withLines(screen *ebiten.Image, cord CoordList, offsetX, offsetY int, tileW, tileH int, gapX, gapY int, clr0, clr1, clr2 color.Color, lineThick float32, aa bool) {
	for _, a := range cord {
		imat.DrawAGridTile_With_Line(screen, a, offsetX, offsetY, tileW, tileH, gapX, gapY, clr0, clr1, clr2, lineThick, aa)
	}
}
