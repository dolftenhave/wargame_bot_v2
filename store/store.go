package store

import (
	"context"
	"wargame-bot/domain"
)

type Store interface {
	WargamePlayerStore
	DiscordPlayerStore
	MapStore
	ModeStore
	IconStore
	CommandStore
}

type DiscordPlayerStore interface {
	NewDiscordPlayer(ctx context.Context, player *domain.DiscordPlayer) error
	LinkDiscordPlayer(ctx context.Context, player *domain.DiscordPlayer, player2 *domain.WargamePlayer) error
	GetDiscordPlayerByName(ctx context.Context, name string) ([]domain.DiscordPlayer, error)
	UpdateDiscordPlayerName(ctx context.Context, player *domain.DiscordPlayer, newName string) error
	DeleteDiscordPlayer(ctx context.Context, player *domain.DiscordPlayer) error
}

type WargamePlayerStore interface {
	NewWargamePlayer(ctx context.Context, player *domain.WargamePlayer) error
	LinkWargamePlayer(ctx context.Context, player *domain.WargamePlayer, player2 *domain.DiscordPlayer) error
	GetWargamePlayerByName(ctx context.Context, name string) ([]domain.WargamePlayer, error)
	UpdateWargamePlayerName(ctx context.Context, player *domain.WargamePlayer, newName string) error
	DeleteWargamePlayer(ctx context.Context, player *domain.WargamePlayer) error
	IsLinkedToDiscordPlayer(ctx context.Context, player *domain.WargamePlayer) (bool, error)
	GetDiscordLink(ctx context.Context, player *domain.WargamePlayer) (*domain.DiscordPlayer, error)
}

type MapStore interface {
	// All maps that contain the substring "name"
	GetMapByName(ctx context.Context, name string) ([]domain.Map, error)
	//All maps that contain the substring code
	GetMapByCode(ctx context.Context, code string) ([]domain.Map, error)
	// All maps that are of Size size
	GetMapBySize(ctx context.Context, size int) ([]domain.Map, error)
	// All maps that are of Kind kind
	GetMapByKind(ctx context.Context, size int) ([]domain.Map, error)
}

type ModeStore interface {
	AddMode(ctx context.Context, newMode *domain.Mode) error
	UpdateMode(ctx context.Context, mode *domain.Mode) error
	DeleteMode(ctx context.Context, mode *domain.Mode) error

	GetModeByName(ctx context.Context, name string) error
}

type IconStore interface {
	GetIconById(ctx context.Context, id int) (*domain.Icon, error)
	GetIconByName(ctx context.Context, name string) (*domain.Icon, error)
	GetAllIcons(ctx context.Context) ([]domain.Icon, error)
}

type CommandStore interface {
	// TODO implement the command interface
}
