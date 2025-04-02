package main

import (
	// 	"fmt"
	"fmt"
	"image/color"
	"log"
	"math"

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
	backgroundColor color.RGBA = color.RGBA{20, 20, 20, 255} //color.RGBA{50, 50, 50, 255}
	clearColor      color.RGBA = color.RGBA{0, 0, 0, 0}
	backgroundImg   *ebiten.Image
	foregroundImg   *ebiten.Image
)

type Game struct {
	initCalled                                             bool
	isQuit                                                 bool
	gameDebugMsg                                           string
	btn00, btn01, btn02, btn03, btn04, btn05, btn06        mypkgs.Button
	btn07, btn08, btn09, btn10, btn11, btn12, btn13        mypkgs.Button
	btn14, btn15, btn16, btn17, btn18, btn19, btn20        mypkgs.Button //btn21, btn22, btn23, btn24, btn25, btn26, btn27 mypkgs.Button
	btn21                                                  mypkgs.Button
	coorAr                                                 mypkgs.CoordList
	numPanel00, numPanel01, numPanel02, numPanel03         mypkgs.NumSelect_Button
	audioTestNumPanel, numPanel05, TileMargin, ScaleNumPad mypkgs.NumSelect_Button
	isRunning                                              bool
	IntGrid                                                mypkgs.IntegerGridManager

	MouseDragStartingPoint mypkgs.CoordInts
	MouseIsDragging        bool

	SoundThing mypkgs.AudioThing
	UIHelp     mypkgs.UI_Helper
}

