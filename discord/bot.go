package discord

import (
	"fmt"
	"log"
	"wargame-bot/wargame"

	"github.com/bwmarrin/discordgo"
)

const (
	//TODO Update to dynamic retrieving of guild and app id.
	APP_ID   = "1445521636533997600"
	GUILD_ID = "1441939728760049809"
)

// Starts a new instance of the discord bot.
func StartBot(conf BotConfig, wargameData *wargame.Wargame) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		log.Fatalf("[Discord] Error creating discordgo session.\n%s", err.Error())
	}
	usr, err := session.User("@me")
	if err != nil {
		log.Fatalf("[Discord] Error setting discordgo user.\n%s", err.Error())
	}

	var botId = usr.ID
	fmt.Println(botId)

	err = session.Open()
	if err != nil {
		log.Fatalf("[Discord] Error opening connection to the discord server.\n%s", err.Error())
	}

	SetCommands(session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// Sets any commands that have not yet been registered.
func SetCommands(s *discordgo.Session) error {

	commands, err := s.ApplicationCommands(APP_ID, "")
	if err != nil {
		return err
	}

	if len(commands) < 1{
		log.Println("[Discord] No commands Registered")
		return nil
	}

	for _, command := range commands{
		fmt.Printf("\t[Registered Command] Name: %s, Description: %s\n",command.Name, command.Description)
	}

	return nil
}
