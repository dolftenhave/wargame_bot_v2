package sqlite3

// A private scanner interface to parse the database content directly into the type structs
type scanner interface {
	Scan(dest ...any) error
}
