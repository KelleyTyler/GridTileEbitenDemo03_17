package mypkgs

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"

	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	//"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
//
//	BtnTypeMomentary ButtonType = 0
//	BtnTypeToggle    ButtonType = 1
//
// )

/*
UI_Helper:
The
*/
type UI_Helper struct {
	IsInitialized               bool
	Button_Colors               []color.Color
	Btn_Sounds                  [][]byte
	Btn_TextSrc                 *text.GoTextFaceSource
	Btn_Text_Mono, Btn_Text_Reg text.Face

	SoundSys *AudioThing
}

// func (ui_Helper *UI_Helper) Init_Default(sound *AudioThing) error {

func (ui_Helper *UI_Helper) Init_Default(sound *AudioThing) error {
	ui_Helper.Button_Colors = []color.Color{color.RGBA{75, 150, 75, 255}, color.RGBA{120, 220, 75, 255}, color.RGBA{140, 240, 100, 255},
		color.RGBA{150, 75, 75, 255}, color.RGBA{220, 120, 75, 255}, color.RGBA{240, 140, 90, 255}}

	ui_Helper.IsInitialized = true
	var err error
	ui_Helper.Btn_TextSrc, err = text.NewGoTextFaceSource(bytes.NewReader(gomono.TTF))
	if err != nil {
		log.Fatal("err: ", err)
	}
	ui_Helper.Btn_Text_Mono = &text.GoTextFace{
		Source: ui_Helper.Btn_TextSrc,
		Size:   20,
	}
	ui_Helper.Btn_TextSrc, err = text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal("err: ", err)
	}
	ui_Helper.Btn_Text_Reg = &text.GoTextFace{
		Source: ui_Helper.Btn_TextSrc,
		Size:   20,
	}

	ui_Helper.SoundSys = sound
	ui_Helper.InitSounds()
	return nil
}

func (ui_Helper *UI_Helper) PlaySound(sound_num int) {
	if sound_num < int(len(ui_Helper.Btn_Sounds)) {
		ui_Helper.SoundSys.PlayByte(ui_Helper.Btn_Sounds[sound_num])
	}
}

func (ui_Helper *UI_Helper) InitSounds() { //1200
	ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 0, 110, []float32{1.0}, []float32{0.0750000}))
	ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 10, 110, []float32{1.0}, []float32{0.0750000}))
	ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 15, 110, []float32{1.0}, []float32{0.0750000}))
	ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 20, 110, []float32{1.0}, []float32{0.0750000}))
	ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 25, 110, []float32{1.0}, []float32{0.0750000}))
	ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 25, 110, []float32{1.0}, []float32{0.0750000}))
}

func (ui_Helper *UI_Helper) ToString() string {
	outstrng := fmt.Sprintf("UI HELPER: sounds %d\n", len(ui_Helper.Btn_Sounds))
	return outstrng
}

/*
 */
func IsMouseOverPos(adj_x, adj_y int, position, size CoordInts) bool {
	if mx, my := ebiten.CursorPosition(); (mx > position.X+adj_x && mx < position.X+size.X+adj_x) && (my > position.Y+adj_y && my < position.Y+size.Y+adj_y) {
		return true
	} else {
		return false
	}
}
func (btn *Button) InitButton(name, label string, uiHelpr *UI_Helper, bType int, Pos_X, Pos_Y, BtnWidth, BtnHeight, OffsetX, OffsetY int) {
	btn.Name, btn.Label = name, label
	btn.Offset.X, btn.Offset.Y, btn.Size.X, btn.Size.Y = OffsetX, OffsetY, BtnWidth, BtnHeight
	btn.Coords.X, btn.Coords.Y = Pos_X, Pos_Y
	btn.IsEnabled = true
	btn.IsToggled = false
	btn.isHovered = false
	btn.Helper = uiHelpr
	btn.BType = bType
	// btn.Color = []color.Color{color.RGBA{75, 150, 75, 255}, color.RGBA{120, 220, 75, 255}, color.RGBA{140, 240, 100, 255},
	// color.RGBA{150, 75, 75, 255}, color.RGBA{220, 120, 75, 255}, color.RGBA{240, 140, 90, 255}}
	btn.Color = btn.Helper.Button_Colors
}

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

	// PointingBool *bool

	Helper *UI_Helper
}

func (btn *Button) ToString() string {
	return ""
}

func (btn *Button) PrintString() {
	fmt.Printf("%s\n", btn.ToString())
}

