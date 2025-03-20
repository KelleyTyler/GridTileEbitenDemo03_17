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
				//fmt.Printf("CLICK ON\n")
				//btn.IsEnabled = !btn.IsEnabled
			} else {
				btn.State = 2
				// btn.IsToggled = true
				//fmt.Printf("CLICK ON_NO %dx\n", btn.BType)
			}
		} else {
			// if btn.BType != 3 && btn.IsToggled {
			// 	fmt.Printf("CLICK OFF\n")
			// 	btn.IsToggled = false
			// }
			btn.isHovered = true
			btn.State = 1
		}
	} else {
		btn.State = 0
	}

}
func (btn *Button) UpdateTwo() bool {
	// if ebitenutil.IsMouseButtonJustReleased(ebiten.MouseButton0) {

	// }
	// Raw_Mouse_X, Raw_Mouse_Y := ebiten.CursorPosition()
	// mouseBool := (Raw_Mouse_X > btn.Coords.X && Raw_Mouse_X < btn.Coords.X+btn.Size.X) && (Raw_Mouse_Y > btn.Coords.Y && Raw_Mouse_Y < btn.Coords.Y+btn.Size.Y)
	// if mouseBool {

	// 	if btn.State > 1 {
	// 		if btn.State < 5 {
	// 			btn.State++
	// 		} else {
	// 			btn.State = 3
	// 		}
	// 	}
	// }
	if btn.IsEnabled {
		if btn.BType == 2 {
			return btn.IsToggled
		} else {
			if btn.State > 1 {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

/* Reasoning that this should depend on the state of the button;
 */
func (btn *Button) CheckIfTriggered() bool {
	return btn.IsEnabled
}

func (btn *Button) GetColor() color.Color {
	//This should return a Color;
	var out int
	if btn.BType == 2 && btn.IsToggled {
		//out = btn.State + 3
		if btn.State > 1 {
			out = 5
		} else {
			out = btn.State + 3
		}
		return btn.Color[out]
	} else {
		if btn.State > 1 {
			out = 2
		} else {
			out = btn.State
		}
		return btn.Color[out]
	}
}

func (btn *Button) InitButton(name, label string, bType int, Pos_X, Pos_Y, BtnWidth, BtnHeight, OffsetX, OffsetY int) {
	btn.Name, btn.Label = name, label
	btn.Offset.X, btn.Offset.Y, btn.Size.X, btn.Size.Y = OffsetX, OffsetY, BtnWidth, BtnHeight
	btn.Coords.X, btn.Coords.Y = Pos_X, Pos_Y
	btn.IsEnabled = true
	btn.IsToggled = false
	btn.isHovered = false
	btn.BType = bType
	btn.Color = []color.Color{color.RGBA{75, 150, 75, 255}, color.RGBA{120, 220, 75, 255}, color.RGBA{140, 240, 100, 255},
		color.RGBA{150, 75, 75, 255}, color.RGBA{220, 120, 75, 255}, color.RGBA{240, 140, 90, 255}}
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
	// out := fmt.Sprintf("%s %t\n", btn.Label, btn.IsToggled)
	out := fmt.Sprintf("%s\n", btn.Label)
	ebitenutil.DebugPrintAt(screen, out, btn.Coords.X, btn.Coords.Y)
	// if sprt.showSimg {
	// 	screen.DrawImage(sprt.animars.GetCurrFrame(), &g.op)
	// 	//screen.DrawImage(&sprt.Simg[sprt.imgArrCurrent], &g.op)
	// }

	// screen.DrawImage(&sprt.Simg[sprt.imgArrCurrent], &g.op)
	// screen.DrawImage(, &opts)
}

/*
---BUTTON PANELS--- Or the basic Button Panel
THE GOAL here is a Button Panel; it will have ButtonPanelStyle as a similar struct that holds a vast majority of it's stats that control for how buttons are place and where
*/
