// Contains all the command handlers.

package discord

import "github.com/bwmarrin/discordgo"

func HelpHandler(c Context) {
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This is the help message!",
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
