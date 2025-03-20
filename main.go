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
	Settings mypkgs.GameSettings
	// imatrix         mypkgs.IntMatrix
	backgroundColor color.RGBA = color.RGBA{150, 100, 250, 255}
	clearColor      color.RGBA = color.RGBA{0, 0, 0, 0}
	backgroundImg   *ebiten.Image
	foregroundImg   *ebiten.Image
)

type Game struct {
	initCalled          bool
	gameDebugMsg        string
	btn0                mypkgs.Button
	btn1, btn2, btn3    mypkgs.Button
	btn4, btn5, btn6    mypkgs.Button
	btn7, btn8, btn9    mypkgs.Button
	btn10, btn11, btn12 mypkgs.Button
	coorAr              mypkgs.CoordArray
	isRunning           bool
	IntGrid             mypkgs.IntegerGridManager
}

func init() {
	Settings = mypkgs.GetSettingsFromJSON()
	// Settings = mypkgs.GetSettingsFromBakedIn()
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
	g.btn0.InitButton("Btn0", "PrintCordArray", 0, Settings.ScreenResX-140, 8, 64, 16, 0, 0)
	g.btn1.InitButton("Btn1", "SortDescOnX", 0, Settings.ScreenResX-72, 8, 64, 16, 0, 0)
	g.btn2.InitButton("Btn2", "remove duplicates", 0, Settings.ScreenResX-140, 28, 64, 16, 0, 0)
	g.btn3.InitButton("Btn3", "clearImat", 0, Settings.ScreenResX-72, 28, 64, 16, 0, 0)
	g.btn4.InitButton("Btn4", "drawArPoints", 0, Settings.ScreenResX-140, 48, 64, 16, 0, 0)
	g.btn5.InitButton("Btn5", "AUTO:OFF", 2, Settings.ScreenResX-72, 48, 64, 16, 0, 0)
	g.btn6.InitButton("Btn6", "prc1_01", 0, Settings.ScreenResX-140, 68, 64, 16, 0, 0)
	g.btn7.InitButton("Btn7", "prc2b_5", 0, Settings.ScreenResX-72, 68, 64, 16, 0, 0)
	g.btn8.InitButton("Btn8", "prc3", 0, Settings.ScreenResX-140, 88, 64, 16, 0, 0)
	g.btn9.InitButton("Btn9", "9", 0, Settings.ScreenResX-72, 88, 64, 16, 0, 0)
	g.btn10.InitButton("Btn10", "Btn10", 0, Settings.ScreenResX-140, 108, 64, 16, 0, 0)
	g.btn11.InitButton("Btn11", "Btn11", 0, Settings.ScreenResX-72, 108, 64, 16, 0, 0)
	g.coorAr = append(g.coorAr, mypkgs.CoordInts{X: 2, Y: 2})
	g.IntGrid.Init(32, 32, 16, 16, 8, 8, 4, 4)
	return nil
}
func (g *Game) PreDraw(screen *ebiten.Image) {
	screen.Clear()
	screen.DrawImage(backgroundImg, nil)
	g.IntGrid.Draw(screen)
	// imatrix.DrawGridTile(screen, tile_offset_X, tile_offset_Y, tileW, tileH, tile_Margin_W, tile_Margin_H) //DrawGridTile(screen, 8, 8, 16, 16, 2, 2)
	g.btn0.DrawButton(screen)
	g.btn1.DrawButton(screen)
	g.btn2.DrawButton(screen)
	g.btn3.DrawButton(screen)
	g.btn4.DrawButton(screen)
	g.btn5.DrawButton(screen)
	g.btn6.DrawButton(screen)
	g.btn7.DrawButton(screen)
	g.btn8.DrawButton(screen)
	g.btn9.DrawButton(screen)
	g.btn10.DrawButton(screen)
	g.btn11.DrawButton(screen)
	//screen.DrawImage()
}

