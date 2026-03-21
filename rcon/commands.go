package rcon

import (
	"context"
	"fmt"
)

// Kick a player from the server using their player id or name.
func (c *rconClient) Kick(ctx context.Context, playerID string) (string, error) {
	var cmd = fmt.Sprintf("kick %s", playerID)
	return c.Execute(ctx, cmd)
}

// Kick a player from the server using their player id or name.
func (c *rconClient) Ban(ctx context.Context, id string, hours int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("ban %s %v", id, hours))
}

// Unban a player from the server using their player id or name.
func (c *rconClient) UnBan(ctx context.Context, id string) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("unban %s", id))
}

// Sets the name of the server
func (c *rconClient) SetServerName(ctx context.Context, name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Name can not be empty")
	}
	return c.Execute(ctx, fmt.Sprintf("setsvar ServerName %s", name))
}

// Sets the number of players in the server
func (c *rconClient) SetNumPlayers(ctx context.Context, n int) (string, error) {
	if n <= 2 || n >= 20 {
		return "Players must be between 2 and 20", fmt.Errorf("Players must be between 2 and 20")
	}
	return c.Execute(ctx, fmt.Sprintf("setsvar NbMaxPlayer %v", n))
}

// Sets the game type of the server. e.g. red vs red.
func (c *rconClient) SetOposition(ctx context.Context, oposition int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar GameType %v", oposition))
}

// Sets the starting money.
func (c *rconClient) SetStartingPoints(ctx context.Context, points int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar InitMoney %v", points))
}

// Sets the income rate of the server.
func (c *rconClient) SetIncomeRate(ctx context.Context, money int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar IncomeRate %v", money))
}

// Sets the game duration in minutes.
func (c *rconClient) SetTimeLimit(ctx context.Context, duration int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar TimeLimit %v", duration))
}

// Sets the map. The mapName is the id of the map. (yes poorly named).
func (c *rconClient) SetMap(ctx context.Context, mapName string) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar Map %s", mapName))
}

// Sets the game mode. e.g. distruction, conquest.
func (c *rconClient) SetGameMode(ctx context.Context, mode int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar VictoryCond %v", mode))
}

// Sets the required number of players needed to start a game.
func (c *rconClient) SetMinPlayers(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar NbMinPlayer %v", n))
}

// Sets the starting countdown in seconds. This begins once the required number of players has been reached.
func (c *rconClient) SetStartCountdown(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar WarmupCountdown %v", n))
}

// Sets the maximum amount of time the server will wait until for players to connect before kicking those who are still connecting.
func (c *rconClient) SetLoadingTime(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar LoadingTimeMax %v", n))
}

// Sets the nation constriant for the server.
func (c *rconClient) SetNations(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar NationConstraint %v", n))
}

// Sets the nation constriant for the server.
func (c *rconClient) SetEra(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar DateConstraint %v", n))
}

// Set the theme of the game.
func (c *rconClient) SetTheme(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar ThematicConstraint %v", n))
}

// Set the warmup time.
func (c *rconClient) SetWarmupTime(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar WarmupCountdown %v", n))
}

// Set the loading time.
func (c *rconClient) SetLodingTime(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar LoadingTimeMax %v", n))
}

// Set the deployment time.
func (c *rconClient) SetDeployTime(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar DeploimentTimeMax %v", n))
}

// Set the debriefing time.
func (c *rconClient) SetDebriefTime(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar DebriefingTimeMax %v", n))
}

// Set a the score limit needed for the game to end.
func (c *rconClient) SetScoreLimit(ctx context.Context, n int) (string, error) {
	return c.Execute(ctx, fmt.Sprintf("setsvar ScoreLimit %v", n))
}

// Launch the game regardless of if the start condistions are met.
func (c *rconClient) Launch(ctx context.Context) (string, error) {
	return c.Execute(ctx, "launch")
}

// Cancel game launch
func (c *rconClient) CancelLaunch(ctx context.Context) (string, error) {
	return c.Execute(ctx, "cancel_launch")
}

// Get all the players that are currently logged in.
func (c *rconClient) GetPlayers(ctx context.Context) (string, error) {
	return c.Execute(ctx, "chat")
	// This command was changes to "chat" with the modiefied bonary
	//return "chat"
}

// Send a message to the server
func (c *rconClient) Chat(ctx context.Context, to string, from string, msg string) (string,error) {
	source_id := 0xffffffff
	dest_id := 0xffffffff

	return c.Execute(ctx, fmt.Sprintf("chat & '%x' '%x' '%s'", source_id, dest_id, msg))
}
