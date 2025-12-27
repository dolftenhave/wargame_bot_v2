// Contains all the command handlers.

package discord

import (
	"fmt"
	"log"
	"wargame-bot/wargame"

	"github.com/bwmarrin/discordgo"
)

// Send an error message to discord.
func SomethingWentWrong(c Context, message string) {
	content := "Oh no, something went wrong."
	if message != "" {
		content += fmt.Sprintf("\nMessage: ", message)
	}

	log.Printf("[Discord] Error: %s", message)

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
// Sends a drop down list with available modes, or sets the mode if a mode was selected.
func setModes(c Context) {
	options := c.Interaction.Interaction.ApplicationCommandData().Options
	options = options[0].Options

	if options == nil || len(options) < 1 || options[0].Name == "" {
		if len(c.Wargame.GameModes) < 1 {
			SomethingWentWrong(c, "There are no game modes to choose from")
			return
		}

		log.Printf("[Discord] %s is selecting a mode.", c.User.GlobalName)

		var mo []discordgo.SelectMenuOption

		for _, m := range c.Wargame.GameModes {
			var label string
			if m.Name == c.Wargame.Server.Mode.Name {
				label = fmt.Sprintf("**%s**", m.Name)
			} else {
				label = m.Name
			}
			mo = append(mo, discordgo.SelectMenuOption{
				Label:       label,
				Value:       m.Name,
				Description: "",
				Default:     false,
			})
		}

		err := c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.SelectMenu{
								CustomID:    "select_mode",
								Placeholder: "Select a mode...",
								Options:     mo,
							},
						},
					},
				},
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})

		if err != nil {
			log.Printf("[Discord] Error: %s", err.Error())
		}
		return
	}

	var mode *wargame.Mode
	for _, m := range c.Wargame.GameModes {
		if m.Name == options[0].Value {
			mode = &m
			break
		}
	}

	if mode == nil {
		SomethingWentWrong(c, fmt.Sprintf("Mode '%s' doesn't exist", options[0].Name))
		return
	}

	c.Wargame.Server.SetMode(mode)
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Modes",
					Description: fmt.Sprintf("Mode set too %s", mode.Name),
				},
			},
		},
	})
}

// Sets the mode of the server
func SetModeHandler(c Context) {
	data := c.Interaction.Interaction.MessageComponentData()
	if len(data.Values) < 1 {
		SomethingWentWrong(c, "The selected mode did not contain a key")
	}

	var mode wargame.Mode

	for _, m := range c.Wargame.GameModes {
		if m.Name == data.Values[0] {
			mode = m
			break
		}
	}

	c.Wargame.Server.SetMode(&mode)
	log.Printf("[Discord] %s set mode too %s", c.User.GlobalName, mode.Name)
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Mode set too %s", mode.Name),
		},
	})
}

// Handles the map command
func MapHandler(c Context) {
	log.Println("[Discord] Map Command")
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
	var content = "Available maps for this mode:\n"

	for _, m := range c.Wargame.Server.Mode.MapList {
		//TODO embed image link to see the map
		content += fmt.Sprintf("- %s (%vv%v)\n", m.Name, m.Type, m.Type)
	}
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Maps",
					Description: content,
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
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
