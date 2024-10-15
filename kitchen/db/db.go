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
		-- Orders table to track each order sent to the kitchen
		CREATE TABLE IF NOT EXISTS kitchen_orders (
			order_id UUID PRIMARY KEY,          -- Unique ID for each order
			user_id VARCHAR(50) NOT NULL,         -- User placing the order
			status VARCHAR(50) DEFAULT 'pending', -- Current status (e.g., pending, accepted, preparing, completed)
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp when the order was created
		);

		-- Order Items table to track items associated with each kitchen order
		CREATE TABLE IF NOT EXISTS kitchen_order_items (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Unique ID for each item entry
			order_id UUID NOT NULL,                     -- Foreign key referring to kitchen_orders
			product_id VARCHAR(50) NOT NULL,               -- Product ID for the ordered item
			quantity INTEGER NOT NULL,                     -- Quantity of the product
			FOREIGN KEY (order_id) REFERENCES kitchen_orders(order_id) -- Enforces relationship to orders table
		);
	`)

	if err != nil {
		return fmt.Errorf("failed to create kitchen tables: %v", err)
	}
	
	return nil
}
