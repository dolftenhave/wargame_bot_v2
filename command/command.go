package command

import "wargame-bot/store"

type PermissionLevel int

const (
	// Has access to a finite set of commands
	PermEveryone PermissionLevel = iota
	// Has access to a greater set of commands than every else
	PermModerator
	// Has access to every command
	PermAdmin
)

// The caller of the command
type Caller struct {
	// (player id) The id of the discord user or wargame user
	ID int
	// (Player name) The id of the discord user or wargame user
	Name string
	// Permission level (everyone, mod or admin)
	Permission PermissionLevel
	// discord or server_chat
	Source string
}

// Provides context to each command
type CommandContext struct {
	Caller Caller
	Args   []string
	// Access to the repo interface
	Store store.Store
	// Access to the rcon client interface
	Rcon RconClient
}

// An interface for the rcon client
type RconClient interface {
	Execute(command string) (string, error)
	Say(to, from, msg string) error
}

// The command handler
type HandlerFunc func(ctx CommandContext) CommandResult

// A single command with meta data
type Command struct {
	// The command name
	Name string
	// A description of what the command does.
	Description string
	// The required permission in order to use the command
	Permission PermissionLevel
	// A pointer to the command handler
	Handler HandlerFunc
}
