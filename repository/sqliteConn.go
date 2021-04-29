package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteInfo struct {
	driverName     string
	dataSourceName string
}

func InitDB(e *sqliteInfo) error {
	e.driverName = "sqlite3"
	e.dataSourceName = "./foo.db"

	db, err := sql.Open(e.driverName, e.dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	checkErr(err)

	rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table' AND name='exchange_rate';`)
	if err != nil {
		return fmt.Errorf("failed to query table: %v", err)
	}
	rows.Close()
	if !rows.Next() {
		sql_table := `
		CREATE TABLE IF NOT EXISTS exchange_rate(
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL COLLATE NOCASE,
			cashBuyingRate REAL NOT NULL,
			cashSellingRate REAL NOT NULL,
			signBuyingRate REAL NOT NULL,
			signSellingRate REAL NOT NULL,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		`
		db.Exec(sql_table)
	}

	return nil
}
