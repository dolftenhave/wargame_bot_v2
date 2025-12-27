package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type (
	Interaction func(Context)

	BotInteraction struct {
		interaction Interaction
		help        string
	}

	InteractionList map[string]BotInteraction

	BotInteractionHandler struct {
		interactions InteractionList
	}
)

func messageReciever(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("[Discord] Recieved: ID: %s, Type: %s", m.ID, m.Type)
}

// Sets up the interaction handler for an interaction.
func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Printf("[D-Interaction] Recieved: ID: %s, Type: %s", i.ID, i.Type)
	var sender string
	var user *discordgo.User

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		sender = i.ID
	default:
		log.Println("[Interaction] Wrong type")
		return
	}

	channel, err := s.State.Channel(i.ChannelID)
	if err != nil {
		log.Println("[Discord] Failed to get channelID")
		return
	}

	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		log.Println("[Discord] Failed to get guildID")
	}

	if i.User == nil {
		user = i.Member.User
	} else {
		user = i.User
	}

	log.Printf("[D-Interaction] Token: %s, Name: %s", i.Token, sender)

	sender = i.ApplicationCommandData().Name
	interaction, found := handler.Find(sender)
	if found != true {
		CommandNotImplemented(s, i)
		log.Printf("[D-Interaction] No handler for %s found", sender)
		return
	}

	context := NewContext(s, i, guild, channel, user, w)
	context.LogRecieved()

	c := *interaction
	c(*context)
}

// Sends back a message notifying the sender that the command has not been implemented yet
func CommandNotImplemented(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Sorry, this command has not yet been implemented.",
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}

func NewInteractionHandler() *BotInteractionHandler {
	return &BotInteractionHandler{make(InteractionList)}
}

func (handler BotInteractionHandler) Find(name string) (*Interaction, bool) {
	interaction, found := handler.interactions[name]

	return &interaction.interaction, found
}

// Register a new interaction
func (handler BotInteractionHandler) Register(name string, interaction Interaction, help string) {
	botInteractionStruct := new(BotInteraction)
	botInteractionStruct.interaction = interaction
	botInteractionStruct.help = help
	handler.interactions[name] = *botInteractionStruct
	if len(name) > 1 {
		handler.interactions[name[:1]] = *botInteractionStruct
	}
}
