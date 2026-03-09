package store

import (
	"context"
	"wargame-bot/domain"
)

type Store interface {
	PlayerStore
	MapStore
	ModeStore
	NationIconStore
	SpecializationIconStore
	EraIconStore
	CommandStore
	RoleStore
}

type PlayerStore interface {
	RegisterPlayer(ctx context.Context, player *domain.Player) error
	LinkDiscord(ctx context.Context, playerID int, discordID int) error
	GetPlayerByName(ctx context.Context, name string) ([]domain.Player, error)
	UpdatePlayerName(ctx context.Context, playerID int, newName string) error
	DeletePlayer(ctx context.Context, playerID int) error
	IsLinked(ctx context.Context, playerID int) (bool, error)
	GetDiscordLink(ctx context.Context, playerID int) (int, error)
	GetWargameId(ctx context.Context, discordID int) (int, error)
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

	GetModeByName(ctx context.Context, name string) ([]domain.Mode, error)
	GetDefaultMode(ctx context.Context) (*domain.Mode, error)
	ListModes(ctx context.Context) ([]domain.Mode, error)
	// Gets all the maps in the current mode's map pool.
	GetMaps(ctx context.Context) ([]domain.MapPoolItem, error)
	AddMap(ctx context.Context, mapItem *domain.MapPoolItem) error
	RemoveMap(ctx context.Context, mapCode string) error
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
	AddCommand(ctx context.Context, commandName string) error
	RemoveCommand(ctx context.Context, commandName string) error
	ListCommands(ctx context.Context) ([]domain.Command, error)
	FindCommand(ctx context.Context, commandName string) ([]domain.Command, error)
}

type RoleStore interface {
	ListRoles(ctx context.Context) ([]domain.Role, error)
	CreateRole(ctx context.Context, name string, level int) (*domain.Role, error)
	AddPlayerToRole(ctx context.Context, playerID int, roleID int) error
	RemovePlayerFromRole(ctx context.Context, playerID int, roleID int) error
	ListPlayerRoles(ctx context.Context, playerID int) ([]domain.Role, error)
	ListCommandsInRole(ctx context.Context) ([]domain.Command, error)
	AddCommandToRole(ctx context.Context, roleID int, commandName string) error
	RemoveCommandFromRole(ctx context.Context, roleID int, commandName string) error
}
