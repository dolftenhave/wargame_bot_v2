package wargame

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	Mode struct {
		Name           string `json:"name"`
		MapListID      []*int `json:"maps"`
		MapList        MapList
		TeamSize       int  `json:"teamSize"`
		StartingPoints int  `json:"startingPoints"`
		TimeLimit      int  `json:"timeLimit"`
		ScoreLimit     int  `json:"scoreLimit"`
		Income         int  `json:"income"`
		GameMode       int  `json:"gameMode"`
		Oposotion      int  `json:"oposotion"`
		Nations        int  `json:"nations"`
		Era            int  `json:"era"`
		Theme          int  `json:"theme"`
		AutoStart      bool `json:"autoStart"`
		MinPlayers     int  `json:"minPlayers"`
		WarmupTime     int  `json:"warmupTime"`
		DeployTime     int  `json:"deployTime"`
		DebriefTime    int  `json:"debriefTime"`
		LoadingTime    int  `json:"loadingTime"`
	}

	// An array of Modes
	ModeList []Mode
)

// Reads the file and returns all the modes as a mode list.
func (modes *ModeList) ReadConfig(filePath string, maps *MapList) error {
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonFile, modes)
	if err != nil {
		return err
	}

	// Creates a map list with only the maps in the MapListID
	for _, mode := range *modes {
		mode.MapList = mode.GetMaps(*maps)
	}

	return nil
}

// Writes the current modes config out to the file filePath.
func (m ModeList) WriteConfig() error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = os.WriteFile("modes_writen.json", data, os.FileMode(0777))
	if err != nil {
		return err
	}
	return nil
}

// Returns he maps for this mode from the porvided MapList
func (m Mode) GetMaps(maps MapList) MapList {
	if len(maps) < 1 {
		fmt.Println("The map list is empty")
		return nil
	}

	var result MapList
	for _, id := range m.MapListID {
		result = append(result, maps[*id])
	}
	return result
}
