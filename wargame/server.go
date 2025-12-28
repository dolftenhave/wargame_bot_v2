package wargame

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorcon/rcon"
)

var GameModes = map[int]string{
	1: "Destruction",
	2: "Siege",
	3: "Economy",
	4: "Conquete",
	5: "BreakThroughConquest",
}

type (
	ServerState int

	RconConfig struct {
		Ip    string `yaml:"ip"`
		Port  string `yaml:"port"`
		Pword string `yaml:"pword"`
	}

	// A struct that holds data about the current state of the server.
	Server struct {
		Conn       rcon.Conn // The rcon connection to the server.
		RconConfig RconConfig
		State      ServerState // The current state of the server.
		Mode       *Mode       // The current mode of the server.
		Players    []Player    // The players currently connected.
	}
)

// Server States
const (
	Waiting ServerState = iota
	Debriefing
	Running
)

func (s *Server) LogIncomeing(msg string) {
	log.Printf("[RCON] Recieved: %s\n", msg)
}

func (s *Server) LogOutgoint(msg string) {
	log.Printf("[RCON] Sending: %s\n", msg)
}

// Creates a connection with the server.
func (s *Server) CreateConn(conf *RconConfig) error {
	conn, err := rcon.Dial(fmt.Sprintf("%s:%s", conf.Ip, conf.Port), conf.Pword)
	if err != nil {
		return err
	}
	log.Println("[Server] Connection Created.")
	defer conn.Close()

	s.Conn = *conn
	return nil
}

func (s *Server) execCommand(command string, wg *sync.WaitGroup) {
	defer wg.Done()

	s.LogOutgoint(command)
	ack, err := s.Conn.Execute(command)
	if err != nil {
		log.Printf("[RCON] Err: %s\n", err.Error())
		return
	}
	s.LogIncomeing(ack)

}

// Executes a group of commands
func (s *Server) Send(commands []string) error {
	conn, err := rcon.Dial(fmt.Sprintf("%s:%s", s.RconConfig.Ip, s.RconConfig.Port), s.RconConfig.Pword)
	if err != nil {
		log.Printf("[RCON] err: %s", err.Error())
	}
	s.Conn = *conn

	/*	Cant run this fast over tcp.
		var wg sync.WaitGroup

		for _, c := range commands {
			wg.Add(1)
			go s.execCommand(c, &wg)
			time.Sleep(20 * time.Millisecond)
		}

		wg.Wait()
	*/

	for _, c := range commands {
		log.Printf("[RCON] Sending: %s", c)
		ack, err := s.Conn.Execute(c)
		if err != nil {
			break
		}
		log.Printf("[RCON] ACK: %s", ack)
	}

	conn.Close()

	if err != nil {
		return err
	}
	return nil
}

// Sets the mode of the server
func (s *Server) SetMode(mode *Mode) error {
	var commands []string
	var m string
	for _, code := range mode.MapList {
		m = fmt.Sprintf("%s%s", GameModes[mode.GameMode], code.Code)
	}

	commands = append(commands, s.setName(mode.Name))
	commands = append(commands, s.setMap(m))
	commands = append(commands, s.setNumPlayers(mode.TeamSize*2))
	commands = append(commands, s.setStartingPoints(mode.StartingPoints))
	commands = append(commands, s.setTimeLimit(mode.TimeLimit))
	commands = append(commands, s.setIncomeRate(mode.Income))
	commands = append(commands, s.setGameMode(mode.GameMode))
	commands = append(commands, s.setOposition(mode.Oposotion))
	commands = append(commands, s.setNations(mode.Nations))
	commands = append(commands, s.setEra(mode.Era))
	commands = append(commands, s.setTheme(mode.Theme))
	commands = append(commands, s.setScoreLimit(mode.ScoreLimit))
	if mode.AutoStart {
		commands = append(commands, s.setMinPlayers(mode.MinPlayers))
	} else {
		commands = append(commands, s.setMinPlayers(mode.MinPlayers+1))
	}
	commands = append(commands, s.setWarmupTime(mode.WarmupTime))
	commands = append(commands, s.setDeployTime(mode.DeployTime))
	commands = append(commands, s.setDebriefTime(mode.DebriefTime))
	commands = append(commands, s.setLoadingTime(mode.LoadingTime))

	err := s.Send(commands)

	if err != nil {
		return err
	}

	s.Mode = mode

	return nil
}

// Sets the map to the current map
func (s *Server) SetMap(m Map) error {
	var commands []string
	commands = append(commands, s.setMap(fmt.Sprintf("%s%s", GameModes[s.Mode.GameMode], m.Code)))
	err := s.Send(commands)
	return err
}

// Kick a player from the server using their player id or name.
func (s *Server) kick(id string) string {
	return fmt.Sprintf("kick %s", id)
}

