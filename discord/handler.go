package discord

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
