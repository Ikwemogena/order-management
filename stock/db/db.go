package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func Init(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func (db *Database) InitTables() error {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS stock (
            item_id VARCHAR(255) PRIMARY KEY,
            quantity INTEGER NOT NULL
        )
    `)
    if err != nil {
        return fmt.Errorf("failed to create stock table: %v", err)
    }
    return nil
}