func (btn *Button) isMouseOverPos(adj_x, adj_y int) bool {
	if mx, my := ebiten.CursorPosition(); (mx > btn.Coords.X+adj_x && mx < btn.Coords.X+btn.Size.X+adj_x) && (my > btn.Coords.Y+adj_y && my < btn.Coords.Y+btn.Size.Y+adj_y) {
		return true
	} else {
		return false
	}

}
func (btn *Button) Update3() bool { //no clue if this works;
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton(0)) && btn.isMouseOverPos(0, 0) { //no clue if this works;
		//fmt.Printf("TICK TICK TICK\n")
		if btn.BType == 2 {
			btn.IsToggled = !btn.IsToggled
			btn.Helper.PlaySound(3)
			//fmt.Printf("TICK 2 2\n")
			return btn.IsToggled
		} else {
			btn.Helper.PlaySound(2)
			btn.State = 2
			return true
		}
	} else if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButton(0)) && btn.isMouseOverPos(0, 0) {
		if btn.BType == 2 {
			btn.isHovered = true
			btn.State = 1
			return btn.IsToggled
		} else {
			btn.isHovered = true
			btn.State = 1
			return false
		}
	} else {
		if btn.IsEnabled {
			if btn.BType == 2 {
				// if btn.IsToggled {
				// 	fmt.Printf("TICK 4 4\n")
				// }
				btn.State = 0
				return btn.IsToggled
			} else {
				if btn.State > 1 {
					return true
				} else {
					btn.State = 0
					return false
				}
			}
		} else {
			return false
		}
	}
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

func (btn *Button) InitButton_deprecated(name, label string, bType int, Pos_X, Pos_Y, BtnWidth, BtnHeight, OffsetX, OffsetY int) {
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
func (btn *Button) UpdateTwo() bool {
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

func (btn *Button) ChangeLabel(strng string) {
	btn.Label = strng
}
func (btn *Button) DrawButton(screen *ebiten.Image) {

	// w := btn.Size.X
	// h := btn.Size.Y
	scaler := 2.0
	// var opts ebiten.DrawImageOptions
	// opts.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	// opts.GeoM.Rotate(2 * math.Pi * float64(btn.Angle) / float64(180))
	// // g.op.GeoM.Translate(float64(w)/2, float64(h)/2)
	// opts.GeoM.Translate(float64(w)/2, float64(h)/2)
	// opts.GeoM.Translate(float64(btn.Coords.X)+float64(w)/2, float64(btn.Coords.X)+float64(h)/2)
	vector.DrawFilledRect(screen, float32(btn.Coords.X), float32(btn.Coords.Y), float32(btn.Size.X), float32(btn.Size.Y), btn.GetColor(), true)
	vector.StrokeRect(screen, float32(btn.Coords.X), float32(btn.Coords.Y), float32(btn.Size.X), float32(btn.Size.Y), 2.0, color.Black, true)
	//out := fmt.Sprintf("%s %t\n", btn.Label, btn.IsToggled)
	out := fmt.Sprintf("%s\n", btn.Label)
	// if btn.PointingBool != nil {
	// 	out += fmt.Sprintf("%t\n", *btn.PointingBool)
	// }	btn.Helper.Btn_Text.
	tops := &text.DrawOptions{}

	tops.GeoM.Translate(float64(btn.Coords.X+4)*scaler, float64(btn.Coords.Y+4)*scaler)
	tops.GeoM.Scale(1/scaler, 1/scaler)
	tops.ColorScale.ScaleWithColor(color.White)
	tops.LineSpacing = float64(20)
	text.Draw(screen, out, btn.Helper.Btn_Text_Reg, tops)
	//ebitenutil.DebugPrintAt(screen, out, btn.Coords.X, btn.Coords.Y)
	// if sprt.showSimg {
	// 	screen.DrawImage(sprt.animars.GetCurrFrame(), &g.op)
	// 	//screen.DrawImage(&sprt.Simg[sprt.imgArrCurrent], &g.op)
	// }

	// screen.DrawImage(&sprt.Simg[sprt.imgArrCurrent], &g.op)
	// screen.DrawImage(, &opts)
}
func (btn *Button) DrawButton_adj(screen *ebiten.Image, adj_X, adj_Y int) {

	// w := btn.Size.X
	// h := btn.Size.Y
	// var opts ebiten.DrawImageOptions
	// opts.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	// opts.GeoM.Rotate(2 * math.Pi * float64(btn.Angle) / float64(180))
	// // g.op.GeoM.Translate(float64(w)/2, float64(h)/2)
	// opts.GeoM.Translate(float64(w)/2, float64(h)/2)
	// opts.GeoM.Translate(float64(btn.Coords.X)+float64(w)/2, float64(btn.Coords.X)+float64(h)/2)
	vector.DrawFilledRect(screen, float32(btn.Coords.X+adj_X), float32(btn.Coords.Y+adj_Y), float32(btn.Size.X), float32(btn.Size.Y), btn.GetColor(), true)
	vector.StrokeRect(screen, float32(btn.Coords.X+adj_X), float32(btn.Coords.Y+adj_Y), float32(btn.Size.X), float32(btn.Size.Y), 2.0, color.Black, true)
	//out := fmt.Sprintf("%s %t\n", btn.Label, btn.IsToggled)
	out := fmt.Sprintf("%s\n", btn.Label)
	// if btn.PointingBool != nil {
	// 	out += fmt.Sprintf("%t\n", *btn.PointingBool)
	// }
	ebitenutil.DebugPrintAt(screen, out, btn.Coords.X, btn.Coords.Y)
	// if sprt.showSimg {
	// 	screen.DrawImage(sprt.animars.GetCurrFrame(), &g.op)
	// 	//screen.DrawImage(&sprt.Simg[sprt.imgArrCurrent], &g.op)
	// }

	// screen.DrawImage(&sprt.Simg[sprt.imgArrCurrent], &g.op)
	// screen.DrawImage(, &opts)
}

type TextPanel struct {
	Position CoordInts
	Size     CoordInts
	Label    string
	Color    color.Color
	Helper   *UI_Helper
}

func (txtPnl *TextPanel) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(txtPnl.Position.X), float32(txtPnl.Position.Y), float32(txtPnl.Size.X), float32(txtPnl.Size.Y), txtPnl.Color, true)
	vector.StrokeRect(screen, float32(txtPnl.Position.X), float32(txtPnl.Position.Y), float32(txtPnl.Size.X), float32(txtPnl.Size.Y), 2.0, color.Black, true)
	scaler := 2.0
	tops := &text.DrawOptions{}

	tops.GeoM.Translate(float64(txtPnl.Position.X+4)*scaler, float64(txtPnl.Position.Y)*scaler)
	tops.GeoM.Scale(1/scaler, 1/scaler)
	tops.ColorScale.ScaleWithColor(color.White)
	tops.LineSpacing = float64(20)
	text.Draw(screen, txtPnl.Label, txtPnl.Helper.Btn_Text_Reg, tops)
	// ebitenutil.DebugPrintAt(screen, txtPnl.Label, txtPnl.Position.X, txtPnl.Position.Y)
}
func (txtPnl *TextPanel) Init(label string, uiHelper *UI_Helper, position, size CoordInts, color color.Color) {
	txtPnl.Position = position
	txtPnl.Size = size
	txtPnl.Label = label
	txtPnl.Color = color
	txtPnl.Helper = uiHelper
}

