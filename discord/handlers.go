// Contains all the command handlers.

package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Send an error message to discord.
func SomethingWentWrong(c Context, message string) {
	content := "Oh no, something went wrong."
	if message != "" {
		content += fmt.Sprintf("\nMessage: ", message)
	}
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

// Handles interactions with the help command.
func HelpHandler(c Context) {
	options := c.Interaction.Interaction.ApplicationCommandData().Options
	if options == nil || options[0].Name == "" {
		c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Help",
						Description: "This bot lets you interact with the wargame server.\n\n**Important!** Please read the <#1452438914215313558> before using it.",
						Color:       0xCF574A,
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:   "Help Command",
								Value:  "Using `/help` will bring up this message.\nOptionaly, you can add the name of another command for help on how to use it. `/help <command name>`.",
								Inline: false,
							},
							{
								Name:   "Where can I use the bot?",
								Value:  "The bot will only work in <#1445051378304028682>.",
								Inline: false,
							},
						},
					},
				},
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})
	} else {
		//TODO
		CommandNotImplemented(c.Session, c.Interaction)
	}
}

// Handles the mode command
func ModeHandler(c Context) {
	options := c.Interaction.Interaction.ApplicationCommandData().Options
	if options == nil {
		SomethingWentWrong(c, "Handler Options is nil")
		return
	}
	switch options[0].Name {
	case "list":
		listModes(c)
	case "set":
		setModes(c)
	default:
		SomethingWentWrong(c, fmt.Sprintf("Mode option '%s' is not yet implemented", options[0].Name))
	}
}

// Sends the user a list of available modes in text format.
func listModes(c Context) {
	var fields = make([]*discordgo.MessageEmbedField, len(c.Wargame.GameModes))
	for i, mode := range c.Wargame.GameModes {
		var name string
		if mode.Name == c.Wargame.Server.Mode.Name {
			name = fmt.Sprintf("__[x] - %s__", mode.Name)
		} else {
			name = fmt.Sprintf("[ ] - %s", mode.Name)
		}

		field := &discordgo.MessageEmbedField{
			Name:   name,
			Value:  fmt.Sprintf("Team Size: %v", mode.TeamSize),
			Inline: false,
		}
		fields[i] = field
	}
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Modes",
					Description: "Here is a list of available modes.",
					Fields:      fields,
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}

// Lets you set the current game mode.
func setModes(c Context) {
	CommandNotImplemented(c.Session, c.Interaction)
}

// Handles the mode command
func MapHandler(c Context) {
	options := c.Interaction.Interaction.ApplicationCommandData().Options
	if options == nil {
		SomethingWentWrong(c, "Handler Options is nil")
		return
	}
	switch options[0].Name {
	case "list":
		listMap(c)
	case "set":
		setMap(c)
	case "vote":
		voteMap(c)
	case "random":
		randomMap(c)
	default:
		SomethingWentWrong(c, fmt.Sprintf("Mode option '%s' is not yet implemented", options[0].Name))
	}
}

func listMap(c Context) {
	SomethingWentWrong(c, "")
}

func setMap(c Context) {
	SomethingWentWrong(c, "")
}

func voteMap(c Context) {
	SomethingWentWrong(c, "")
}

func randomMap(c Context) {
	SomethingWentWrong(c, "")
}
