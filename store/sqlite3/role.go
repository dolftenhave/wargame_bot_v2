package sqlite3

import (
	"context"
	"wargame-bot/domain"
)


func (s *Sqlite3Store) ListRoles(ctx context.Context) ([]domain.Role, error) { return nil,nil }
func (s *Sqlite3Store) CreateRole(ctx context.Context, name string, level int) (*domain.Role, error) { return nil,nil }
func (s *Sqlite3Store) AddPlayerToRole(ctx context.Context, playerID int, roleID int) error { return nil }
func (s *Sqlite3Store) RemovePlayerFromRole(ctx context.Context, playerID int, roleID int) error { return nil }
func (s *Sqlite3Store) ListPlayerRoles(ctx context.Context, playerID int) ([]domain.Role, error) { return nil, nil }
func (s *Sqlite3Store) ListCommandsInRole(ctx context.Context) ([]domain.Command, error) { return nil, nil }
func (s *Sqlite3Store) AddCommandToRole(ctx context.Context, roleID int, commandName string) error { return nil }
func (s *Sqlite3Store) RemoveCommandFromRole(ctx context.Context, roleID int, commandName string) error { return nil }
