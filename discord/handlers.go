// Contains all the command handlers.

package discord

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"wargame-bot/wargame"

	"github.com/bwmarrin/discordgo"
)

// Send an error message to discord.
func SomethingWentWrong(c Context, message string) {
	content := "Oh no, something went wrong."
	if message != "" {
		content += fmt.Sprintf("\nMessage: %s", message)
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
	// Ack the response
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	options := c.Interaction.Interaction.ApplicationCommandData().Options

	var embeds []*discordgo.MessageEmbed

	log.Printf("Command ID %s", c.Interaction.ApplicationCommandData().ID)

	switch options[0].Value {
	case "deck":
		embeds = []*discordgo.MessageEmbed{
			{
				Title:       "Help",
				Description: "This lets you set or decode a deck.",
				Color:       0xCF574A,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "</deck set:1454259422908907784>",
						Value:  fmt.Sprintf("Will let you set the deck of a player on the server."),
						Inline: false,
					},
					{
						Name:   "</deck decode:1454259422908907784>",
						Value:  "Will tell you what the nation, spec and era a deck is.\n- `code`: The deck code.",
						Inline: false,
					},
				},
			},
		}
	case "mode":
		embeds = []*discordgo.MessageEmbed{
			{
				Title:       "Help",
				Description: "See or change the server mode.",
				Color:       0xCF574A,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "What is a mode?",
						Value:  "A mode is a way for me to group similar game experiences together. e.g. Tactical 10v10, or 2v2. The name of a mode will be the name of the server.",
						Inline: false,
					},
					{
						Name:   "</mode set:1454259421273395434>",
						Value:  "Lets you set the mode of the server to a different mode.",
						Inline: false,
					},
					{
						Name:   "</mode list:1454259421273395434>",
						Value:  "Displays all the available modes.",
						Inline: false,
					},
				},
			},
		}

	case "map":
		embeds = []*discordgo.MessageEmbed{
			{
				Title:       "Help",
				Description: "See available maps or change them.",
				Color:       0xCF574A,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "</map set:1454259419851264031>",
						Value:  "Lets you select a map from the available maps for the current mode.",
						Inline: false,
					},
					{
						Name:   "</map list:1454259419851264031>",
						Value:  "List all the available maps for the current mode. You can click the maps name to see an image of it!",
						Inline: false,
					},
					{
						Name:   "</map random:1454259419851264031>",
						Value:  "Randomly sets a map from the pool of available maps for the mode.",
						Inline: false,
					},
					{
						Name:   "</map vote:1454259419851264031>",
						Value:  "**NOTE:** Not implemented yet.\nVote on a random selection of up to 5 maps.",
						Inline: false,
					},
				},
			},
		}

	case "help":
		embeds = []*discordgo.MessageEmbed{
			{
				Title:       "Help",
				Description: "This bot lets you interact with the wargame server.\n\n**Important!** Please read the <#1452438914215313558> before using it.",
				Color:       0xCF574A,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Help Command",
						Value:  fmt.Sprintf("</help:1451051630588858472> - Will display this message.\n\n*Tip!* You can add the name of another command as an option to learn more about how to use it."),
						Inline: false,
					},
					{
						Name:   "Where can I use the bot?",
						Value:  "The bot will only work in <#1445051378304028682>.",
						Inline: false,
					},
				},
			},
		}

	default:
		embeds = []*discordgo.MessageEmbed{
			{
				Title:       "Help",
				Description: "This bot lets you interact with the wargame server.\n\n**Important!** Please read the <#1452438914215313558> before using it.",
				Color:       0xCF574A,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Help Command",
						Value:  fmt.Sprintf("</help:1451051630588858472> - Will display this message.\n\n*Tip!* You can add the name of another command as an option to learn more about how to use it."),
						Inline: false,
					},
					{
						Name:   "Where can I use the bot?",
						Value:  "The bot will only work in <#1445051378304028682>.",
						Inline: false,
					},
				},
			},
		}
	}

	c.Session.InteractionResponseEdit(c.Interaction.Interaction, &discordgo.WebhookEdit{
		Embeds: &embeds,
	})
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
			var def = false
			if m.Name == c.Wargame.Server.Mode.Name {
				def = true
			}
			mo = append(mo, discordgo.SelectMenuOption{
				Label:       m.Name,
				Value:       m.Name,
				Description: "",
				Default:     def,
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
	data := c.Interaction.MessageComponentData()
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

	// Delete the modal
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    fmt.Sprintf("You selected *%s*\nPlease wait while the setting are sent to the server...", mode.Name),
			Components: []discordgo.MessageComponent{},
		},
	})

	// ack the message
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	err := c.Wargame.Server.SetMode(&mode)
	if err != nil {
		log.Printf("[Discord] Error setting mode.\n%s", err.Error())
		SomethingWentWrong(c, "Error setting the mode, please check the logs")
		return
	}

	c.Wargame.Server.Mode = &mode

	var confirm = "Done!"
	_, err = c.Session.InteractionResponseEdit(c.Interaction.Interaction, &discordgo.WebhookEdit{
		Content: &confirm,
	})

	log.Printf("[Discord] %s set mode to %s", c.User.GlobalName, mode.Name)
	if err != nil {
		log.Printf("[Discord] Error: Failed to delete Set Mode Interaction.\n%s", err.Error())
	}

	_, err = c.Session.FollowupMessageCreate(c.Interaction.Interaction, false, &discordgo.WebhookParams{
		Content: fmt.Sprintf("<@%s> set the mode to %s", c.User.ID, mode.Name),
	})
	if err != nil {
		log.Printf("[Discord] Error: Setting mode\n%s", err.Error())
	}
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
	var content = "Click on the map to see an image.\n\n**Available maps:**\n"

	for _, m := range c.Wargame.Server.Mode.MapList {
		//TODO embed image link to see the map
		if m.Image != "" {
			content += fmt.Sprintf("- [%s (%vv%v)](%s)\n", m.Name, m.Type, m.Type, m.Image)
		} else {

			content += fmt.Sprintf("- %s (%vv%v)\n", m.Name, m.Type, m.Type)
		}
	}
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       c.Wargame.Server.Mode.Name,
					Description: content,
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}

