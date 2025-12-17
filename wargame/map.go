package wargame

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type (
	// Contains the map data.
	Map struct {
		Type  int    `csv:"type"`
		Name  string `csv:"name"`
		Code  string `csv:"code"`
		Image string `csv:"image,omitempty"`
	}

	// An array of maps
	MapList []Map
)

// Reads the file and adds all the maps to the maplist.
func (maps *MapList) ReadConfig(filePath string) error {
	csvFile, err := os.Open(filePath)


	if err != nil {
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Read()
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		t, err := strconv.Atoi(line[0])
		if err != nil {
			return fmt.Errorf("%s is not an int.\n%s", line[0], err)
		}

		var mapVar = Map{
			Type:  t,
			Name:  line[1],
			Code:  line[2],
			Image: line[3],
		}
		*maps = append(*maps, mapVar)
	}
	return nil
}

func (m Map) PrintMap() {
	fmt.Printf("%s, %s, %v, %s\n", m.Code, m.Name, m.Type, m.Image)
}
