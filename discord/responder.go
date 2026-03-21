package discord

import (
	"fmt"
	"wargame-bot/command"

	"github.com/bwmarrin/discordgo"
)

type DiscordResponder struct {
	Session     *discordgo.Session
	Interaction *discordgo.InteractionCreate
}

func (d *DiscordResponder) SendResult(result command.CommandResult) error {
	switch result.Type {
	case command.ResultError:
		return d.Session.InteractionRespond(d.Interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("**Error:** %s", result.Message),
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
	case command.ResultSuccess:
		return d.Session.InteractionRespond(d.Interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("**Success!** %s", result.Message),
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
	case command.ResultList:
		var fields []*discordgo.MessageEmbedField
		for _, f := range result.Fields {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  f.Name,
				Value: f.Value,
			})
		}
		return d.Session.InteractionRespond(d.Interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:  result.Title,
							Fields: fields,
						},
					},
				},
			})
	case command.ResultMessage:
		return d.Session.InteractionRespond(d.Interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: result.Message,
				},
			})
	}

	return nil
}
