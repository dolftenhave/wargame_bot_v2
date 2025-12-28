package wargame

type (
	// A collection of components used for wargame.
	Wargame struct {
		GameModes    ModeList
		Maps         MapList
		Server       Server
		DeckCodeData DeckCodeData
	}
)

// Initialises a new wargame variable.
func NewWargame(modeDataPath string, mapDataPath string, rconConfig RconConfig, deckDataPath string) (*Wargame, error) {
	var wargame = new(Wargame)

	err := wargame.Maps.ReadConfig(mapDataPath)
	if err != nil {
		return nil, err
	}

	err = wargame.GameModes.ReadConfig(modeDataPath, &wargame.Maps)
	if err != nil {
		return nil, err
	}

	err = wargame.DeckCodeData.ReadConfig(deckDataPath)
	if err != nil {
		return nil, err
	}

	err = wargame.Server.CreateConn(&rconConfig)
	if err != nil {
		return nil, err
	}

	wargame.Server.RconConfig = rconConfig

	for _, mode := range wargame.GameModes {
		wargame.Server.Mode = &mode
		break
	}

	err = wargame.Server.SetMode(wargame.Server.Mode)
	if err != nil {
		return nil, err
	}

	return wargame, nil
}
