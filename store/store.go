package store

import (
	"context"
	"wargame-bot/domain"
)

type Store interface {
	PlayerStore
	MapStore
	ModeStore
	IconStore
	CommandStore
}

type PlayerStore interface {
	NewPlayer(ctx context.Context, player *domain.Player) error
	LinkPlayer(ctx context.Context, player *domain.Player, player2 *domain.Player) error
	GetPlayerByName(ctx context.Context, name string) (error, []domain.Player)
	UpdatePlayerName(ctx context.Context, player *domain.Player, newName string) error
	DeletePlayer(ctx context.Context, player *domain.Player) error
}

type MapStore interface {
	GetMapByName(ctx context.Context, name string) (error, []domain.Map)
	GetMapByCode(ctx context.Context, code string) (error, []domain.Map)
	GetMapBySize(ctx context.Context, size int) (error, []domain.Map)
}

type ModeStore interface {
	AddMode(ctx context.Context, newMode *domain.Mode) error
	UpdateMode(ctx context.Context, mode *domain.Mode) error
	DeleteMode(ctx context.Context, mode *domain.Mode) error

	GetModeByName(ctx context.Context, name string) error
}

type IconStore interface {
	GetIconById(ctx context.Context, id int) (error, *domain.Icon)
	GetIconByName(ctx context.Context, name string) (error, *domain.Icon)
	GetAllIcons(ctx context.Context) (error, []domain.Icon)
}

type CommandStore interface {
	// TODO implement the command interface
}
