package domain

// Map represents a wargame map
type Map struct {
	Code  string // The code of the map
	Name  string // The name of the map
	Image string // A link to the maps image
	Kind  string // The maps kind. Land, Mixed, Naval.
	Size  int	// The size of the map. 1: 1v1, 2: 2v2, 3: 3v3, 4: 4v4, 10: 10v10
}
