all: 
	sqlite3 wargame.db < db.sql
	sqlite3 wargame.db < maps.sql
	sqlite3 wargame.db < deck.sql
	sqlite3 wargame.db < modes.sql
	sqlite3 wargame.db < roles.sql

clean:
	rm wargame.db
