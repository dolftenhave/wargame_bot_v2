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
				Name:        "setdeck",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
				Description: "Sets your deck",
			},
			{
				Name:        "help",
				Type:        discordgo.ApplicationCommandOptionType(discordgo.ChatApplicationCommand),
				Description: "Lists all commands and what they do",
			},
		},
	}
}

// Set Deck
// Replies to the command with the emojis's of the nation, sepc and era of the deck.
// If it is not a deck then the command says so.
func SetDeck() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{}
}

// Wargame Pannel
// This will bring up a complex pannel for managing the state of the server.
func RCON() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{}
}

// Set Mode
// This will let you set and see the mode of the current server.
func SetMode() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "mode",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Set or change the mode of the server.",
	}
}

// Set Map
//
// Sub Commands.
//   - Auto complete list for all maps in the mode.
//   - Auto a random map.
func Map() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{}
}
