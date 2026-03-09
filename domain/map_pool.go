package domain

//MapPool represents a map that is in the map pool of a mode
type MapPoolItem struct {
	MapCode string
	Name string
	IncomeRate int
	StartingPoints int
	ScoreLimit int
	TimeLimit int
}
