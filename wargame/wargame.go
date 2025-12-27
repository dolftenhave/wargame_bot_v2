package wargame

type (
	// A collection of components used for wargame.
	Wargame struct {
		GameModes ModeList
		Maps      MapList
		Server    Server
	}
)

// Initialises a new wargame variable.
func NewWargame(modeDataPath string, mapDataPath string, rconConfig RconConfig) (*Wargame, error) {
	var wargame = new(Wargame)

	err := wargame.Maps.ReadConfig(mapDataPath)
	if err != nil {
		return nil, err
	}

	err = wargame.GameModes.ReadConfig(modeDataPath, &wargame.Maps)
	if err != nil {
		return nil, err
	}

	err = wargame.Server.CreateConn(&rconConfig)
	if err != nil {
		return nil, err
	}

	for _, mode := range wargame.GameModes {
		wargame.Server.Mode = &mode
		break
	}
	return wargame, nil
}
