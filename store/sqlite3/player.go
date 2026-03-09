package sqlite3

import (
	"context"
	"wargame-bot/domain"
)

// Register a new player
func (s *Sqlite3Store) RegisterPlayer(ctx context.Context, player *domain.Player) error {
	return nil
}

func (s *Sqlite3Store) LinkDiscord(ctx context.Context, playerID int, discordID int) error {
	return nil
}
func (s *Sqlite3Store) GetPlayerByName(ctx context.Context, name string) ([]domain.Player, error) {
	return nil, nil
}
func (s *Sqlite3Store) UpdatePlayerName(ctx context.Context, playerID int, newName string) error {
	return nil
}
func (s *Sqlite3Store) DeletePlayer(ctx context.Context, playerID int) error          { return nil }
func (s *Sqlite3Store) IsLinked(ctx context.Context, playerID int) (bool, error)      { return false, nil }
func (s *Sqlite3Store) GetDiscordLink(ctx context.Context, playerID int) (int, error) { return 0, nil }
func (s *Sqlite3Store) GetWargameId(ctx context.Context, discordID int) (int, error)  { return 0, nil }
