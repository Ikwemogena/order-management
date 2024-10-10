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
		CREATE TABLE IF NOT EXISTS orders (
			order_id SERIAL PRIMARY KEY,
			user_id VARCHAR(50) NOT NULL,
			status VARCHAR(50) DEFAULT 'pending',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS order_items (
			id UUID PRIMARY KEY,
			order_id UUID NOT NULL,
			product_id VARCHAR(50) NOT NULL,
			FOREIGN KEY (order_id) REFERENCES orders(id)
		);
	`)
    if err != nil {
        return fmt.Errorf("failed to create stock table: %v", err)
    }
    return nil
}