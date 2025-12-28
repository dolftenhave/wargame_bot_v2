package wargame

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

const TEST_CODE = "@HsEC3UOisqDSaVEoq6eUVdN6KuntFZUakkqBSY07pq6gValOKSOsLqmqoitxquaxmjRpsasOmSo6qyqjjpG60GkDqcqvqlzpSqqBWJSrEglZhqqasRIw"

const (
	// 1-11, 11 bits- 0111 1111 1111 0000
	NAT_MASK uint16 = 32752

	// 12-14, 3 bits - 0000 0111
	SPEC_MASK uint8 = 7

	// 15-16, 2 bits - 0000 0011
	ERA_MASK uint8 = 3
)

type (
	DeckData struct {
		Name string `json:"name"`

		Code string `json:"code"`

		DiscID string `json:"discord_id"`
		// The path to the icon.
		Icon string `json:"icon"`
	}

	DeckCodeData struct {
		Nations         map[uint16]DeckData `json:"nations"`
		Specializations map[uint8]DeckData  `json:"specializations"`
		Eras            map[uint8]DeckData  `json:"eras"`
	}
)


// Reads the data from the config
func (d *DeckCodeData) ReadConfig(filePath string) error {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, d)

	if err != nil {
		return err
	}
	return nil
}

// Decodes a wargame red dragon deck and returns an array of icon paths if there are any.
func DecodeDeck(deckCode string, data *DeckCodeData) ([]DeckData, error) {
	var icons []DeckData
	// 1. Break the deck code down into binary
	bytes, err := toByteArray(deckCode)
	if err != nil {
		return nil, err
	}

	// Find the nation bits 1-11
	decodeNation(bytes, data, &icons)
	decodeSpec(bytes, data, &icons)
	decodeEra(bytes, data, &icons)
	log.Println("")

	return icons, nil
}

func decodeEra(bytes []byte, data *DeckCodeData, icons *[]DeckData) {
	if len(bytes) < 1 {
		return
	}

	era := bytes[1] & 1
	era = ((bytes[2] >> 6) & 2) | era
	era = era & ERA_MASK
	fmt.Printf("era: %v", era)

	icon, found := data.Eras[era]
	if found {
		*icons = append(*icons, icon)
	}
}

func decodeSpec(bytes []byte, data *DeckCodeData, icons *[]DeckData) {
	if len(bytes) < 1 {
		return
	}

	spec := bytes[1] >> 1
	spec = spec & SPEC_MASK
	fmt.Printf("spec: %v", spec)

	icon, found := data.Specializations[spec]
	if found {
		*icons = append(*icons, icon)
	}
}

func decodeNation(bytes []byte, data *DeckCodeData, icons *[]DeckData) {
	if len(bytes) < 1 {
		return
	}

	nation := uint16(bytes[0])<<8 | uint16(bytes[1])
	nation = nation & NAT_MASK
	nation = nation >> 4
	fmt.Printf("nation: %v", nation)

	icon, found := data.Nations[nation]
	if found {
		*icons = append(*icons, icon)
	}
}

func toByteArray(deckCode string) ([]byte, error) {
	deckCode = strings.TrimPrefix(deckCode, "@")

	bytes, err := base64.StdEncoding.DecodeString(deckCode)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
