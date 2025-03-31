package mypkgs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type GameSettings struct {
	VersionID   string `json:"versionID,omitempty"`
	WindowSizeX int    `json:"window_size_x"`
	WindowSizeY int    `json:"window_size_y"`
	ScreenResX  int    `json:"screen_res_x,"`
	ScreenResY  int    `json:"screen_res_y,"`
}

/*
This will load from a JSON file;
*/
func GetBytesFromJSON(filePath string) []byte {
	fmt.Print("INIT JSON HELLO!\n\n")
	jSonFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := jSonFile.Close(); err != nil {
			panic(err)
		}
	}()
	var rdal []byte
	rdr := bufio.NewReader(jSonFile)
	rdal, err = io.ReadAll(rdr)
	if err != nil {
		panic(err)
	}
	return rdal
}

//	func(gSets *GameSettings) GetSettingsFromJSON(){
//		get
//	}
func GetSettingsFromJSON() GameSettings {
	var gSets GameSettings
	err := json.Unmarshal(GetBytesFromJSON("init.JSON"), &gSets)
	if err != nil {
		panic(err)
	}
	return gSets
}
func (sets *GameSettings) ToString() string {
	return fmt.Sprintf("SETTINGS:\n%12s: %s\n%12s: %3d, %3d\n%12s: %3d,%3d\n", "VERSION", sets.VersionID, "Window Size", sets.WindowSizeX, sets.WindowSizeY, "Screen Res", sets.ScreenResX, sets.ScreenResY)
}

func GetSettingsFromBakedIn() GameSettings {
	var gSets GameSettings = GameSettings{
		VersionID:   "0.0.08",
		WindowSizeX: 960, //860//892
		WindowSizeY: 640, //660 //720
		ScreenResX:  960, //860 //892
		ScreenResY:  640,
	}
	return gSets
}
