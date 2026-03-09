package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"wargame-bot/domain"
)

// Scans a single row of the database into a new map struct
func scanMode(s scanner) (*domain.Mode, error) {
	var m domain.Mode
	err := s.Scan(
		&m.Name,
		&m.ServerName,
		&m.StartingPoints,
		&m.TimeLimit,
		&m.ScoreLimit,
		&m.IncomeRate,
		&m.GameMode,
		&m.Oposition,
		&m.Nation,
		&m.Era,
		&m.Theme,
		&m.TeamSize,
		&m.MinPlayers,
		&m.WarmupTime,
		&m.DeployTime,
		&m.DebriefTime,
		&m.LoadingTime,
		&m.AutoStart,
		&m.AutoRotate,
		&m.MapVote,
		&m.EnableCommands,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Scan each row into a domain.Map struct
func scanModes(rows *sql.Rows) ([]domain.Mode, error) {
	defer rows.Close()

	var modes []domain.Mode
	for rows.Next() {
		m, err := scanMode(rows)
		if err != nil {
			return nil, fmt.Errorf("scanning modes: %w", err)
		}
		modes = append(modes, *m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating modes: %w", err)
	}
	return modes, nil
}

// Returns a list of all the maps that contain the substring name
func (sqlite3 *Sqlite3Store) GetModeByName(ctx context.Context, name string) ([]domain.Mode, error) {
	rows, err := sqlite3.db.QueryContext(ctx, "SELECT * FROM mode WHERE name is like '%?%' LIMIT 25", name)
	if err != nil {
		return nil, err
	}
	return scanModes(rows)
}

// Return the default mode (the first mode in the mode table)
func (sqlite3 *Sqlite3Store) GetDefaultMode(ctx context.Context) (*domain.Mode, error) {
	rows, err := sqlite3.db.QueryContext(ctx, "SELECT * FROM mode ORDER BY id ASC LIMIT 1")
	if err != nil {
		return nil, err
	}
	return scanMode(rows)
}
