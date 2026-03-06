package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"wargame-bot/domain"
)

// Scans a single row of the database into a new icon struct
func scanIcon(s scanner) (*domain.Icon, error) {
	var i domain.Icon
	err := s.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.EmoteID,
	)
	if err != nil {
		return nil, fmt.Errorf("scanning icon: %w", err)
	}
	return &i, nil
}

// Scan each row into a domain.Icon struct
func scanIcons(rows *sql.Rows) ([]domain.Icon, error) {
	defer rows.Close()

	var icons []domain.Icon
	for rows.Next() {
		i, err := scanIcon(rows)
		if err != nil {
			return nil, err
		}
		icons = append(icons, *i)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating icons: %w", err)
	}
	return icons, nil
}

// Gets an icon from the icon table table
func (sqlite3 *Sqlite3Store) getIconFromTable(ctx context.Context, table string, id int) (*domain.Icon, error) {
	row := sqlite3.db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id = ?", table), id)
	return scanIcon(row)
}

// Gets all icons from the icon table table
func (sqlite3 *Sqlite3Store) getAllIconsFromTable(ctx context.Context, table string) ([]domain.Icon, error) {
	rows, err := sqlite3.db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s", table))
	if err != nil {
		return nil, fmt.Errorf("querying %s icons: %w", table, err)
	}
	return scanIcons(rows)
}

// Returns a nation icon that matches the icon id, or an error if it does not exist
func (sqlite3 *Sqlite3Store) GetNationIcon(ctx context.Context, id int) (*domain.Icon, error) {
	return sqlite3.getIconFromTable(ctx, "nation", id)
}

// Returns all of the nation icons
func (sqlite3 *Sqlite3Store) GetAllNationIcons(ctx context.Context) ([]domain.Icon, error) {
	return sqlite3.getAllIconsFromTable(ctx, "nation")
}

// Returns a specialization icon that matches the icon id, or an error if it does not exist
func (sqlite3 *Sqlite3Store) GetSpecializationIcon(ctx context.Context, id int) (*domain.Icon, error) {
	return sqlite3.getIconFromTable(ctx, "specialization", id)
}

// Returns all of the specialization icons
func (sqlite3 *Sqlite3Store) GetAllSpecializationIcons(ctx context.Context) ([]domain.Icon, error) {
	return sqlite3.getAllIconsFromTable(ctx, "specialization")
}

// Returns an era icon that matches the icon id, or an error if it does not exist
func (sqlite3 *Sqlite3Store) GetEraIcon(ctx context.Context, id int) (*domain.Icon, error) {
	return sqlite3.getIconFromTable(ctx, "era", id)
}

// Returns all of the era icons
func (sqlite3 *Sqlite3Store) GetAllEraIcons(ctx context.Context) ([]domain.Icon, error) {
	return sqlite3.getAllIconsFromTable(ctx, "era")
}
