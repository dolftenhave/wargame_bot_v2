package wargame

import (
	"encoding/json"
	"os"
)

type (
	Mode struct {
		Name           string `json:"name"`
		MapListID      []int  `json:"maps"`
		TeamSize       int    `json:"teamSize"`
		StartingPoints int    `json:"startingPoints"`
		TimeLimit      int    `json:"timeLimit"`
		ScoreLimit     int    `json:"scoreLimit"`
		Income         int    `json:"income"`
		GameMode       int    `json:"gameMode"`
		Oposotion      int    `json:"oposotion"`
		Nations        int    `json:"nations"`
		Era            int    `json:"era"`
		Theme          int    `json:"theme"`
		AutoStart      bool   `json:"autoStart"`
		MinPlayers     int    `json:"minPlayers"`
		WarmupTime     int    `json:"warmupTime"`
		DeployTime     int    `json:"deployTime"`
		DebriefTime    int    `json:"debriefTime"`
		LoadingTime    int    `json:"loadingTime"`
	}

	ModeList []Mode
)

// Reads the file and returns all the modes as a mode list.
func (modes *ModeList) UnmarshalFile(filePath string, maps MapList) error {
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonFile, modes)
	if err != nil {
		return err
	}
	return nil
}

// Returns he maps for this mode from the porvided MapList
func (m Mode) GetMaps(maps MapList) MapList {
	var result MapList
	for _, id := range m.MapListID {
		result = append(result, maps[id])
	}
	return result
}
