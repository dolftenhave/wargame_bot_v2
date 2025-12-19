package discord

import (
	"log"
	"wargame-bot/wargame"

	"github.com/bwmarrin/discordgo"
)

type (
	// Context of a message recieved form the discord channel.
	Context struct {
		Session     *discordgo.Session
		Interaction *discordgo.InteractionCreate
		Guild       *discordgo.Guild
		Channel     *discordgo.Channel
		User        *discordgo.User

		Wargame     *wargame.Wargame
	}
)

// Initialises a new Context variable
func NewContext(session *discordgo.Session, interaction *discordgo.InteractionCreate, guild *discordgo.Guild, channel *discordgo.Channel, user *discordgo.User, wargame *wargame.Wargame) *Context {
	context := new(Context)

	context.Interaction = interaction
	context.Guild = guild
	context.Channel = channel
	context.User = user
	context.Wargame = wargame

	return context
}

// Logs the recieved request.
func (c Context) LogRecieved() {
	log.Printf("[Discord] %s:%s\n", c.Guild, c.User)
}
