CREATE TABLE IF NOT EXISTS game_mode(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	mode_name TEXT NOT NULL DEFAULT 'New Server',
	server_name TEXT NOT NULL DEFAULT '[Tactical] New Server',
	starting_points INTEGER NOT NULL DEFAULT 5500,
	time_limit INTEGER NOT NULL DEFAULT 3000,
	score_limit INTEGER NOT NULL DEFAULT 8000,
	income_rate INTEGER NOT NULL DEFAULT 4,
	game_mode INTEGER NOT NULL DEFAULT 1,
	oposition INTEGER NOT NULL DEFAULT 0,
	nations INTEGER NOT NULL DEFAULT -1,
	era INTEGER NOT NULL DEFAULT -1,
	theme INTEGER NOT NULL DEFAULT -1,
	team_size INTEGER NOT NULL DEFAULT 10,
	min_players INTEGER NOT NULL DEFAULT 19,
	warmup_time INTEGER NOT NULL DEFAULT 15,
	deploy_time INTEGER NOT NULL DEFAULT 180,
	debrief_time INTEGER NOT NULL DEFAULT 60,
	loading_time INTEGER NOT NULL DEFAULT 60, 

	auto_start INTEGER NOT NULL DEFAULT 0,
	auto_rotate INTEGER NOT NULL DEFAULT 0,
	map_vote INTEGER NOT NULL DEFAULT 0,
	enable_commands INTEGER NOT NULL DEFAULT 0
);

-- A table that contains maps currently in the modes map pool.
-- The maps may have overides for the current modes settings.
CREATE TABLE IF NOT EXISTS map_pool(
	mode_id INTEGER NOT NULL,
	map_id TEXT NOT NULL,
	income_rate INTEGER,
	starting_points INTEGER,
	score_limit INTEGER,
	time_limit INTEGER,

	PRIMARY KEY (mode_id, map_id),
	FOREIGN KEY (mode_id) REFERENCES mode(id) ON DELETE CASCADE,
	FOREIGN KEY (map_id) REFERENCES map(id) ON DELETE CASCADE
);

-- A table that conains data about a wargame map.
CREATE TABLE IF NOT EXISTS map(
	id TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	image TEXT,
	kind TEXT NOT NULL DEFAULT 'Land',
	intended_size TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS nation(
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	code TEXT NOT NULL,
	emote_id TEXT
);

CREATE TABLE IF NOT EXISTS specialization(
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	code TEXT NOT NULL,
	emote_id TEXT
);

CREATE TABLE IF NOT EXISTS era(
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	code TEXT NOT NULL,
	emote_id TEXT
);

-- A player using the game. They could link their discord account to their wargame account.
CREATE TABLE IF NOT EXISTS player(
	wargame_id INTEGER PRIMARY KEY,
	discord_id INTEGER,
	name TEXT NOT NULL
);

-- A table that displays the name history for a player.
CREATE TABLE IF NOT EXISTS name_history(
	wargame_id INTEGER NOT NULL,
	name TEXT,
	
	PRIMARY KEY (wargame_id, name),
	FOREIGN KEY (wargame_id) REFERENCES wargame_player(id) ON DELETE CASCADE
);

-- A list of all roles.
CREATE TABLE IF NOT EXISTS role(
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL UNIQUE
); 

-- A table that groups roles and members
CREATE TABLE IF NOT EXISTS role_member(
	player_id INTEGER NOT NULL,
	role_id INTEGER NOT NULL,

	PRIMARY KEY (player_id, role_id),
	FOREIGN KEY (player_id) REFERENCES player(wargame_id) ON DELETE CASCADE,
	FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE
);

-- A list of all registered text commands.
CREATE TABLE IF NOT EXISTS command(
	name TEXT PRIMARY KEY		
	description TEXT
);

-- The commands that the role has access too.
CREATE TABLE IF NOT EXISTS role_command(
	role_id INTEGER NOT NULL,
	command_name TEXT NOT NULL,

	PRIMARY KEY (role_id, command_name),
	FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
	FOREIGN KEY (command_name) REFERENCES command(name) ON DELETE CASCADE
);

-- The current server state
CREATE TABLE IF NOT EXISTS server_state(
	id INTEGER PRIMARY KEY,
	mode_id INTEGER NOT NULL,
	map_id INTEGER NOT NULL,
	state INTEGER NOT NULL,

	FOREIGN KEY (mode_id) REFERENCES mode(id),
	FOREIGN KEY (map_id) REFERENCES map(id)
);

-- A table that contains the id's of all the players currently online.
CREATE TABLE IF NOT EXISTS online_players(
	player_id INTEGER PRIMARY KEY,
	team INTEGER NOT NULL CHECK(team == 0 OR team == 1),

	FOREIGN KEY (player_id) REFERENCES player(wargame_id)
);
