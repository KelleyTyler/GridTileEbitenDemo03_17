package mypkgs

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type General_UI_Interface interface {
	Init(helper *UI_Helper, settings *GameSettings, coords []CoordInts)
	Update()
	UpdateAdj(parentPos CoordInts)
	Draw(screen *ebiten.Image)
	ToString() string
	GetType() string
}

type UI_Container_Interface interface {
}

type UI_Container_type1 struct {
}

//	type UI_General_Values struct {
//		SettingsPtr *GameSettings
//		UIHelpPtr   *UI_Helper
//		Position    CoordInts
//		Dimensions  CoordInts
//		IsActive    bool
//		IsVisible   bool
//		IsHovered   bool
//	}
type Label_Mk2 struct {
	// UIValues UI_General_Values
	SettingsPtr *GameSettings
	UIHelpPtr   *UI_Helper
	Position    *CoordInts
	Dimensions  *CoordInts
	DisplayImg  *ebiten.Image
	IsActive    bool
	IsVisible   bool
	IsHovered   bool
	Label       string
}

/*
	Label_Mk2.Init

-coords[0]=Position,
-coords[1]=Dimensions,
*/
func (lab *Label_Mk2) Init(helper *UI_Helper, settings *GameSettings, coords []CoordInts, strngs []string, isBool []bool) { //textOPs *text.DrawOptions,
	lab.Position = &CoordInts{X: coords[0].X, Y: coords[0].Y}
	lab.Dimensions = &CoordInts{X: coords[1].X, Y: coords[1].Y}
	lab.Label = strngs[0]
	lab.IsActive = true
	lab.IsVisible = true
	lab.SettingsPtr = settings
	lab.UIHelpPtr = helper
	lab.MakeImage(lab.Label)
}
func (lab *Label_Mk2) Draw(screen *ebiten.Image) {

}
func (lab *Label_Mk2) MakeImage(labeler string) {
	lab.DisplayImg = ebiten.NewImage(lab.Dimensions.X, lab.Dimensions.Y)
	lab.DisplayImg.Fill(lab.UIHelpPtr.Button_Colors[0])
	textOps := &text.DrawOptions{}
	textOps.GeoM.Reset()
	textOps.LayoutOptions.PrimaryAlign = text.AlignCenter
	textOps.LayoutOptions.SecondaryAlign = text.AlignCenter

}

type Button_Mk2 struct { //GenerUI_Thing
	SettingsPtr *GameSettings
	UIHelpPtr   *UI_Helper
	Position    CoordInts
	Dimensions  CoordInts
	IsActive    bool
	IsVisible   bool
	// IsHovered   bool
	Name  string
	Label string

	General   *func()
	Hovered   *func()
	Click     *func()
	Activated *func()
}

type MenuBar struct {
	MainBackgroundImg *ebiten.Image
	SettingsPtr       *GameSettings
	UIHelpPtr         *UI_Helper
	Position          CoordInts
	Dimensions        CoordInts
}

func (mBar *MenuBar) ToString() string {
	return fmt.Sprintf("MENUBAR! has SettingsPtr %5t UI Helper %5t", mBar.SettingsPtr != nil, mBar.UIHelpPtr != nil)
}
func (mBar *MenuBar) Init(helper *UI_Helper, settings *GameSettings, coords []CoordInts) {
	mBar.Dimensions = CoordInts{settings.ScreenResX, coords[0].Y}
	mBar.Position = CoordInts{X: 0, Y: 0}
	mBar.SettingsPtr = settings
	// mBar.MainBackgroundImg

}

type DropdownMenu struct {
	SurfaceButton   Button
	DropdownButtons []Button_Mk2
	Label           string
	ButtonMargin    CoordInts
	Panel_OpenSize  CoordInts
	SettingPtr      *GameSettings
	UIHelpPtr       *UI_Helper
	DropDownState   uint8
}
