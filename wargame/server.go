package wargame

import (
	"fmt"
	"os"

	rcon "github.com/gorcon/rcon"
	"gopkg.in/yaml.v2"
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

// Loads the rcon connection string
func (c *RconConfig) GetConf(path string) error {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}
	return nil
}

// Creates a connection with the server.
func (s *Server) CreateConn(confPath string) error {
	var (
		err  error
		conf RconConfig
	)

	err = conf.GetConf(confPath)
	if err != nil {
		return err
	}

	conn, err := rcon.Dial(fmt.Sprintf("%s:%s", conf.Ip, conf.Port), conf.Pword)
	if err != nil {
		return err
	}
	conn.Close()

	s.Conn = *conn
	return nil
}
