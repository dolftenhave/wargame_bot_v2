package discord

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	Command    func(Context)
	BotCommand struct {
		command Command
		help    string
	}

	CommandList map[string]BotCommand

	BotCommandHandler struct {
		commands CommandList
	}

	// A configuration for the discord bot.
	BotConfig struct {
		Token    string `yaml:"bot_token"`
		Owner_id string `yaml:"owner_id"`
		Prefix   string `yaml:"command_prefix"`
	}
)

func (bConf *BotConfig) ReadConfig(filePath string) (error){
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, bConf)
	if err != nil {
		return fmt.Errorf("Unmarshal: %s\n", err)
	}

	if len(bConf.Token) == 0 {
		return fmt.Errorf("bot_token lenth is 0.\n")
	}
	return nil
}

func NewCommandHandler() *BotCommandHandler {
	return &BotCommandHandler{make(CommandList)}
}

func (handler BotCommandHandler) Find(name string) (*Command, bool) {
	command, found := handler.commands[name]

	return &command.command, found
}

// Register a new command
func (handler BotCommandHandler) Register(name string, command Command, help string){
	botCommandStruct := new(BotCommand)
	botCommandStruct.command = command
	botCommandStruct.help = help
	handler.commands[name] = *botCommandStruct
	if len(name) > 1 {
		handler.commands[name[:1]] = *botCommandStruct
	}
}
