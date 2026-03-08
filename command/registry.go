package command

import (
	"fmt"
	"log"
	"wargame-bot/wargame"
)

// A collection of all registered commands
type Registry struct {
	commands map[string]*Command
	wargame *wargame.Wargame
}

// Create a new command registry
func NewRegistry(w *wargame.Wargame) *Registry {
	return &Registry{
		commands: make(map[string]*Command),
		wargame: w,
	}
}

// Register a command
func (r *Registry) Register(cmd *Command) {
	r.commands[cmd.Name] = cmd
	log.Printf("[Registry] Registered command: %s", cmd.Name)
}

// Execute a command if it exists.
func (r Registry) Execute(name string, caller Caller, args []string) CommandResult {
	cmd, exists := r.commands[name] 

	if !exists {
		return CommandResult {
			Type: ResultError,
			Message: fmt.Sprintf("Unknown Command: %s", name),
		}
	}

	if caller.Permission < cmd.Permission {
		return CommandResult{
			Type: ResultError,
			Message: "Sorry, you do not have permission to use this command.",
		}
	}

	// Execute the command handler
	return cmd.Handler(r.wargame, caller, args)
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
