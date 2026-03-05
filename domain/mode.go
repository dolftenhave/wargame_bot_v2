package domain

// Mode represends a the mode the servere is currently running on.
type Mode struct {
	ID int
	Name string
	ServerName string
	StartingPoints int
	TimeLimit int
	ScoreLimit int
	IncomeRate int
	GameMode int
	Oposition int
	Nation int
	Era int
	Theme int
	TeamSize int
	MinPlayers int
	WarmupTime int
	DeployTime int
	DebriefTime int
	LoadingTime int

	AutoStart bool
	AutoRotate bool
	MapVote bool
	EnableCommands bool
}
