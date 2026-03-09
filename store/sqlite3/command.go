package sqlite3

import (
	"context"
	"wargame-bot/domain"
)

func (s *Sqlite3Store) AddCommand(ctx context.Context, commandName string) error    { return nil }
func (s *Sqlite3Store) RemoveCommand(ctx context.Context, commandName string) error { return nil }
func (s *Sqlite3Store) ListCommands(ctx context.Context) ([]domain.Command, error)  { return nil, nil }
func (s *Sqlite3Store) FindCommand(ctx context.Context, commandName string) ([]domain.Command, error) {
	return nil, nil
}
