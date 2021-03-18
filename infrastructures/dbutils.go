package infrastructures

import (
	"database/sql"
	"fmt"
	"time"
)

func openDb() (*sql.DB, error) {
	return sql.Open("sqlite", "./db/godrider.db")
}

func isDbUpdated(rows int, lastUpdate time.Time) bool {
	isDbSet := rows > 0
	timeAfterLastUpdate := time.Now().Unix() - lastUpdate.Unix()
	return isDbSet && (timeAfterLastUpdate <= 3600)
}

func countRegisters(db *sql.DB, tableName string, count *int) error {
	return db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s;", tableName)).Scan(count)
}
