package discord

import (
	"fmt"
	"log"
	"strconv"
	"wargame-bot/command"

	"github.com/bwmarrin/discordgo"
)

type DiscordAdapter struct {
	Registry *command.Registry
}

func (d *DiscordAdapter) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate){
	// Only accepts application commands
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()

	cmdName := data.Name
	var args []string

	if len(data.Options) > 0 {
		subCmd := data.Options[0]
		cmdName = fmt.Sprintf("%s_%s",cmdName, subCmd.Name)
		for _, opt := range subCmd.Options {
			args = append(args, opt.StringValue())
		}
	}

	var user *discordgo.User
	if i.Member != nil {
		user = i.Member.User
	} else {
		user = i.User
	}

	id, err := strconv.Atoi(user.ID)
	if err != nil {
		log.Printf("[Discord] Error converting player id into an int. id: %s", user.ID)
		return
	}

	caller := command.Caller{
		ID: id,
		Name: user.Username, 
		//TODO look up perms in the db
		Permission: command.PermEveryone, 
		Source: "discord",
	}

	log.Printf("[Discord] %s called %s with args %v", caller.Name, cmdName, args)

	result := d.Registry.Execute(cmdName, caller, args)
	responder := &DiscordResponder{
		Session: s,
		Interaction: i,
	}

	if err := responder.SendResult(result); err != nil {
		log.Printf("[Discord] Error sending response: %s", err)
	}
}
