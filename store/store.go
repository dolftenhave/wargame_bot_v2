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
	NationIconStore
	SpecializationIconStore
	EraIconStore
	CommandStore
	RoleStore
}

type DiscordPlayerStore interface {
	NewDiscordPlayer(ctx context.Context, player *domain.DiscordPlayer) error
	LinkDiscordPlayer(ctx context.Context, player *domain.DiscordPlayer, player2 *domain.Player) error
	GetDiscordPlayerByName(ctx context.Context, name string) ([]domain.DiscordPlayer, error)
	UpdateDiscordPlayerName(ctx context.Context, player *domain.DiscordPlayer, newName string) error
	DeleteDiscordPlayer(ctx context.Context, player *domain.DiscordPlayer) error
}

type WargamePlayerStore interface {
	NewWargamePlayer(ctx context.Context, player *domain.Player) error
	LinkWargamePlayer(ctx context.Context, player *domain.Player, player2 *domain.DiscordPlayer) error
	GetWargamePlayerByName(ctx context.Context, name string) ([]domain.Player, error)
	UpdateWargamePlayerName(ctx context.Context, player *domain.Player, newName string) error
	DeleteWargamePlayer(ctx context.Context, player *domain.Player) error
	IsLinkedToDiscordPlayer(ctx context.Context, player *domain.Player) (bool, error)
	GetDiscordLink(ctx context.Context, player *domain.Player) (*domain.DiscordPlayer, error)
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

	GetModeByName(ctx context.Context, name string) (*domain.Mode, error)
	GetDefaultMode(ctx context.Context) (*domain.Mode, error)
	// Gets all the maps in the mode's map pool.
	GetMaps(ctx context.Context) ([]domain.MapPoolItem, error)
}

type NationIconStore interface {
	// Get the nation icon that matches the id, or an error if it does not exist
	GetNationIcon(ctx context.Context, id int) (*domain.Icon, error)
	// Get all the nation icons
	GetAllNationIcons(ctx context.Context) ([]domain.Icon, error)
}


type SpecializationIconStore interface {
	// Get the specialization icon that matches the id, or an error if it does not exist
	GetSpecializationIcon(ctx context.Context, id int) (*domain.Icon, error)
	// Get all the specialization icons
	GetAllSpecializationIcons(ctx context.Context) ([]domain.Icon, error)
}

type EraIconStore interface {
	// Get the era icon that matches the id, or an error if it does not exist
	GetEraIcon(ctx context.Context, id int) (*domain.Icon, error)
	// Get all the era icons
	GetAllEraIcons(ctx context.Context) ([]domain.Icon, error)
}

type CommandStore interface {
	// TODO implement the command interface
}

type RoleStore interface {
	// TODO Implement interface
	ListRoles(ctx context.Context) ([]domain.Role, error)
	CreateRole(ctx context.Context, name string, level int) (*domain.Role, error)
}
