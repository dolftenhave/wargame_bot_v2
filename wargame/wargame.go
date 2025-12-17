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
func InitWargame(modes ModeList, maps MapList, server Server) *Wargame {
	return &Wargame{
		GameModes: modes,
		Maps:      maps,
		Server:    server,
	}
}
