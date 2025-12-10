package wargame

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type (
	Player struct {
		Name string
		ID   string
	}

	PlayerList map[string]Player
)

// Returns a player list created from a list returned by the server.
func ToPlayerList(players string) (PlayerList, error) {
	list := strings.Split(players, "\n")
	if len(list) < 2 {
		return nil, fmt.Errorf("No players")
	}

	l := list[1 : len(list)-1]
	var pl = make(PlayerList)

	for _, item := range l {
		id, name, found := strings.Cut(item, " ")
		if found != true {
			break
		}
		player := Player{
			Name: name,
			ID:   id,
		}
		pl[id] = player
	}
	return pl, nil
}

// Returns a list of players as a string.
func (pl PlayerList) ToString() string {
	var list = ""
	for _, p := range pl {
		list += fmt.Sprintf("- **%s** (%s)\n", p.Name, p.ID)
	}
	return list
}