// Kick a player from the server using their player id or name.
func (s *Server) ban(id string, hours int) string {
	return fmt.Sprintf("ban %s %v", id, hours)
}

// Unban a player from the server using their player id or name.
func (s *Server) unban(id string) string {
	return fmt.Sprintf("unban %s", id)
}

// Sets the name of the server
func (s *Server) setName(name string) string {
	return fmt.Sprintf("setsvar ServerName %s", name)
}

// Sets the number of players in the server
func (s *Server) setNumPlayers(n int) string {
	if n >= 2 && n <= 20 {
		return fmt.Sprintf("setsvar NbMaxPlayer %v", n)
	}
	return ""
}

// Sets the game type of the server. e.g. red vs red.
func (s *Server) setOposition(oposition int) string {
	return fmt.Sprintf("setsvar GameType %v", oposition)
}

// Sets the starting money.
func (s *Server) setStartingPoints(points int) string {
	return fmt.Sprintf("setsvar InitMoney %v", points)
}

// Sets the income rate of the server.
func (s *Server) setIncomeRate(money int) string {
	return fmt.Sprintf("setsvar IncomeRate %v", money)
}

// Sets the game duration in minutes.
func (s *Server) setTimeLimit(duration int) string {
	return fmt.Sprintf("setsvar TimeLimit %v", duration)
}

// Sets the map. The mapName is the id of the map. (yes poorly named).
func (s *Server) setMap(mapName string) string {
	return fmt.Sprintf("setsvar Map %s", mapName)
}

// Sets the game mode. e.g. distruction, conquest.
func (s *Server) setGameMode(mode int) string {
	return fmt.Sprintf("setsvar VictoryCond %v", mode)
}

// Sets the required number of players needed to start a game.
func (s *Server) setMinPlayers(n int) string {
	return fmt.Sprintf("setsvar NbMinPlayer %v", n)
}

// Sets the starting countdown in seconds. This begins once the required number of players has been reached.
func (s *Server) setStartCountdown(n int) string {
	return fmt.Sprintf("setsvar WarmupCountdown %v", n)
}

// Sets the maximum amount of time the server will wait until for players to connect before kicking those who are still connecting.
func (s *Server) setLoadingTime(n int) string {
	return fmt.Sprintf("setsvar LoadingTimeMax %v", n)
}

// Sets the nation constriant for the server.
func (s *Server) setNations(n int) string {
	return fmt.Sprintf("setsvar NationConstraint %v", n)
}

// Sets the nation constriant for the server.
func (s *Server) setEra(n int) string {
	return fmt.Sprintf("setsvar DateConstraint %v", n)
}

func (s *Server) setTheme(n int) string {
	return fmt.Sprintf("setsvar ThematicConstraint %v", n)
}

func (s *Server) setWarmupTime(n int) string {
	return fmt.Sprintf("setsvar WarmupCountdown %v", n)
}

func (s *Server) setLodingTime(n int) string {
	return fmt.Sprintf("setsvar LoadingTimeMax %v", n)
}

func (s *Server) setDeployTime(n int) string {
	return fmt.Sprintf("setsvar DeploimentTimeMax %v", n)
}

func (s *Server) setDebriefTime(n int) string {
	return fmt.Sprintf("setsvar DebriefingTimeMax %v", n)
}

func (s *Server) setScoreLimit(n int) string {
	return fmt.Sprintf("setsvar ScoreLimit %v", n)
}

func (s *Server) launch() string {
	return "launch"
}
func (s *Server) cancelLaunch() string {
	return "cancel_launch"
}

func (s *Server) getPlayers() string {
	return "display_all_clients"
}
func (s *Server) GetPlayers() string {
	conn, err := rcon.Dial(fmt.Sprintf("%s:%s", s.RconConfig.Ip, s.RconConfig.Port), s.RconConfig.Pword)
	if err != nil {
		log.Printf("[RCON] err: %s", err.Error())
	}
	s.Conn = *conn

	ack, err := s.Conn.Execute(s.getPlayers())

	conn.Close()

	return ack
}

func (s *Server) setDeckCode(playerID string, deckCode string) string {
	return fmt.Sprintf("setpvar %s PlayerDeckContent %s", playerID, deckCode)
}

func (s *Server) SetDeckCode(playerID string, deckCode string) string {

	conn, err := rcon.Dial(fmt.Sprintf("%s:%s", s.RconConfig.Ip, s.RconConfig.Port), s.RconConfig.Pword)
	if err != nil {
		log.Printf("[RCON] err: %s", err.Error())
	}
	s.Conn = *conn

	ack, err := s.Conn.Execute(s.setDeckCode(playerID, deckCode))

	conn.Close()

	return ack
}