func init() {
	Settings = mypkgs.GetSettingsFromJSON()
	// Settings = mypkgs.GetSettingsFromBakedIn()
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
	g.SoundThing.Init01(&Settings, 3200, 220, 0, 110) //4800, 220, -15, 20
	g.UIHelp.Init_Default(&g.SoundThing)
	col0 := Settings.ScreenResX - 140
	col1 := Settings.ScreenResX - 72
	g.btn00.InitButton("btn00", "PrintCordArray", &g.UIHelp, 0, col0, 8, 64, 32, 0, 0)
	g.btn01.InitButton("btn01", "SortDescOnX", &g.UIHelp, 0, col1, 8, 64, 32, 0, 0)
	g.btn02.InitButton("btn02", "remove\nduplicates", &g.UIHelp, 0, col0, 44, 64, 32, 0, 0)
	g.btn03.InitButton("btn03", "Clear\nInt Matrix", &g.UIHelp, 0, col1, 44, 64, 32, 0, 0)
	block00 := 86
	g.btn04.InitButton("btn04", "HL Select\nPoints", &g.UIHelp, 2, col0, block00, 64, 32, 0, 0)
	g.btn05.InitButton("btn05", "AUTO:OFF", &g.UIHelp, 2, col1, block00, 64, 32, 0, 0)
	g.btn06.InitButton("btn06", "Simple\nDecay", &g.UIHelp, 0, col0, block00+36, 64, 32, 0, 0)
	g.btn07.InitButton("btn07", "Primlike\nMaze Gen", &g.UIHelp, 0, col1, block00+36, 64, 32, 0, 0)
	g.btn08.InitButton("btn08", "MazeGen\n3c", &g.UIHelp, 0, col0, block00+72, 64, 32, 0, 0)
	g.btn09.InitButton("btn09", "Select\nPoints", &g.UIHelp, 2, col1, block00+72, 64, 32, 0, 0)
	g.btn10.InitButton("Btn10", "Clear\nArea", &g.UIHelp, 0, col0, block00+108, 64, 32, 0, 0)
	g.btn11.InitButton("Btn11", "Culling", &g.UIHelp, 0, col1, block00+108, 64, 32, 0, 0)
	block2 := 236
	g.btn12.InitButton("Btn12", "Pathfind\nSet Start", &g.UIHelp, 2, col0, block2, 64, 32, 0, 0)
	g.btn13.InitButton("Btn13", "Pathfind\nSet Stop", &g.UIHelp, 2, col1, block2, 64, 32, 0, 0)
	g.btn14.InitButton("Btn14", "Reset\nStart/Stop", &g.UIHelp, 0, col0, block2+36, 64, 32, 0, 0)
	g.btn15.InitButton("Btn15", "Pathfind\nINIT", &g.UIHelp, 0, col1, block2+36, 64, 32, 0, 0)
	block3 := 314
	g.btn16.InitButton("Btn16", "Pathfind\nBRESENHAM", &g.UIHelp, 0, col0, block3, 64, 32, 0, 0)
	g.btn17.InitButton("Btn17", "Pathfind\nBreadth", &g.UIHelp, 0, col1, block3, 64, 32, 0, 0)
	g.btn18.InitButton("Btn18", "Pathfind\nManhattan", &g.UIHelp, 0, col0, block3+36, 64, 32, 0, 0)
	g.btn19.InitButton("Btn19", "Draw\nCircle", &g.UIHelp, 2, col1, block3+36, 64, 32, 0, 0)
	g.btn20.InitButton("Btn20", "ShowCursr\nneighbors", &g.UIHelp, 0, col0, block3+72, 64, 32, 0, 0)
	g.btn21.InitButton("Btn21", "AddCirc\ntoMazeGen", &g.UIHelp, 2, col1, block3+72, 64, 32, 0, 0)
	block4 := 444
	g.numPanel00.Init("nums00", "Maze3Param", &g.UIHelp, true, col0, block4, 32, 16, 0, 10, 20, 1)
	g.numPanel01.Init("nums01", "Maze3Param", &g.UIHelp, true, col1, block4, 32, 16, 0, 6, 20, 1)
	g.numPanel02.Init("nums02", "Maze3Param", &g.UIHelp, true, col0, block4+36, 32, 16, 1, 8, 16, 1)
	g.numPanel03.Init("circRadPanel", "circ. Rad", &g.UIHelp, true, col1, block4+36, 32, 16, 0, 0, 20, 1)
	// g.audioTestNumPanel.Init("AudioTest", "AudioTest", true, col0, block4+36+36, 32, 16, 0, 0, 3, 1)
	g.numPanel05.Init("nums05", "FindPath", &g.UIHelp, true, col1, block4+36+36, 32, 16, 0, 0, 3, 1)

	g.audioTestNumPanel.Init("AudioTest", "AudioTest", &g.UIHelp, true, col0, block4+36+36, 32, 16, 0, 0, 5, 1)
	//=----------
	g.TileMargin.Init("TileMargin_Selector", "TileMargin", &g.UIHelp, true, col0, block4+36+36+36, 32, 16, 0, 4, 16, 1)
	g.ScaleNumPad.Init("Scale_Selector", "SCALE", &g.UIHelp, true, col1, block4+36+36+36, 32, 16, 1, 4, 32, 1)
	g.coorAr = append(g.coorAr, mypkgs.CoordInts{X: 2, Y: 2})
	// g.IntGrid.Init(32, 32, 16, 16, 64, 8, 2, 2)
	g.IntGrid.Init(&g.UIHelp, 64, 64, 8, 8, 168, 8, 0, 0, 4, 4)
	// g.IntGrid.Init(96, 96, 8, 8, 64, 8, 0, 0, 4, 4)
	g.MouseDragStartingPoint = mypkgs.CoordInts{X: 0, Y: 0}
	g.MouseIsDragging = false

	g.SoundThing.AddToAudioThing(10, 110) //01
	g.SoundThing.AddToAudioThing(15, 110) //02
	g.SoundThing.AddToAudioThing(20, 110) //03
	g.SoundThing.AddToAudioThing(25, 110) //04
	g.SoundThing.AddToAudioThing(25, 110) //05
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
	g.audioTestNumPanel.Draw(screen)
	g.numPanel05.Draw(screen)
	g.TileMargin.Draw(screen)
	g.ScaleNumPad.Draw(screen)
}