func (g *Game) Update() error {
	if !g.initCalled {
		g.init()
	}

	mx, my := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		g.IntGrid.UpdateOnMouseEvent(mx, my)
		g.btn0.Update(mx, my, true)
		g.btn1.Update(mx, my, true)
		g.btn2.Update(mx, my, true)
		g.btn3.Update(mx, my, true)
		g.btn4.Update(mx, my, true)
		g.btn5.Update(mx, my, true)
		g.btn6.Update(mx, my, true)
		g.btn7.Update(mx, my, true)
		g.btn8.Update(mx, my, true)
		g.btn9.Update(mx, my, true)
		g.btn10.Update(mx, my, true)
		g.btn11.Update(mx, my, true)
	} else {
		g.btn0.Update(mx, my, false)
		g.btn1.Update(mx, my, false)
		g.btn2.Update(mx, my, false)
		g.btn3.Update(mx, my, false)
		g.btn4.Update(mx, my, false)
		g.btn5.Update(mx, my, false)
		g.btn6.Update(mx, my, false)
		g.btn7.Update(mx, my, false)
		g.btn8.Update(mx, my, false)
		g.btn9.Update(mx, my, false)
		g.btn10.Update(mx, my, false)
		g.btn11.Update(mx, my, false)
	}

	if g.btn0.UpdateTwo() {
		g.IntGrid.DEMO_COORDS_00(4, 0, 0) //igd.Coords.PrintCordArray()
	}
	if g.btn1.UpdateTwo() {
		g.IntGrid.DEMO_COORDS_00(5, 0, 0) //igd.Coords.SortDescOnX()
	}
	if g.btn2.UpdateTwo() {
		g.IntGrid.DEMO_COORDS_00(6, 0, 0) //remove duplicates
	}
	if g.btn3.UpdateTwo() {
		g.IntGrid.ClearImat()
	}
	if g.btn4.UpdateTwo() {
		g.IntGrid.DrawCoordsOnImat()
	}
	if g.btn5.UpdateTwo() {
		if g.btn6.BType != 2 {
			g.btn6.BType = 2
			g.btn5.Label = "AUTO:ON"
		}
		if g.btn7.BType != 2 {
			g.btn7.BType = 2
		}
		if g.btn8.BType != 2 {
			g.btn8.BType = 2
		}
	} else {
		if g.btn6.BType == 2 {
			g.btn6.BType = 1
			g.btn5.Label = "AUTO:OFF"
		}
		if g.btn7.BType == 2 {
			g.btn7.BType = 1
		}
		if g.btn8.BType == 2 {
			g.btn8.BType = 1
		}
	}
	if g.btn6.UpdateTwo() {
		g.IntGrid.Process()
	}
	if g.btn7.UpdateTwo() {
		g.IntGrid.Process2b(5)
	}
	if g.btn8.UpdateTwo() {
		if len(g.IntGrid.Coords) > 0 {
			g.IntGrid.Process3(8, 4, []int{0, 2})
		}
	}
	if g.btn9.UpdateTwo() {
		g.IntGrid.CullCoords(2, true, []int{0, 2})
	}
	if g.btn10.UpdateTwo() {
		g.IntGrid.CullCoords(4, true, []int{0, 2})
	}
	if g.btn11.UpdateTwo() {
		g.IntGrid.CullCoords(8, true, []int{0, 2})
	}
	g.PreDraw(foregroundImg)
	g.gameDebugMsg = fmt.Sprintf("FPS:%8.3f TPS:%8.3f\n", ebiten.ActualFPS(), ebiten.ActualTPS())
	g.gameDebugMsg += fmt.Sprintf("%s\n", Settings.ToString())
	g.gameDebugMsg += fmt.Sprintf("BTN0: %2d BTN1:%2d BTN2:%2d\n", g.btn0.State, g.btn1.State, g.btn2.State)
	g.gameDebugMsg += "------------------------\n"
	g.gameDebugMsg += fmt.Sprintf("%s\n", g.IntGrid.ToString())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.DrawImage(backgroundImg, nil)
	screen.DrawImage(foregroundImg, nil)
	ebitenutil.DebugPrint(screen, g.gameDebugMsg)

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
