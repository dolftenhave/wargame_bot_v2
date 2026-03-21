package discord

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordResponder struct {
	session *discordgo.Session
	channelID string
}

// Create a new DiscordResponder
func NewDiscordResponder(session *discordgo.Session, channelID string) *DiscordResponder{
	return &DiscordResponder{
		session: session,
		channelID: channelID,
	}
}

func (r *DiscordResponder) SendMessage(ctx context.Context, message string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := r.session.ChannelMessageSend(r.channelID, message)
	if err != nil {
		return fmt.Errorf("discord send message: %w", err)
	}

	return nil
}

func (r *DiscordResponder) SendError(ctx context.Context, err error) error {
	return r.SendMessage(ctx, fmt.Sprintf("Error: %s", err.Error()))
}
