package db

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/mattn/go-sqlite3"
)

type (
	Mode struct {
		ID             int    `db:"id" sqltype:"INTEGER PRIMARY KEY"`
		ModeName       string `db:"mode_name" sqltype:"TEXT NOT NULL"`
		ServerName     string `db:"server_name" sqltype:"TEXT NOT NULL"`
		StartingPoints int    `db:"starting_points" sqltype:"INTEGER"`
		TimeLimit      int    `db:"time_limit" sqltype:"INTEGER"`
		ScoreLimit     int    `db:"score_limit" sqltype:"INTEGER"`
		IncomeRate     int    `db:"income_rate" sqltype:"INTEGER"`
		GameMode       int    `db:"game_mode" sqltype:"INTEGER"`
		Oposition      int    `db:"oposition" sqltype:"INTEGER"`
		Nations        int    `db:"nations" sqltype:"INTEGER"`
		Era            int    `db:"era" sqltype:"INTEGER"`
		Theme          int    `db:"theme" sqltype:"INTEGER"`
		TeamSize       int    `db:"team_size" sqltype:"INTEGER"`
		MinPlayers     int    `db:"min_players" sqltype:"INTEGER"`
		WarmupTime     int    `db:"warmup_time" sqltype:"INTEGER"`
		DeployTime     int    `db:"deploy_time" sqltype:"INTEGER"`
		DebriefTime    int    `db:"debrief_time" sqltype:"INTEGER"`
		LoadingTime    int    `db:"loading_time" sqltype:"INTEGER"`
		AutoStart      bool   `db:"auto_start" sqltype:"INTEGER"`
		AutoRotate     bool   `db:"auto_rotate" sqltype:"INTEGER"`
		MapVote        bool   `db:"map_vote" sqltype:"INTEGER"`
		EnableCommands bool   `db:"enable_commands" sqltype:"INTEGER"`
	}

	MapPool struct {
		ModeId          int `db:"mode_id" sqltype:"INTEGER"`
		Name            int `db:"name" sqltype:"TEXT NOT NULL"`
		Income_rate     int `db:"income_rate" sqltype:"INTEGER"`
		Starting_points int `db:"starting_points" sqltype:"INTEGER"`
		Score_limit     int `db:"score_limit" sqltype:"INTEGER"`
		Time_limit      int `db:"time_limit" sqltype:"INTEGER"`
	}

	Map struct {
		ID    int    `db:"id" sqltype:"TEXT PRIMARY KEY"`
		Name  string `db:"name" sqltype:"TEXT NOT NULL"`
		Image string `db:"image" sqltype:"TEXT"`
		Size  string `db:"size" sqltype:"TEXT NOT NULL"`
	}

	Nation struct {
		ID       int    `db:"id" sqltype:"INTEGER PRIMARY KEY"`
		Name     string `db:"name" sqltype:"TEXT NOT NULL"`
		Code     string `db:"code" sqltype:"TEXT NOT NULL"`
		Emode_id string `db:"emode_id" sqltype:"TEXT"`
	}

	Era struct {
		ID       int    `db:"id" sqltype:"INTEGER PRIMARY KEY"`
		Name     string `db:"name" sqltype:"TEXT NOT NULL"`
		Code     string `db:"code" sqltype:"TEXT NOT NULL"`
		Emode_id string `db:"emode_id" sqltype:"TEXT"`
	}

	Wargame_Player struct {
		ID   int    `db:"id" sqltype:"INTEGER PRIMARY KEY"`
		Name string `db:"name" sqltype:"TEXT NOT NULL"`
	}

	Discord_Player struct {
		ID   int    `db:"id" sqltype:"INTEGER PRIMARY KEY"`
		Name string `db:"name" sqltype:"TEXT NOT NULL"`
	}
)

// A helper function that makes it faster to define sqlite tables.
func CreateTableSQL(tableName string, model interface{}) string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var columns []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		colName := field.Tag.Get("db")
		colType := field.Tag.Get("sqltype")
		if colName == "" || colName == "-" {
			continue
		}

		if colType == "" {
			switch field.Type.Kind() {
			case reflect.Int, reflect.Int64:
				colType = "INTEGER"
			case reflect.Float64:
				colType = "float"
			case reflect.Bool:
				colType = "INTEGER"
			case reflect.String:
				colType = "TEXT"
			// TODO add time
			default:
				colType = "BLOB"
			}
		}
		columns = append(columns, fmt.Sprintf("%s %s", colName, colType))
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n\t%s\n);", tableName, strings.Join(columns, ",\n\t"))
}

// Migrate creates tables for all provided models
func Migrate(db *sql.DB, tables map[string]interface{}) error {
	for name, model := range tables {
		statement := CreateTableSQL(name, model)
		log.Printf("[DB] Executing:\n%s\n", statement)
		if _, err := db.Exec(statement); err != nil {
			return fmt.Errorf("[DB] Failed to create table %s: %w", name, err)
		}
	}
	return nil
}
