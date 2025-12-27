// Contains all the command handlers.

package discord

import "github.com/bwmarrin/discordgo"

type Option struct {
	Name        string
	Description string
	Required    bool
}

type CommandsHelp struct {
	Name        string
	Description string
	Options     []Option
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

}
