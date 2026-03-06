package sqlite3

import (
	"context"
	"wargame-bot/domain"
)


// Create a new wargame player
func (sqlite3 *Sqlite3Store) NewWargamePlayer(ctx context.Context, player *domain.WargamePlayer) error {
	return nil
}

// Links the accounts of a discord player and a wargame player
func (sqlite3 *Sqlite3Store) LinkWargamePlayer(ctx context.Context, player *domain.WargamePlayer, player2 *domain.WargamePlayer) error {
	return nil
}

// Gets the players details by name
func (sqlite3 *Sqlite3Store) GetWargamePlayerByName(ctx context.Context, name string) ([]domain.WargamePlayer, error) {
	return nil, nil
}

// Updates the name of a player
func (sqlite3 *Sqlite3Store) UpdateWargamePlayerName(ctx context.Context, player *domain.WargamePlayer, newName string) error {
	return nil
}

// Deletes a player from the database
func (sqlite3 *Sqlite3Store) DeleteWargamePlayer(ctx context.Context, player *domain.WargamePlayer) error { return nil }
