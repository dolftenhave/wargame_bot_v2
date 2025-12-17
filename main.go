package main

import (
	"fmt"
	"log"
	"os"
	"wargame-bot/discord"
	"wargame-bot/wargame"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

const CONFIGPATH = "conf.yaml"

type (
	// Contains the paths to the config files for the various components.
	Conf struct {
		DiscordConf  string `yaml:"discord_conf"`
		RconConf     string `yaml:"rcon_conf"`
		WargameModes string `yaml:"wargame_modes"`
		WargameMaps  string `yaml:"wargame_maps"`
	}
)

var (
	botId       string
	wargameData *wargame.Wargame
)

// Initalises the conf variale using the yaml file in CONFIGPATH
func (conf *Conf) initConf() error {
	yamlFile, err := os.ReadFile(CONFIGPATH)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return err
	}
	log.Println("Log file loaded sucessfully.")
	return nil
}

func main() {
	var (
		conf Conf
		err  error

		session *discordgo.Session
	)

	log.SetFlags(log.Ldate | log.Ltime)

	// Initialises the bot and reads the relevent config files. Stopping if there are any errors.
	err = conf.initConf()
	if err != nil {
		log.Fatalf("Error loading the the main conf.yaml.\n%s", err.Error())
	}

	wargameData = new(wargame.Wargame)

	err = wargameData.Maps.ReadConfig(conf.WargameMaps)
	if err != nil {
		log.Fatalf("Error loading maps.\n%s", err.Error())
	}

	err = wargameData.GameModes.ReadConfig(conf.WargameModes, &wargameData.Maps)
	if err != nil {
		log.Fatalf("Error loading maps.\n%s", err.Error())
	}

	err = wargameData.Server.CreateConn(conf.RconConf)
	if err != nil {
		log.Fatalf("Error creating a connection to the wargame server\n%s", err.Error())
	}

	var bConf = new(discord.BotConfig)
	err = bConf.ReadConfig(conf.DiscordConf)
	if err != nil {
		log.Fatalf("Error Creating a discord session.\n%s", err.Error())
	}

	session, err = discordgo.New("Bot " + bConf.Token)
	if err != nil {
		log.Fatalf("Error creating discordgo session.\n%s", err.Error())
	}
	usr, err := session.User("@me")
	if err != nil {
		log.Fatalf("Error setting discordgo user.\n%s", err.Error())
	}

	botId = usr.ID

	err = session.Open()
	if err != nil {
		log.Fatalf("Error opening connection to the discord server.\n%s", err.Error())
	}

	log.Println("Bot Ready")
	//<-make(chan struct{})
	fmt.Println("Done.")
}
