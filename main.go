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

const (
	buttonHieght0 = 16
	buttonX_1     = 140
	buttonX_2     = 72
	buttonHieght1 = 32
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
	initCalled                                      bool
	gameDebugMsg                                    string
	btn00, btn01, btn02, btn03, btn04, btn05, btn06 mypkgs.Button
	btn07, btn08, btn09, btn10, btn11, btn12, btn13 mypkgs.Button
	btn14, btn15, btn16, btn17, btn18, btn19, btn20 mypkgs.Button //btn21, btn22, btn23, btn24, btn25, btn26, btn27 mypkgs.Button
	btn21                                           mypkgs.Button
	coorAr                                          mypkgs.CoordList
	numPanel00, numPanel01, numPanel02, numPanel03  mypkgs.NumSelect_Button
	numPanel04, numPanel05, numPanel06, numPanel07  mypkgs.NumSelect_Button
	isRunning                                       bool
	IntGrid                                         mypkgs.IntegerGridManager
}

func init() {
	// Settings = mypkgs.GetSettingsFromJSON()
	Settings = mypkgs.GetSettingsFromBakedIn()
	fmt.Printf("DONE INIT\n")
	backgroundImg = ebiten.NewImage(Settings.ScreenResX, Settings.ScreenResY)
	foregroundImg = ebiten.NewImage(Settings.ScreenResX, Settings.ScreenResY)
	// foregroundImg = ebiten.NewImage(320, 240)
	backgroundImg.Fill(backgroundColor)
	foregroundImg.Fill(backgroundColor)
}

func (g *Game) init() error {
	defer func() {
		g.initCalled = true
	}()
	col0 := Settings.ScreenResX - 140
	col1 := Settings.ScreenResX - 72
	g.btn00.InitButton("btn00", "PrintCordArray", 0, col0, 8, 64, 32, 0, 0)
	g.btn01.InitButton("btn01", "SortDescOnX", 0, col1, 8, 64, 32, 0, 0)
	g.btn02.InitButton("btn02", "remove\nduplicates", 0, col0, 44, 64, 32, 0, 0)
	g.btn03.InitButton("btn03", "Clear\nInt Matrix", 0, col1, 44, 64, 32, 0, 0)
	block00 := 86
	g.btn04.InitButton("btn04", "HL Select\nPoints", 0, col0, block00, 64, 32, 0, 0)
	g.btn05.InitButton("btn05", "AUTO:OFF", 2, col1, block00, 64, 32, 0, 0)
	g.btn06.InitButton("btn06", "Process01\nsimpleDecay", 0, col0, block00+36, 64, 32, 0, 0)
	g.btn07.InitButton("btn07", "MazeGen\n2b_noCull", 0, col1, block00+36, 64, 32, 0, 0)
	g.btn08.InitButton("btn08", "MazeGen\n3c", 0, col0, block00+72, 64, 32, 0, 0)
	g.btn09.InitButton("btn09", "Select\nPoints", 2, col1, block00+72, 64, 32, 0, 0)
	g.btn10.InitButton("Btn10", "Clear\nArea", 0, col0, block00+108, 64, 32, 0, 0)
	g.btn11.InitButton("Btn11", "Culling", 0, col1, block00+108, 64, 32, 0, 0)
	block2 := 236
	g.btn12.InitButton("Btn12", "Pathfind\nSet Start", 2, col0, block2, 64, 32, 0, 0)
	g.btn13.InitButton("Btn13", "Pathfind\nSet Stop", 2, col1, block2, 64, 32, 0, 0)
	g.btn14.InitButton("Btn14", "Reset\nStart/Stop", 0, col0, block2+36, 64, 32, 0, 0)
	g.btn15.InitButton("Btn15", "Pathfind\nINIT", 0, col1, block2+36, 64, 32, 0, 0)
	block3 := 314
	g.btn16.InitButton("Btn16", "Pathfind\nBRESENHAM", 0, col0, block3, 64, 32, 0, 0)
	g.btn17.InitButton("Btn17", "Pathfind\nBreadth", 0, col1, block3, 64, 32, 0, 0)
	g.btn18.InitButton("Btn18", "Pathfind\nManhattan", 0, col0, block3+36, 64, 32, 0, 0)
	g.btn19.InitButton("Btn19", "Draw\nCircle", 2, col1, block3+36, 64, 32, 0, 0)
	g.btn20.InitButton("Btn20", "", 0, col0, block3+72, 64, 32, 0, 0)
	g.btn21.InitButton("Btn21", "", 0, col1, block3+72, 64, 32, 0, 0)
	block4 := 444
	g.numPanel00.Init("nums00", "Maze3Param", true, col0, block4, 32, 16, 0, 10, 20, 1)
	g.numPanel01.Init("nums01", "Maze3Param", true, col1, block4, 32, 16, 0, 6, 20, 1)
	g.numPanel02.Init("nums02", "Maze3Param", true, col0, block4+36, 32, 16, 1, 8, 16, 1)
	g.numPanel03.Init("", "circ. Rad", true, col1, block4+36, 32, 16, 0, 0, 20, 1)
	g.numPanel04.Init("nums03", "", true, col0, block4+36+36, 32, 16, 0, 0, 10, 1)
	g.numPanel05.Init("nums05", "FindPath", true, col1, block4+36+36, 32, 16, 0, 0, 3, 1)
	g.numPanel06.Init("nums03", "", true, col0, block4+36+36+36, 32, 16, 0, 0, 10, 1)
	g.numPanel07.Init("nums03", "SCALE", true, col1, block4+36+36+36, 32, 16, 1, 1, 4, 1)
	g.coorAr = append(g.coorAr, mypkgs.CoordInts{X: 2, Y: 2})
	// g.IntGrid.Init(32, 32, 16, 16, 64, 8, 2, 2)
	g.IntGrid.Init(32, 32, 16, 16, 144, 8, 2, 2, 4, 4)

	return nil
}

