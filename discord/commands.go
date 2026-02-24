// This file contains all the commands exposed to discord that let the user interact with the wargame bot.

package discord

import "github.com/bwmarrin/discordgo"

// A list of all commands.
// Options. The name of another command, this will list a clearer desrcrioption of what the command does.
func HelpCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "help",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Lists all commands and what they do",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "command",
				Description: "The command you want help for.",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "deck",
						Value: "deck",
					},
					{
						Name:  "help",
						Value: "help",
					},
					{
						Name:  "map",
						Value: "map",
					},
					{
						Name:  "mode",
						Value: "mode",
					},
				},
				Required: false,
			},
		},
	}
}

// Set Deck
// Replies to the command with the emojis's of the nation, sepc and era of the deck.
// If it is not a deck then the command says so.
func DeckCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "deck",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Set your deck.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "set",
				Description: "Set your deck.",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
				// TODO integrate permissions.
			},
			{
				Name:        "decode",
				Description: "Descodes a deck for you.",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "code",
						Description: "The deck code.",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
					},
				},
			},
		},
	}
}

// Wargame Pannel
// This will bring up a complex pannel for managing the state of the server.
func PannelCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{}
}

// Set Mode
// This will let you set and see the mode of the current server.
func ModeCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "mode",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Set or change the mode of the server.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "set",
				Description: "Set the mode to one of the available options.",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
				// TODO integrate permissions.
			},
			{
				Name:        "list",
				Description: "See all the available modes.",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
			},
		},
	}
}

// Set Map
//
// Sub Commands.
//   - Auto complete list for all maps in the mode.
//   - Auto a random map.
func MapCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "map",
		Type:        discordgo.ChatApplicationCommand,
		Description: "See or change the current map.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "set",
				Description: "Set the map.",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
				// TODO integrate permissions.
			},
			{
				Name:        "list",
				Description: "See a list of available maps.",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
			},
			{
				Name:        "vote",
				Description: "Start a map vote.",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
			},
			{
				Name:        "random",
				Description: "Sets a random map.",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
			},
		},
	}
}

// Ban a player from the server.
// ( For now just permanently)
func Ban() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand {
		Name: "ban", 
		Type: discordgo.ChatApplicationCommand,
		Description:  "ban a player",
	}
}


// Kick a player from the server.
// ( For now just permanently)
func Kick() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand {
		Name: "kick", 
		Type: discordgo.ChatApplicationCommand,
		Description:  "kick a player",
	}
}

// Unban a player from the server
func UnBan() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name: "unban",
		Type: discordgo.ChatApplicationCommand,
		Description: "unban a player",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: "id",
				Description: "Player ID",
				Type: discordgo.ApplicationCommandOptionInteger,
				Required: true,
			},
		},
	}
}

// Send a message to the wargame server.
func Say() *discordgo.ApplicationCommand {
	return & discordgo.ApplicationCommand{
		Name: "say",
		Type: discordgo.ChatApplicationCommand,
		Description: "Send a message to the server",
		Options: []* discordgo.ApplicationCommandOption{
			{
				Name: "msg",
				Description: "Your message",
				Type: discordgo.ApplicationCommandOptionString,
				Required: true,
			},
		},
	}
}