func (g *Game) PreDraw(screen *ebiten.Image) {
	screen.Clear()
	screen.DrawImage(backgroundImg, nil)
	// g.IntGrid.RedrawBoard()
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
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) {
		g.MouseIsDragging = true
		mx, my := ebiten.CursorPosition()
		g.MouseDragStartingPoint = mypkgs.CoordInts{X: mx, Y: my}
		//fmt.Printf("DRAGGING YOUR MOUSE\n")
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton2) {
		if g.MouseIsDragging {
			g.MouseIsDragging = false
			// x0, y0 := ebiten.CursorPosition()
			// x1 := g.MouseDragStartingPoint.X
			// y1 := g.MouseDragStartingPoint.Y

			// //fmt.Printf("HEY YOU DRAGGED YOUR MOUSE:\n%8s %6.2d,%6.2d\n%8s %6.2f %6.2f\n%8s %6.2d,%6.2d\n", "RAW", x1-x0, y1-y0, "Red.(f)", float32(x1-x0)/4.00, float32(y1-y0)/4.0, "Red.(i)", (x1-x0)/4.00, (y1-y0)/4.0)
			// x2, y2 := (x1-x0)/4.00, (y1-y0)/4.0
			// g.IntGrid.BoardPosition.Y -= y2
			// g.IntGrid.BoardPosition.X -= x2
			// g.MouseDragStartingPoint = mypkgs.CoordInts{X: 0, Y: 0}
		}

	}

	if g.MouseIsDragging {
		x0, y0 := ebiten.CursorPosition()
		x1 := g.MouseDragStartingPoint.X
		y1 := g.MouseDragStartingPoint.Y
		t0 := g.MouseDragStartingPoint.X == x0 && g.MouseDragStartingPoint.Y == y0
		t1 := int(math.Abs(float64(x0-x1))) < 2 && int(math.Abs(float64(y0-y1))) < 2
		if t0 || t1 {
			g.MouseDragStartingPoint = mypkgs.CoordInts{X: x0, Y: y0}
			//fmt.Printf("DRAGGING YOUR MOUSE\n")
		} else {
			x2, y2 := (x1-x0)/4.00, (y1-y0)/4.0
			g.IntGrid.BoardPosition.Y -= y2
			g.IntGrid.BoardPosition.X -= x2
			g.MouseDragStartingPoint = mypkgs.CoordInts{X: x0, Y: y0}
		}
	}
	g.TileMargin.Update()
	g.ScaleNumPad.Update()
	if g.ScaleNumPad.Btns[1].Update3() {
		fmt.Printf("SCALE  %3d + %3d\n", 4*g.ScaleNumPad.CurValue, g.TileMargin.CurValue)
		g.IntGrid.Rescale(4*g.ScaleNumPad.CurValue, 4*g.ScaleNumPad.CurValue, g.TileMargin.CurValue, g.TileMargin.CurValue)
	}
	g.numPanel00.Update()
	g.numPanel01.Update()
	g.numPanel02.Update()
	g.numPanel03.Update()
	g.audioTestNumPanel.Update()
	if g.audioTestNumPanel.Btns[1].Update3() {
		g.UIHelp.PlaySound(g.audioTestNumPanel.CurValue)
		// g.SoundThing.PlayThing(g.audioTestNumPanel.CurValue)
		// fmt.Printf("%s\n\n", g.UIHelp.ToString())
	}
	g.numPanel05.Update()

	if g.btn00.Update3() {
		// g.IntGrid.DEMO_COORDS_00(4, 0, 0) //igd.Coords.PrintCordArray()
		// if g.IntGrid.Tile_Size.X == 16 {
		// 	g.IntGrid.Rescale(32, 32, 4, 4)
		// } else {
		// 	g.IntGrid.Rescale(16, 16, 2, 2)
		// }
		g.IntGrid.PFinder.Cursor.ShowCircle = !g.IntGrid.PFinder.Cursor.ShowCircle
	}
	if g.btn01.Update3() {
		g.IntGrid.DEMO_COORDS_00(5, 0, 0) //igd.Coords.SortDescOnX()
	}
	if g.btn02.Update3() {
		g.IntGrid.DEMO_COORDS_00(6, 0, 0) //remove duplicates
	}
	if g.btn03.Update3() {
		g.IntGrid.ClearImat()
		g.IntGrid.MazeM.ClearCords0()
	}
	if g.btn04.Update3() {
		g.IntGrid.MazeM.Cords0_IsVisible = true
	} else {
		g.IntGrid.MazeM.Cords0_IsVisible = false
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
		// go g.IntGrid.Process()
		go g.IntGrid.MazeM.BasicDecayProcess([]int{1, 2, 3, 4, 5}, [4]int{5, 6, 6, 5})
	}

	if g.btn07.Update3() {
		// g.IntGrid.Process2b(5)
		// go g.IntGrid.MazeM.MoreAdvancedDecay([]int{1, 2, 3, 4, 5}, [4]int{1, 2, 2, 1})
		// g.IntGrid.MazeM.PrimLike_Maze_Algorithm00_Looper([]int{1, 2, 3, 4, 5}, []int{-1, 1, 2, 4}, [4]int{1, 2, 2, 1}, true)
		g.IntGrid.MazeM.PrimeLike_Wrapper(5, []int{1, 2, 3, 4, 5}, []int{-1, 1, 2, 4}, [4]int{1, 2, 2, 1}, true)

	}
	if g.btn08.Update3() {

		// fmt.Printf("%s", g.SoundThing.ToString())
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

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// backgroundImg.Fill(color.RGBA{150, 150, 150, 255})
		// g.IntGrid.Img.Fill(color.RGBA{150, 150, 150, 255})
		g.IntGrid.ResetCoordPosition()
	}
	if inpututil.IsKeyJustPressed(ebiten.Key8) {
		backgroundImg.Fill(backgroundColor)
	}
	//inpututil.IsKeyJustPressed(ebiten.KeyW)
	if ebiten.IsKeyPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyShiftLeft) { //inpututil.IsKeyJustPressed(ebiten.KeyArrowUp)
		g.IntGrid.BoardPosition.Y += 1
		// g.IntGrid.RedrawBoard()
		// g.IntGrid.Img.Fill(color.RGBA{150, 150, 150, 255})
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) && ebiten.IsKeyPressed(ebiten.KeyShiftLeft) { //inpututil.IsKeyJustPressed(ebiten.KeyArrowDown)
		g.IntGrid.BoardPosition.Y -= 1
		// g.IntGrid.RedrawBoard()

		//g.IntGrid.Img.Fill(color.RGBA{150, 150, 150, 255})
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && ebiten.IsKeyPressed(ebiten.KeyShiftLeft) { //inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft)
		g.IntGrid.BoardPosition.X += 1
		// g.IntGrid.RedrawBoard()

		//g.IntGrid.Img.Fill(color.RGBA{150, 150, 150, 255})
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyShiftLeft) { //inpututil.IsKeyJustPressed(ebiten.KeyArrowRight)
		g.IntGrid.BoardPosition.X -= 1
		// g.IntGrid.RedrawBoard()

		//g.IntGrid.Img.Fill(color.RGBA{150, 150, 150, 255})
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {

		//g.numPanel00.CurValue
		if g.IntGrid.MoveCursorFreely(0, 1, []int{0, 2, 3, 4}) {
			g.IntGrid.BoardPosition.Y += (g.IntGrid.Tile_Size.Y + g.IntGrid.Margin.Y)
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		// g.IntGrid.Position.X -= 1

		if g.IntGrid.MoveCursorFreely(3, 1, []int{0, 2, 3, 4}) {
			//g.IntGrid.Position.X += 18
			g.IntGrid.BoardPosition.X += (g.IntGrid.Tile_Size.X + g.IntGrid.Margin.X)
		}

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		// g.IntGrid.Position.Y += 1
		if g.IntGrid.MoveCursorFreely(2, 1, []int{0, 2, 3, 4}) {
			g.IntGrid.BoardPosition.Y -= (g.IntGrid.Tile_Size.Y + g.IntGrid.Margin.Y)

		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.IntGrid.MoveCursorFreely(1, 1, []int{0, 2, 3, 4}) {
			g.IntGrid.BoardPosition.X -= (g.IntGrid.Tile_Size.X + g.IntGrid.Margin.X)
		}

		// g.IntGrid.Position.X += 1
	}

	go g.IntGrid.UpdateOnMouseEvent()
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
	game := &Game{}
	ebiten.SetWindowSize(Settings.WindowSizeX, Settings.WindowSizeY)
	ebiten.SetWindowTitle("Grid Tile Ebitengine Demo")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
