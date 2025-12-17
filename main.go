package main

import (
	"fmt"
	"log"
	"os"
	"wargame-bot/wargame"

	//"github.com/bwmarrin/discordgo"
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
	maps        *wargame.MapList
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
		conf  Conf
		err   error
		modes wargame.ModeList

		//server  wargame.Server
		//discord discordgo.Session
	)

	log.SetFlags(log.Ldate | log.Ltime)

	// Initialises the bot and reads the relevent config files. Stopping if there are any errors.
	err = conf.initConf()
	if err != nil {
		log.Fatalf("Error loading the the main conf.yaml.\n%s", err.Error())
	}

	maps = new(wargame.MapList)
	err = maps.ReadConfig(conf.WargameMaps)
	if err != nil {
		log.Fatalf("Error loading maps.\n%s", err.Error())
	}

	err = modes.ReadConfig(conf.WargameModes, maps)
	if err != nil {
		log.Fatalf("Error loading maps.\n%s", err.Error())
	}

	//err = server.CreateConn(conf.RconConf)
	//if err != nil {
	//	log.Fatalf("Error creating a connection to the wargame server\n%s", err.Error())
	//}

	//err = discord.CreateDiscord(conf.DiscordConf)
	//if err != nil {
	//	log.Fatalf("Error Creating a discord session.\n%s", err.Error())
	//}

	modes[0].PrintMaps()
	modes.WriteConfig()
	//<-make(chan struct{})
	fmt.Println("Done.")
}
