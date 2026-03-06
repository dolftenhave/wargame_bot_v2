package sqlite3

import (
	"context"
	"wargame-bot/domain"
)

const (
	playerTable = "wargame_player"
	idColumn    = "id"
	nameColumn  = "name"
)

// Create a new wargame player
func (sqlite3 *Sqlite3Store) NewDiscordPlayer(ctx context.Context, player *domain.DiscordPlayer) error {
	_, err := sqlite3.db.ExecContext(ctx, "INSERT INTO ?(?, ?) VALUES('?', '?');", playerTable, idColumn, nameColumn, player.ID, player.Name)
	return err
}

// Links the accounts of a discord player and a wargame player
func (sqlite3 *Sqlite3Store) LinkDiscordPlayer(ctx context.Context, wargame_player *domain.DiscordPlayer, discord_player *domain.DiscordPlayer) error {
	return nil
}

// Gets the players details by name
func (sqlite3 *Sqlite3Store) GetDiscordPlayerByName(ctx context.Context, name string) ([]domain.DiscordPlayer, error) {
	return nil, nil
}

// Updates the name of a player
func (sqlite3 *Sqlite3Store) UpdateDiscordPlayerName(ctx context.Context, player *domain.DiscordPlayer, newName string) error {
	return nil
}

// Deletes a player from the database
func (sqlite3 *Sqlite3Store) DeleteDiscordPlayer(ctx context.Context, player *domain.DiscordPlayer) error {
	return nil
}
