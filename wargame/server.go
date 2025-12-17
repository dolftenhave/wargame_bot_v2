package wargame

import (
	"fmt"
	"log"

	rcon "github.com/gorcon/rcon"
)

type (
	ServerState int

	RconConfig struct {
		Ip    string `yaml:"ip"`
		Port  string `yaml:"port"`
		Pword string `yaml:"pword"`
	}

	// A struct that holds data about the current state of the server.
	Server struct {
		Conn    rcon.Conn   // The rcon connection to the server.
		State   ServerState // The current state of the server.
		Mode    *Mode        // The current mode of the server.
		Players []Player    // The players currently connected.
	}
)

// Server States
const (
	Waiting ServerState = iota
	Debriefing
	Running
)


// Creates a connection with the server.
func (s *Server) CreateConn(conf *RconConfig) error {
	conn, err := rcon.Dial(fmt.Sprintf("%s:%s", conf.Ip, conf.Port), conf.Pword)
	if err != nil {
		return err
	}
	log.Println("[Server] Connection Created.")
	conn.Close()

	s.Conn = *conn
	return nil
}
