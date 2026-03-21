package command

import (
	"context"
	"fmt"
)

// Registers all the default rcon commands
func RegisterRconCommands(r *Registry) {
	r.Register(&Command{})
}

// Registers all the available commands
func RegisterAllCommands(r *Registry) {
	r.Register(
		&Command{
			Name:        "help",
			Description: "Shows a list of all available commands, what they do and how to use them",
			Permission:  PermEveryone,
			Handler:     handleHelpCommand,
		},
		&Command{
			Name:        "map_list",
			Description: "List available maps for the current mode.",
			Permission:  PermEveryone,
			Handler:     handleMapList,
		},
	)
}

// Handles the help command
func handleHelpCommand(ctx CommandContext) CommandResult {
	return CommandResult{
		Type:    ResultMessage,
		Message: "Help Command message",
	}
}

// Returns a list of maps for the current mode.
func handleMapList(ctx CommandContext) CommandResult {
	maps, err := ctx.Store.GetMapsForMode(context.Background())
	if err != nil {
		return CommandResult{Type: ResultError, Message: err.Error(), Error: err}
	}

	if len(maps) == 0 {
		return CommandResult{
			Type:    ResultMessage,
			Title:   "Maps",
			Message: "No maps available for this mode.",
		}
	}

	var fields []ResultField
	for _, m := range maps {
		fields = append(fields, ResultField{
			Name:  m.Name,
			Value: fmt.Sprintf("Map Code: %s", m.MapCode),
		})
	}

	return CommandResult{
		Type:   ResultList,
		Title:  "Available Maps",
		Fields: fields,
	}
}