type NumSelect_Button struct {
	Position CoordInts
	Size     CoordInts
	Btns     [3]Button //has a numSelect button
	CurValue int
	MinValue int
	MaxValue int
	iterator int
	ShowLbl  bool
	Label    TextPanel
}

// func (nsel *NumSelect_Button) Init(name, label string, Pos_X, Pos_Y, BtnWidth, BtnHeight, OffsetX, OffsetY int) {
// 	nsel.Position = CoordInts{X: Pos_X, Y: Pos_Y}
// 	nsel.Size = CoordInts{X: BtnWidth, Y: BtnHeight}

// }
func (nsel *NumSelect_Button) Init(name, label string, helpr *UI_Helper, showlbl bool, Pos_X, Pos_Y, PWidth, PHeight, mintVal, startVal, maxVal, iterate int) {
	nsel.Position = CoordInts{X: Pos_X, Y: Pos_Y}
	nsel.Size = CoordInts{X: 64, Y: PHeight}
	nsel.Btns[0].InitButton("LButton", " -", helpr, 0, Pos_X, Pos_Y, 16, PHeight, 0, 0)
	nsel.Btns[1].InitButton("DButton", "", helpr, 0, Pos_X+16, Pos_Y, 32, PHeight, 0, 0)
	nsel.Btns[2].InitButton("RButton", " +", helpr, 0, Pos_X+48, Pos_Y, 16, PHeight, 0, 0)
	nsel.MinValue = mintVal
	nsel.CurValue = startVal
	nsel.MaxValue = maxVal
	nsel.iterator = iterate
	// nsel.Label = label
	nsel.Label.Init(label, helpr, CoordInts{X: Pos_X, Y: Pos_Y - 16}, CoordInts{X: 64, Y: 16}, color.RGBA{75, 150, 75, 255})
	nsel.ShowLbl = showlbl
}
func (nsel *NumSelect_Button) Draw(screen *ebiten.Image) {
	nsel.Btns[0].DrawButton(screen)
	//DrawArrow(screen, nsel.Btns[0].Coords, nsel.Btns[0].Size, 1.0, color.Black, true)
	nsel.Btns[1].DrawButton(screen)

	nsel.Btns[2].DrawButton(screen)
	if nsel.ShowLbl {
		nsel.Label.Draw(screen)
	}
	// DrawArrow01(screen, nsel.Btns[2].Coords, nsel.Btns[2].Size, 1.0, color.RGBA{255, 100, 100, 255}, false)
}

