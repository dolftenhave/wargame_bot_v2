package db

import (
	"database/sql"
)

type (
	db struct {

	}

	Player struct {
		id int
		wargame_name
		discord_name
		verified
		last_played
		banned
		commands
	}
)

// Create a template for the tables.

// Player

// Map
// Mode
// Last_State
