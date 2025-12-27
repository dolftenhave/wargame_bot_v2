package discord

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"wargame-bot/wargame"

	"github.com/bwmarrin/discordgo"
)

const (
	//TODO Update to dynamic retrieving of guild and app id.
	APP_ID        = "1445521636533997600"
	GUILD_ID      = "1441939728760049809"
	BOT_TESTER_ID = "1451051992758878340"
)

var (
	handler *BotInteractionHandler
	w       *wargame.Wargame
)

// Configuration for the discord bot.
type (
	BotConfig struct {
		Token        string `yaml:"bot_token"`
		Owner_id     string `yaml:"owner_id"`
		Prefix       string `yaml:"interaction_prefix"`
		CommandsFile string `yaml:"commands"`
	}

	CommandStruct struct {
		Interactions []*discordgo.ApplicationCommand `yaml:"interarctions"`
	}
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

	w = wargameData
	//err = registerCommands(session)
	//if err != nil {
	//		return nil, err
	//	}

	handler = NewInteractionHandler()
	registerHandlers()
	session.AddHandler(interactionHandler)
	//session.AddHandler(messageReciever)

	if err != nil {
		return nil, err
	}
	return session, nil
}

func registerCommands(s *discordgo.Session) error {
	// Manually listing out commands since there are only a few

	log.Println("[Discord] Regisetring slash commands")

	s.ApplicationCommandCreate(APP_ID, GUILD_ID, MapCommand())
	s.ApplicationCommandCreate(APP_ID, GUILD_ID, ModeCommand())
	s.ApplicationCommandCreate(APP_ID, GUILD_ID, DeckCommand())
	s.ApplicationCommandEdit(APP_ID, GUILD_ID, "1451051630588858472", HelpCommand())
	return nil
}

// Loads all the commands in the commands file into an array
func loadCommandFile(filePath string) (error, *CommandStruct) {
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		return err, nil
	}

	var commStr = new(CommandStruct)
	err = json.Unmarshal(jsonFile, commStr)
	if err != nil {
		return err, nil
	}

	return nil, commStr
}

// Sets any commands that have not yet been registered.
func AdvertiseCapabilites(s *discordgo.Session) error {

	//s.ApplicationCommandCreate(APP_ID, GUILD_ID, HelpCommand())
	//s.ApplicationCommandEdit(APP_ID,GUILD_ID, "1451051630588858472", HelpCommand())

	commands, err := s.ApplicationCommands(APP_ID, GUILD_ID)
	if err != nil {
		return err
	}

	if len(commands) < 1 {
		log.Println("[Discord] No commands Registered")
		return nil
	}

	return nil
}

// Registers a list of interactions that will be used by the interaction handler.
func registerHandlers() {
	handler.Register("help", HelpHandler, "This lists all the available commands.")
	handler.Register("mode", ModeHandler, "Lets you set or change the game mode.")
	handler.Register("map", MapHandler, "Map stuff")
	handler.Register("set_mode", SetModeHandler, "Sets the mode")
}
