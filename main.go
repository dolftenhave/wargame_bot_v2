package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"gopkg.in/yaml.v2"
	"wargame-bot/command"
	"wargame-bot/discord"
	"wargame-bot/rcon"
	sqlite3store "wargame-bot/store/sqlite3"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// The path to the config file
const CONFIGPATH = "conf.yaml"

type config struct {
	Discord struct{
		bot_token string `yaml:"bot_token"`
		owner_id string `yaml:"owner_id"`
	} `yaml:"discord"`
	Rcon struct {
		ip string `yaml:"ip"`
		port string `yaml:"port"`
		pword string `yaml:"pword"`
	} `yaml:"rcon"`
}

// Initialise the config struct
func initConf() (config, error){
	var conf config
	confFile, err := os.ReadFile(CONFIGPATH)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(confFile, &conf)
	if err != nil {
		return conf, err
	}

	log.Println("Log file loaded successfully.")
	return conf, nil
}

func main(){
	// Sets the log output to go the the log file.
	log.SetFlags(log.Ldate | log.Ltime)
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	conf, err := initConf()
	if err != nil {
		log.Fatal(err)
	}

	// Open the database
	db, err := sql.Open("sqlite3", "db/wargame.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := sqlite3store.NewSqlite3Store(db)

	rcon, err := rcon.NewRconClient(conf.rcon.ip, conf.rcon.port, conf.rcon.pword)
	if err != nil {
		log.Fatal(err)
	}

	registry := command.NewRegistry(store, rcon)
	command.RegisterAllCommands(registry)

	session, _ := discordgo.New(fmt.Sprintf("Bot %s", conf.discord.bot_token))
	adapter := &discord.DiscordAdapter{Registry: registry}
	session.AddHandler(adapter.HandleInteraction)
	session.Open()

	log.Println("Bot Ready.")
	<-make(chan struct{})
	fmt.Println("Done.")
}
