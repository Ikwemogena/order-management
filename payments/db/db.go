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
       CREATE TABLE IF NOT EXISTS payments (
            id UUID PRIMARY KEY,
            order_id UUID NOT NULL,
            reference VARCHAR(255) UNIQUE NOT NULL,
            amount DECIMAL(10, 2) NOT NULL,
            currency VARCHAR(10) DEFAULT 'NGN',
            status VARCHAR(50) NOT NULL,
            payment_provider TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        return fmt.Errorf("failed to create payments table: %v", err)
    }
    return nil
}