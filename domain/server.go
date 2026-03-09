package domain

type Team int

const (
	Blue Team = iota
	Red
)

// Server represents the current wargame server state.
type Server struct {
	State int
	Players map[int]Team
	Map string
}