func (nsel *NumSelect_Button) Update() {
	if nsel.Btns[0].Update3() {
		if (nsel.CurValue - nsel.iterator) >= nsel.MinValue {
			nsel.CurValue -= nsel.iterator
		} else {
			nsel.CurValue = nsel.MinValue
		}
	}
	if nsel.Btns[1].Update3() {

	}
	if nsel.Btns[2].Update3() {
		if (nsel.CurValue + nsel.iterator) <= nsel.MaxValue {
			nsel.CurValue += nsel.iterator
		} else {
			nsel.CurValue = nsel.MaxValue
		}
	}
	nsel.Btns[1].ChangeLabel(fmt.Sprintf(" %03d", nsel.CurValue))
}
func (nsel *NumSelect_Button) GetCurrValue() int {
	return nsel.CurValue
}
func DrawArrow(screen *ebiten.Image, pos CoordInts, cellSize CoordInts, swidth float32, colr color.Color, aa bool) {
	midH := float32(pos.Y) + (float32(cellSize.Y) / 2)
	midW := float32(pos.X) + (float32(cellSize.X) / 2)
	vector.StrokeLine(screen, float32(pos.X), midH, midW, midH, swidth, colr, aa)
	vector.StrokeLine(screen, float32(pos.X), float32(pos.Y), midW, midH, swidth, colr, aa)
	vector.StrokeLine(screen, float32(pos.X), float32(pos.Y+cellSize.Y), midW, midH, swidth, colr, aa)
}
func DrawArrow01(screen *ebiten.Image, pos CoordInts, cellSize CoordInts, swidth float32, colr color.Color, aa bool) {
	// aX := 13
	midH := float32(pos.Y) + (float32(cellSize.Y) / 2)
	midW := float32(pos.X) + (float32(cellSize.X) / 2)
	vector.StrokeLine(screen, midW-2, midH, float32(pos.X+cellSize.X-2), midH, swidth, colr, aa)
	// vector.StrokeLine(screen, float32(pos.X), float32(pos.Y), midW, midH, swidth, colr, aa)
	// vector.StrokeLine(screen, float32(pos.X), float32(pos.Y+cellSize.Y), midW, midH, swidth, colr, aa)
}

/*
---BUTTON PANELS--- Or the basic Button Panel
THE GOAL here is a Button Panel; it will have ButtonPanelStyle as a similar struct that holds a vast majority of it's stats that control for how buttons are place and where
*/

type ButtonPanel struct {
	Position         CoordInts
	PnlDimensions    CoordInts
	PnlBackgroundImg *ebiten.Image
	Buttons          []Button
	BorderMargin     int
	Button_Buffer    int
	Label            string
	Name             string
	// Rows          int
	// Columns       int
}

func (btnPnl *ButtonPanel) Draw(screen *ebiten.Image) {

	// test1X := ((len(imat[0]) * tileW) + (len(imat[0]) * GapX)) + OffsetX
	// test1Y := ((len(imat) * tileH) + (len(imat) * GapY)) + OffsetY
	for _, btn := range btnPnl.Buttons {

		btn.DrawButton_adj(screen, btnPnl.Position.X, btnPnl.Position.Y)
	}
}

func (btnPnl *ButtonPanel) InitBtns(cols int, helper *UI_Helper, size CoordInts) {
	// xx_re, yy_re := btnPnl.BorderMargin, btnPnl.BorderMargin
	xx_i, yy_i := 0, 0
	for i, _ := range btnPnl.Buttons {
		xx_re := btnPnl.BorderMargin + (size.X+btnPnl.Button_Buffer)*xx_i
		yy_re := btnPnl.BorderMargin + (size.Y+btnPnl.Button_Buffer)*yy_i
		strng := fmt.Sprintf("Btn%02d", i)
		btnPnl.Buttons[i].InitButton(strng, strng, helper, 0, xx_re, yy_re, size.X, size.Y, 0, 0)
		if xx_i > cols {
			xx_i = 0
			yy_i++
		} else {
			xx_i++
		}
	}
}
