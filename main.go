package main

import (
	"fmt"
	"log"
	"os"
	"wargame-bot/discord"
	"wargame-bot/wargame"

	"gopkg.in/yaml.v2"
)

const CONFIGPATH = "conf.yaml"

type (
	// Contains the paths to the config files for the various components.
	Conf struct {
		Discord discord.BotConfig  `yaml:"discord"`
		Rcon    wargame.RconConfig `yaml:"rcon"`
		Wargame struct {
			Modes string `yaml:"modes"`
			Maps  string `yaml:"maps"`
			Deck  string `yaml:"deck"`
		} `yaml:"wargame"`
	}
)

var (
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
	)

	log.SetFlags(log.Ldate | log.Ltime)

	// Initialises the bot and reads the relevent config files. Stopping if there are any errors.
	err = conf.initConf()
	if err != nil {
		log.Fatalf("Error loading the the main conf.yaml.\n%s", err.Error())
	}

	wargameData, err = wargame.NewWargame(conf.Wargame.Modes, conf.Wargame.Maps, conf.Rcon, conf.Wargame.Deck)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	_, err = discord.StartBot(conf.Discord, wargameData)
	if err != nil {
		log.Fatalf("Error starting discord bot.\n%s", err.Error())
	}

	log.Println("Bot Ready")
	<-make(chan struct{})
	fmt.Println("Done.")
}