func (g *Game) PreDrawGUI(screen *ebiten.Image) {
	// mx, _ := ebiten.CursorPosition()
	// if mx > Settings.ScreenResX-200 {

	// }
	g.btn00.DrawButton(screen)
	g.btn01.DrawButton(screen)
	g.btn02.DrawButton(screen)
	g.btn03.DrawButton(screen)

	g.btn04.DrawButton(screen)
	g.btn05.DrawButton(screen)
	g.btn06.DrawButton(screen)
	g.btn07.DrawButton(screen)
	g.btn08.DrawButton(screen)
	g.btn09.DrawButton(screen)
	g.btn10.DrawButton(screen)
	g.btn11.DrawButton(screen)

	g.btn12.DrawButton(screen)
	g.btn13.DrawButton(screen)
	g.btn14.DrawButton(screen)
	g.btn15.DrawButton(screen)
	g.btn16.DrawButton(screen)

	g.btn17.DrawButton(screen)
	g.btn18.DrawButton(screen)
	g.btn19.DrawButton(screen)
	g.btn20.DrawButton(screen)
	g.btn21.DrawButton(screen)

	g.numPanel00.Draw(screen)
	g.numPanel01.Draw(screen)
	g.numPanel02.Draw(screen)
	g.numPanel03.Draw(screen)
	g.numPanel04.Draw(screen)
	g.numPanel05.Draw(screen)
	g.numPanel06.Draw(screen)
	g.numPanel07.Draw(screen)
}

func (g *Game) PreDraw(screen *ebiten.Image) {
	screen.Clear()
	screen.DrawImage(backgroundImg, nil)
	g.IntGrid.Draw(screen) //if made into a goroutine this needs to have some better way of streaming output;
	g.PreDrawGUI(screen)
	//alternatively this perhaps is a great way to
	// imatrix.DrawGridTile(screen, tile_offset_X, tile_offset_Y, tileW, tileH, tile_Margin_W, tile_Margin_H) //DrawGridTile(screen, 8, 8, 16, 16, 2, 2)
	//screen.DrawImage()
}

