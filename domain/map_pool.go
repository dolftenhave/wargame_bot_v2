package domain

//MapPool represents a map that is in the map pool of a mode
type MapPool struct {
	ModeId int
	MapId int
	Name string
	IncomeRate int
	StartingPoints int
	ScoreLimit int
	TimeLimit int
}
