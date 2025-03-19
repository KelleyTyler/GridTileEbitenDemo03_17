package mypkgs

import (
	"fmt"
	"image/color"

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// type ButtonState uint8

// const (
// 	BtnEnabled ButtonState = 5
// 	BtnFocused ButtonState = 10
// 	BtnPressed ButtonState = 15
// 	// BtnPressed_Post ButtonState = 12
// )

// type ButtonType uint8

// const (
// 	BtnTypeMomentary ButtonType = 0
// 	BtnTypeToggle    ButtonType = 1
// )

type Button struct {
	Coords    CoordInts
	Offset    CoordInts
	Size      CoordInts
	Color     []color.Color
	Angle     int
	Label     string
	Name      string
	State     int
	BType     int
	IsEnabled bool //not to be confused with active; this is
	isHovered bool
	IsToggled bool
}

func (btn *Button) ToString() string {
	return ""
}

func (btn *Button) PrintString() {
	fmt.Printf("%s\n", btn.ToString())
}

func (btn *Button) Update(Raw_Mouse_X, Raw_Mouse_Y int, isTriggered bool) {
	// Raw_Mouse_X, Raw_Mouse_Y := ebiten.CursorPosition()
	if (Raw_Mouse_X > btn.Coords.X && Raw_Mouse_X < btn.Coords.X+btn.Size.X) && (Raw_Mouse_Y > btn.Coords.Y && Raw_Mouse_Y < btn.Coords.Y+btn.Size.Y) {
		if isTriggered {

			if btn.BType == 2 {
				btn.IsToggled = !btn.IsToggled
				//btn.IsEnabled = !btn.IsEnabled
			} else {
				btn.State = 3
				// btn.IsToggled = true
				// fmt.Printf("CLICK ON\n")
			}
		} else {
			// if btn.BType != 3 && btn.IsToggled {
			// 	fmt.Printf("CLICK OFF\n")
			// 	btn.IsToggled = false
			// }
			btn.isHovered = true
			btn.State = 2
		}
	} else {
		btn.State = 0
	}

}

/* Reasoning that this should depend on the state of the button;
 */
func (btn *Button) CheckIfTriggered() bool {
	return btn.IsEnabled
}

func (btn *Button) GetColor() color.Color {
	//This should return a Color;
	if btn.BType == 2 && btn.IsToggled {
		out := btn.State + 3
		return btn.Color[out]
	} else {
		return btn.Color[btn.State]
	}
}

func (btn *Button) InitButton(Pos_X, Pos_Y, BtnWidth, BtnHeight int) {

}

func (btn *Button) DrawButton(screen *ebiten.Image) {

	// w := btn.Size.X
	// h := btn.Size.Y
	// var opts ebiten.DrawImageOptions
	// opts.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	// opts.GeoM.Rotate(2 * math.Pi * float64(btn.Angle) / float64(180))
	// // g.op.GeoM.Translate(float64(w)/2, float64(h)/2)
	// opts.GeoM.Translate(float64(w)/2, float64(h)/2)
	// opts.GeoM.Translate(float64(btn.Coords.X)+float64(w)/2, float64(btn.Coords.X)+float64(h)/2)
	vector.DrawFilledRect(screen, float32(btn.Coords.X), float32(btn.Coords.Y), float32(btn.Size.X), float32(btn.Size.Y), btn.GetColor(), true)
	vector.StrokeRect(screen, float32(btn.Coords.X), float32(btn.Coords.Y), float32(btn.Size.X), float32(btn.Size.Y), 2.0, color.Black, true)
	out := fmt.Sprintf("%s %t", btn.Label, btn.IsToggled)
	ebitenutil.DebugPrintAt(screen, out, btn.Coords.X, btn.Coords.Y)
	// if sprt.showSimg {
	// 	screen.DrawImage(sprt.animars.GetCurrFrame(), &g.op)
	// 	//screen.DrawImage(&sprt.Simg[sprt.imgArrCurrent], &g.op)
	// }

	// screen.DrawImage(&sprt.Simg[sprt.imgArrCurrent], &g.op)
	// screen.DrawImage(, &opts)
}
