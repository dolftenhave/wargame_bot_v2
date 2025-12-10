package wargame

import (
	"encoding/json"
	"os"
)

type (
	// Contains the map data.
	Map struct {
		Type  int    `csv:"type"`
		Name  string `csv:"name"`
		Code  string `csv:"code"`
		Image string `csv:"image,omitempty"`
	}

	MapList []Map
)

// Reads the file and returns all the maps as a maplist.
func (maps MapList) Unmarshal(filePath string) error {
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonFile, maps)
	if err != nil {
		return err
	}

	return nil
}
