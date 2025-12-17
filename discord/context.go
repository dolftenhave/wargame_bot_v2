package discord

import (
	"wargame-bot/wargame"

	"github.com/bwmarrin/discordgo"
)

type (
	// Context of a message recieved form the discord channel.
	Context struct {
		Discord     *discordgo.Session
		Type        *discordgo.InteractionType
		Guild       *discordgo.Guild
		TextChannel *discordgo.Channel
		User        *discordgo.User
		Message     *discordgo.MessageCreate
		Wargame     *wargame.Wargame
		Args        []string

		Interactions BotInteractionHandler
		Prefix   string
	}
)

// Initialises a new Context variable
func NewContext(discord *discordgo.Session, guild *discordgo.Guild, textChannl *discordgo.Channel, user *discordgo.User, message *discordgo.MessageCreate, interactions BotInteractionHandler, prefix string, wargame *wargame.Wargame) *Context {
	context := new(Context)
	context.Discord = discord
	context.Guild = guild
	context.TextChannel = textChannl
	context.User = user
	context.Message = message

	context.Interactions = interactions
	context.Prefix = prefix
	context.Wargame = wargame
	return context
}