func setMap(c Context) {
	options := c.Interaction.Interaction.ApplicationCommandData().Options
	options = options[0].Options

	if options == nil || len(options) < 1 || options[0].Name == "" {
		if len(c.Wargame.GameModes) < 1 {
			SomethingWentWrong(c, "There are no maps in this mode.")
			return
		}

		log.Printf("[Discord] %s is selecting a Map.", c.User.GlobalName)

		var mo []discordgo.SelectMenuOption

		for i, m := range c.Wargame.Server.Mode.MapList {
			mo = append(mo, discordgo.SelectMenuOption{
				Label:       fmt.Sprintf("%s (%vv%v)", m.Name, m.Type, m.Type),
				Value:       fmt.Sprintf("%v", i),
				Description: "",
			})
		}

		err := c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.SelectMenu{
								CustomID:    "select_map",
								Placeholder: "Please choose a map...",
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

	return
	//TODO add auto complete
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

func SelectMapHandler(c Context) {
	data := c.Interaction.MessageComponentData()
	if len(data.Values) < 1 {
		SomethingWentWrong(c, "The selected map did not contain a key")
	}
	key, err := strconv.Atoi(data.Values[0])
	if err != nil {
		SomethingWentWrong(c, fmt.Sprintf("The map key `%v` is not in the correct format (int)", data.Values[0]))
		return
	}

	if key >= len(c.Wargame.Server.Mode.MapList) {
		SomethingWentWrong(c, fmt.Sprintf("The key does not match a map in this modes maplist."))
		return
	}

	m := c.Wargame.Server.Mode.MapList[key]

	// Delete the modal
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    fmt.Sprintf("You selected *%s (%vv%v)*\nPlease wait while the setting are sent to the server...", m.Name, m.Type, m.Type),
			Components: []discordgo.MessageComponent{},
		},
	})

	// ack the message
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	err = c.Wargame.Server.SetMap(m)

	if err != nil {
		log.Printf("[Discord] Error setting map.\n%s", err.Error())
		SomethingWentWrong(c, "Error setting the map, please check the logs")
		return
	}
	var confirm = "Done!"
	_, err = c.Session.InteractionResponseEdit(c.Interaction.Interaction, &discordgo.WebhookEdit{
		Content: &confirm,
	})

	log.Printf("[Discord] %s set map to %s (%vv%v)", c.User.GlobalName, m.Name, m.Type, m.Type)
	if err != nil {
		log.Printf("[Discord] Error: Failed to delete Set Map Interaction.\n%s", err.Error())
	}

	_, err = c.Session.FollowupMessageCreate(c.Interaction.Interaction, false, &discordgo.WebhookParams{
		Content: fmt.Sprintf("<@%s> set the map to %s (%vv%v)", c.User.ID, m.Name, m.Type, m.Type),
	})
	if err != nil {
		log.Printf("[Discord] Error: Setting map\n%s", err.Error())
	}
}

func voteMap(c Context) {
	SomethingWentWrong(c, "")
}

func randomMap(c Context) {
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	var nMaps = len(c.Wargame.Server.Mode.MapList)
	if nMaps < 1 {
		SomethingWentWrong(c, "There are no maps set for this mode.")
	}

	var i = rand.Intn(nMaps)
	var m = c.Wargame.Server.Mode.MapList[i]

	content := fmt.Sprintf("Setting map too %s (%vv%v)", m.Name, m.Type, m.Type)
	c.Session.InteractionResponseEdit(c.Interaction.Interaction, &discordgo.WebhookEdit{
		Content: &content,
	})

	err := c.Wargame.Server.SetMap(m)
	if err != nil {
		content = fmt.Sprintf("There was an error setting the map... \n%s", err.Error())
	} else {
		content = fmt.Sprintf("<@%s> Set the map too %s (%vv%v)", c.User.ID, m.Name, m.Type, m.Type)
	}
	c.Session.InteractionResponseEdit(c.Interaction.Interaction, &discordgo.WebhookEdit{
		Content: &content,
	})
}