func (g *Game) Update() error {
	if !g.initCalled {
		g.init()
	}

	mx, my := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		if g.btn19.IsToggled {
			g.IntGrid.DrawACircleOnClick(mx, my, g.numPanel03.GetCurrValue(), 1)
			//g.btn20.IsToggled = false
		}
		// 	//g.btn21.Update(mx, my, true)
	} else {
	}

	g.numPanel00.Update()
	g.numPanel01.Update()
	g.numPanel02.Update()
	g.numPanel03.Update()
	g.numPanel04.Update()
	g.numPanel05.Update()
	g.numPanel06.Update()
	g.numPanel07.Update()
	if g.numPanel07.Btns[1].Update3() {
		fmt.Printf("NPANEL 7  %d\n", g.numPanel07.CurValue)
	}
	if g.btn00.Update3() {
		g.IntGrid.DEMO_COORDS_00(4, 0, 0) //igd.Coords.PrintCordArray()
	}
	if g.btn01.Update3() {
		g.IntGrid.DEMO_COORDS_00(5, 0, 0) //igd.Coords.SortDescOnX()
	}
	if g.btn02.Update3() {
		g.IntGrid.DEMO_COORDS_00(6, 0, 0) //remove duplicates
	}
	if g.btn03.Update3() {
		g.IntGrid.ClearImat()
	}
	if g.btn04.Update3() {
		g.IntGrid.DrawCoordsOnImat()
	}
	if g.btn05.Update3() {
		if g.btn06.BType != 2 {
			g.btn06.BType = 2
			g.btn05.Label = "AUTO:ON"
		}
		if g.btn07.BType != 2 {
			g.btn07.BType = 2
		}
		if g.btn08.BType != 2 {
			g.btn08.BType = 2
		}
	} else {
		if g.btn06.BType == 2 {
			g.btn06.BType = 1
			g.btn05.Label = "AUTO:OFF"
		}
		if g.btn07.BType == 2 {
			g.btn07.BType = 1
		}
		if g.btn08.BType == 2 {
			g.btn08.BType = 1
		}
	}
	if g.btn06.Update3() {
		go g.IntGrid.Process()
	}

	if g.btn07.Update3() {
		g.IntGrid.Process2b(5)
	}
	if g.btn08.Update3() {
		if len(g.IntGrid.Coords) > 0 {
			if !g.IntGrid.AlgorithmRunning {
				g.IntGrid.AlgorithmRunning = true
			}
			// g.IntGrid.Process3c(50, 10, 6, []int{0, 2, 3}) //8,4
			// g.IntGrid.Process3c(50, g.numPanel00.CurValue, g.numPanel01.CurValue, g.numPanel02.CurValue, g.numPanel03.CurValue, g.numPanel04.CurValue, []int{0, 2, 3}) //8,4
			g.IntGrid.Process3c(50, g.numPanel00.CurValue, g.numPanel01.CurValue, []int{0, 2, 3}) //8,4 //, g.numPanel02.CurValue

		}
	}
	if g.btn09.Update3() {
		g.IntGrid.SelectPoints = true
	} else {
		g.IntGrid.SelectPoints = false
	}
	if g.btn10.Update3() {
		g.IntGrid.Imat.ClearAnArea(3, 3, 29, 29, 1)
	}
	if g.btn11.Update3() {
		g.IntGrid.CullCoords(8, true, []int{0, 2})
	}
	if g.btn12.Update3() && !g.IntGrid.PFinder.IsStartInit {
		g.IntGrid.PFinderStartSelect = true
	} else {
		g.IntGrid.PFinderStartSelect = false
	}
	if g.btn13.Update3() && !g.IntGrid.PFinder.IsEndInit {
		g.IntGrid.PFinderEndSelect = true
	} else {
		g.IntGrid.PFinderEndSelect = false
	}
	if g.btn14.Update3() { //RESET
		g.IntGrid.RESETPathfinder()
	}
	if g.btn15.Update3() {
		g.IntGrid.PathfindingProcess()
	}
	if g.btn16.Update3() {
		g.IntGrid.PFindr_DrawBresenHamLine([]int{0, 2, 3, 4})
		// go g.IntGrid.MoveCursorAround(mypkgs.CoordInts{X: 2, Y: 2}, []int{0, 2, 3, 4})
	}
	if g.btn17.Update3() {
		// g.IntGrid.PFindr_DrawSlope()
		// g.IntGrid.PFindr_DrawManhattan()
		g.IntGrid.FindPath(g.numPanel05.CurValue)
		//mypkgs.FindPath(g.IntGrid.Imat,g.IntGrid.PFinder.StartPos,g.I)
		//g.IntGrid.PFinder.HasFalsePos = !g.IntGrid.PFinder.HasFalsePos
	}
	if g.btn18.Update3() {
		g.IntGrid.PFindr_DrawManhattan2([]int{0, 2, 3, 4})
	}
	if g.btn19.Update3() {

	}
	if g.btn20.Update3() {
		g.IntGrid.PFinder.Cursor.ShowNeighbors = !g.IntGrid.PFinder.Cursor.ShowNeighbors
	}
	//inpututil.IsKeyJustPressed(ebiten.KeyW)
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {

		//g.numPanel00.CurValue
		if g.IntGrid.MoveCursorFreely(0, 1, []int{0, 2, 3, 4}) {
			//g.IntGrid.Position.Y += 18
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		// g.IntGrid.Position.X -= 1

		if g.IntGrid.MoveCursorFreely(3, 1, []int{0, 2, 3, 4}) {
			//g.IntGrid.Position.X += 18
		}

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		// g.IntGrid.Position.Y += 1
		if g.IntGrid.MoveCursorFreely(2, 1, []int{0, 2, 3, 4}) {
			//g.IntGrid.Position.Y -= 18
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.IntGrid.MoveCursorFreely(1, 1, []int{0, 2, 3, 4}) {
			//g.IntGrid.Position.X -= 18
		}

		// g.IntGrid.Position.X += 1
	}

	go g.IntGrid.UpdateOnMouseEvent2()
	g.PreDraw(foregroundImg)
	g.gameDebugMsg = fmt.Sprintf("FPS:%8.3f TPS:%8.3f\n", ebiten.ActualFPS(), ebiten.ActualTPS())
	g.gameDebugMsg += fmt.Sprintf("%s\n", Settings.ToString())
	//g.gameDebugMsg += fmt.Sprintf("BTN0: %2d btn01:%2d btn02:%2d\n", g.btn00.State, g.btn01.State, g.btn02.State)
	g.gameDebugMsg += "------------------------\n"
	g.gameDebugMsg += g.IntGrid.PFinder.ToString()
	//g.gameDebugMsg += fmt.Sprintf("\tIS INIT?:\n\t\tSTART:%t\n\t\tSTOP:%t\n\t\tFULL:%t\n", g.IntGrid.PFinder.IsEndInit, g.IntGrid.PFinder.IsEndInit, g.IntGrid.PFinder.IsFullyInitialized)
	// g.gameDebugMsg += fmt.Sprintf("\t")
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
