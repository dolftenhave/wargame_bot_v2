package discord

import (
)

type (
	Interaction    func(Context)
	BotInteraction struct {
		interaction Interaction
		help        string
	}

	InteractionList map[string]BotInteraction

	BotInteractionHandler struct {
		interactions InteractionList
	}

	// A configuration for the discord bot.
	BotConfig struct {
		Token    string `yaml:"bot_token"`
		Owner_id string `yaml:"owner_id"`
		Prefix   string `yaml:"interaction_prefix"`
	}
)

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