func DeckHandler(c Context) {
	log.Println("[Discord] Deck Command")
	options := c.Interaction.Interaction.ApplicationCommandData().Options
	if options == nil {
		SomethingWentWrong(c, "Handler Options is nil")
		return
	}
	switch options[0].Name {
	case "set":
		setDeck(c)
	case "decode":
		decodeDeck(c)
	default:
		SomethingWentWrong(c, fmt.Sprintf("Mode option '%s' is not yet implemented", options[0].Name))
	}
}

func SetDeck(c Context) {
	data := c.Interaction.MessageComponentData()
	if len(data.Values) < 1 {
		SomethingWentWrong(c, "There is no player id.")
	}
	split := strings.Split(data.Values[0], ",")
	key, err := strconv.Atoi(split[0])
	if err != nil {
		SomethingWentWrong(c, fmt.Sprintf("The player id `%v` is not in the correct format (int)", data.Values[0]))
		return
	}

	//TODO turn this into an auto complete so that players can just select a name that comes up.
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: fmt.Sprintf("deck_code_modal:%v:%s", key, split[1]),
			Title:    "Set Deck",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "deck_code",
							Label:       "Deck Code",
							Style:       discordgo.TextInputShort,
							Placeholder: "Please enter your deck code...",
							Required:    true,
						},
					},
				},
			},
		},
	})

	content := fmt.Sprintf("Setting deck for %s", split[1])
	c.Session.InteractionResponseEdit(c.Interaction.Interaction, &discordgo.WebhookEdit{
		Content:    &content,
		Components: &[]discordgo.MessageComponent{},
	})
}

func SetDeckCode(c Context) {
	data := c.Interaction.ModalSubmitData()

	parts := strings.Split(data.CustomID, ":")
	playerID := parts[1]
	playerName := parts[2]

	deckCode := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

	icons, err := wargame.DecodeDeck(deckCode, &c.Wargame.DeckCodeData)
	if err != nil {
		SomethingWentWrong(c, err.Error())
		return
	}

	iconString := ""

	for _, i := range icons {
		iconString += fmt.Sprintf("<:%s:%s>", strings.ToLower(i.Code), i.DiscID)
	}
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{

			Content: fmt.Sprintf("<@%s> set deck for player **%s**.\n%s\n```\n%s\n```", c.User.ID, playerName, iconString, deckCode),
		},
	})
	c.Wargame.Server.SetDeckCode(playerID, deckCode)

}

func setDeck(c Context) {
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	list, err := wargame.ToPlayerList(c.Wargame.Server.GetPlayers())

	if err != nil {
		SomethingWentWrong(c, err.Error())
	}

	if len(list) < 1 {
		content := "No Players."
		_, err = c.Session.InteractionResponseEdit(c.Interaction.Interaction, &discordgo.WebhookEdit{
			Content: &content,
		})
		return
	}

	var mo []discordgo.SelectMenuOption

	for _, m := range list {
		mo = append(mo, discordgo.SelectMenuOption{
			Label:       m.Name,
			Value:       fmt.Sprintf("%v,%s", m.ID, m.Name),
			Description: "",
		})
	}

	_, err = c.Session.InteractionResponseEdit(c.Interaction.Interaction, &discordgo.WebhookEdit{
		Components: &[]discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.SelectMenu{
						CustomID:    "set_deck",
						Placeholder: "Please choose your name...",
						Options:     mo,
					},
				},
			},
		},
	})

	if err != nil {
		SomethingWentWrong(c, err.Error())
	}
}

func decodeDeck(c Context) {
	c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	options := c.Interaction.Interaction.ApplicationCommandData().Options
	options = options[0].Options
	if len(options) < 1 {
		SomethingWentWrong(c, "Please add a deck code")
		return
	}
	var code = options[0].StringValue()

	if "" == code {
		SomethingWentWrong(c, "Please add a deck code.")
		return
	}

	if len(c.Wargame.DeckCodeData.Eras) < 1 {
		SomethingWentWrong(c, "Deck code data is not loaded.")
		return
	}
	icons, err := wargame.DecodeDeck(code, &c.Wargame.DeckCodeData)

	if err != nil {
		SomethingWentWrong(c, "Error decoding your deck")
		log.Printf("[Discord] Error: Decoding deck.\n%s", err.Error())
		return
	}

	if len(icons) < 1 {
		if len(c.Wargame.DeckCodeData.Eras) < 1 {
			SomethingWentWrong(c, "Deck code data is not loaded.")
			return
		}
		log.Printf("[Discord] The deck code was invalid")
		c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "There was a proplem decoding your deck",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	content := ""
	for _, icon := range icons {
		content += fmt.Sprintf("<:%s:%s>", strings.ToLower(icon.Code), icon.DiscID)
	}
	content += fmt.Sprintf("\n```\n%s\n```", code)
	_, err = c.Session.InteractionResponseEdit(c.Interaction.Interaction, &discordgo.WebhookEdit{
		Content: &content,
	})
	if err != nil {
		SomethingWentWrong(c, err.Error())
	}
}
