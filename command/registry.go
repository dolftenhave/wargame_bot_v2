package command

import (
	"fmt"
	"log"
	"wargame-bot/rcon"
	"wargame-bot/store"
)

// A collection of all registered commands
type Registry struct {
	commands map[string]*Command
	store    store.Store
	rcon     rcon.Client
}

// Create a new command registry
func NewRegistry(store store.Store, rcon rcon.Client) *Registry {
	return &Registry{
		commands: make(map[string]*Command),
		store:    store,
		rcon:     rcon,
	}
}

// Register a command
func (r *Registry) Register(cmd ...*Command) {
	// TODO complete
	for _, c := range cmd {
		// TODO add collision check and return something
		// Add aliases?
		r.commands[c.Name] = c
		log.Printf("[Registry] Registered command: %s", c.Name)
	}
}

// Execute a command if it exists.
func (r *Registry) Execute(name string, caller Caller, args []string) CommandResult {
	cmd, exists := r.commands[name]

	if !exists {
		return CommandResult{
			Type:    ResultError,
			Message: fmt.Sprintf("Unknown Command: %s", name),
		}
	}

	// check command permission.
	if caller.Permission < cmd.Permission {
		return CommandResult{
			Type:    ResultError,
			Message: "Sorry, you do not have permission to use this command.",
		}
	}

	ctx := CommandContext{
		Caller: caller,
		Args:   args,
		Store:  r.store,
		Rcon:   r.rcon,
	}
	// Execute the command handler
	return cmd.Handler(ctx)
}

// Return a registered command
func (r *Registry) GetCommand(name string) (*Command, bool) {
	cmd, exists := r.commands[name]
	return cmd, exists
}

// Return a list of all registered commands
func (r *Registry) ListCommands() []*Command {
	var cmds []*Command
	for _, cmd := range r.commands {
		cmds = append(cmds, cmd)
	}
	return cmds
}
