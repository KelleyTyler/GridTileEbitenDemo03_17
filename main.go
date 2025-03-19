package main

import (
	// 	"fmt"
	"fmt"
	"image/color"
	"log"

	mypkgs "github.com/KelleyTyler/GridTileEbitenDemo03_17/myPkgs"
	//"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	Settings        mypkgs.GameSettings
	imatrix         mypkgs.IntMatrix
	backgroundColor color.RGBA = color.RGBA{150, 100, 250, 255}
	clearColor      color.RGBA = color.RGBA{0, 0, 0, 0}
	backgroundImg   *ebiten.Image
	foregroundImg   *ebiten.Image
)

type Game struct {
	initCalled bool
	//settings mypkgs.GameSettings
	gameDebugMsg string
	btn0         mypkgs.Button
	btn1         mypkgs.Button
	btn2         mypkgs.Button
	coorAr       mypkgs.CoordArray
	isRunning    bool
}

func init() {
	Settings = mypkgs.GetSettingsFromJSON()
	// imatrix = imatrix.MakeIntMatrix(64, 64)
	// imatrix.InitBlankMatrix(64, 64, 0)
	imatrix = imatrix.MakeIntMatrix(32, 32)
	imatrix.InitBlankMatrix(32, 32, 0)
	//imatrix.MakeIntMatrixRowBlank(0, 5, 1)
	// imatrix[0][5] = 0
	// imatrix[1][5] = 0
	// imatrix[2][5] = 0
	// imatrix[3][5] = 0
	// imatrix[3][6] = 0
	// imatrix[3][7] = 0
	// imatrix[3][8] = 0
	//imatrix.PrintMatrix()
	fmt.Printf("DONE INIT\n")
	backgroundImg = ebiten.NewImage(Settings.ScreenResX, Settings.ScreenResY)
	foregroundImg = ebiten.NewImage(Settings.ScreenResX, Settings.ScreenResY)
	// foregroundImg = ebiten.NewImage(320, 240)
	backgroundImg.Fill(backgroundColor)
	foregroundImg.Fill(clearColor)
}

func (g *Game) init() error {
	defer func() {
		g.initCalled = true
	}()
	g.btn0 = mypkgs.Button{
		Coords: mypkgs.CoordInts{X: Settings.ScreenResX - 72, Y: 8},
		Offset: mypkgs.CoordInts{X: 0, Y: 0},
		Size:   mypkgs.CoordInts{X: 64, Y: 16},
		Color: []color.Color{color.RGBA{75, 150, 75, 255}, color.RGBA{120, 220, 75, 255}, color.RGBA{140, 230, 90, 255},
			color.RGBA{150, 75, 75, 255}, color.RGBA{220, 120, 75, 255}, color.RGBA{240, 140, 90, 255}},
		Angle:     0,
		Label:     "->",
		Name:      "Btn0",
		State:     0,
		BType:     0,
		IsEnabled: true,
	}
	g.btn1 = mypkgs.Button{
		Coords: mypkgs.CoordInts{X: Settings.ScreenResX - 72, Y: 28},
		Offset: mypkgs.CoordInts{X: 0, Y: 0},
		Size:   mypkgs.CoordInts{X: 64, Y: 16},
		Color: []color.Color{color.RGBA{75, 150, 75, 255}, color.RGBA{120, 220, 75, 255}, color.RGBA{140, 230, 90, 255},
			color.RGBA{150, 75, 75, 255}, color.RGBA{220, 120, 75, 255}, color.RGBA{240, 140, 90, 255}},
		Angle:     0,
		Label:     "->",
		Name:      "Btn0",
		State:     0,
		BType:     2,
		IsEnabled: true,
	}
	g.btn2 = mypkgs.Button{
		Coords: mypkgs.CoordInts{X: Settings.ScreenResX - 72, Y: 48},
		Offset: mypkgs.CoordInts{X: 0, Y: 0},
		Size:   mypkgs.CoordInts{X: 64, Y: 16},
		Color: []color.Color{color.RGBA{75, 150, 75, 255}, color.RGBA{120, 220, 75, 255}, color.RGBA{140, 230, 90, 255},
			color.RGBA{150, 75, 75, 255}, color.RGBA{220, 120, 75, 255}, color.RGBA{240, 140, 90, 255}},
		Angle:     0,
		Label:     "->",
		Name:      "Btn0",
		State:     0,
		BType:     0,
		IsEnabled: true,
	}
	g.coorAr = append(g.coorAr, mypkgs.CoordInts{2, 2})
	return nil
}
func (g *Game) PreDraw(screen *ebiten.Image) {
	screen.Clear()
	screen.DrawImage(backgroundImg, nil)

	imatrix.DrawGridTile(screen, 8, 8, 16, 16, 2, 2)
	g.btn0.DrawButton(screen)
	g.btn1.DrawButton(screen)
	g.btn2.DrawButton(screen)
	//screen.DrawImage()
}

func (g *Game) Update() error {
	if !g.initCalled {
		g.init()
	}

	mx, my := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		imatrix.ChangeValOnMouseEvent(mx, my, 8, 8, 16, 16, 2, 2, 0, 2)
		g.btn0.Update(mx, my, true)
		g.btn1.Update(mx, my, true)
		g.btn2.Update(mx, my, true)
	} else {
		g.btn0.Update(mx, my, false)
		g.btn1.Update(mx, my, false)
		g.btn2.Update(mx, my, false)
	}

	if g.btn0.State == 3 {
		//g.coorAr = imatrix.Process(g.coorAr, 4, 20)
		g.coorAr = imatrix.Process(g.coorAr, 1, 20)
	}
	if g.btn1.IsToggled {
		g.coorAr = imatrix.Process(g.coorAr, 1, 20)
	}
	if g.btn2.State == 3 {

		g.coorAr = imatrix.SETEVERYTHINGTOGREEN(g.coorAr)
		imatrix.Hilight_FromCoord(g.coorAr)
	}
	g.PreDraw(foregroundImg)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.DrawImage(backgroundImg, nil)
	screen.DrawImage(foregroundImg, nil)
	ebitenutil.DebugPrint(screen, "Hello, World!\n"+Settings.ToString()+"\n"+g.coorAr.ToString())

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Settings.ScreenResX, Settings.ScreenResY
}

func main() {

	ebiten.SetWindowSize(Settings.WindowSizeX, Settings.WindowSizeY)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
