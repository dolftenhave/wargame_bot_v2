package domain

// Server represents the current wargame server state.
type Server struct {
	State int
	Players []WargamePlayer
	Map []Map
}
