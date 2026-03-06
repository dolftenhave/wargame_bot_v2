package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"wargame-bot/domain"
)

// Scans a single row of the database into a new map struct
func scanMap(s scanner) (*domain.Map, error) {
	var m domain.Map
	err := s.Scan(
		&m.Code,
		&m.Name,
		&m.Image,
		&m.Kind,
		&m.Size,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Scan each row into a domain.Map struct
func scanMaps(rows *sql.Rows) ([]domain.Map, error) {
	defer rows.Close()

	var maps []domain.Map
	for rows.Next() {
		m, err := scanMap(rows)
		if err != nil {
			return nil, fmt.Errorf("scanning maps: %w", err)
		}
		maps = append(maps, *m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating maps: %w", err)
	}
	return maps, nil
}

// Returns a list of all the maps that contain the substring name
func (sqlite3 *Sqlite3Store) GetMapByName(ctx context.Context, name string) ([]domain.Map, error) {
	rows, err := sqlite3.db.QueryContext(ctx, "SELECT * FROM map WHERE name is like '%?%' LIMIT 25", name)
	if err != nil {
		return nil, err
	}
	return scanMaps(rows)

}

// Returns all the maps that contain the substring code
func (sqlite3 *Sqlite3Store) GetMapByCode(ctx context.Context, code string) ([]domain.Map, error) {
	rows, err := sqlite3.db.QueryContext(ctx, "SELECT * FROM map WHERE code IS LIKE '%?%' LIMIT 25", code)
	if err != nil {
		return nil, err
	}
	return scanMaps(rows)
}

// Returns all the maps that have a specific size
func (sqlite3 *Sqlite3Store) GetMapBySize(ctx context.Context, size int) ([]domain.Map, error) {
	rows, err := sqlite3.db.QueryContext(ctx, "SELECT * FROM map WHERE intended_size IS ? LIMIT 25", size)
	if err != nil {
		return nil, err
	}
	return scanMaps(rows)
}

// Returns all the maps that have a specific size
func (sqlite3 *Sqlite3Store) GetMapByKind(ctx context.Context, kind string) ([]domain.Map, error) {
	rows, err := sqlite3.db.QueryContext(ctx, "SELECT * FROM map WHERE kind IS ? LIMIT 25", kind)
	if err != nil {
		return nil, err
	}
	return scanMaps(rows)
}
