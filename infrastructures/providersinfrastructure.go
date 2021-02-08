package infrastructures

import (
	"database/sql"
	"fmt"
	"time"

	"godrider/infrastructures/models"

	_ "modernc.org/sqlite"
)

type ProvidersInfrastructure struct {
	providerDb []models.Provider
	rows       int
	lastUpdate time.Time
}

var ProvidersDb = ProvidersInfrastructure{}

// GetAllProviders returns all registers from user table and its count as integer
func (infrastructure *ProvidersInfrastructure) GetAllProviders() ([]models.Provider, int, error) {
	if infrastructure.isDbUpdated() {
		return infrastructure.providerDb, infrastructure.rows, nil
	}

	db, _ := sql.Open("sqlite", "./db/godrider.db")
	defer db.Close()

	var count int
	if err := infrastructure.countRegisters(db, &count); err != nil {
		return nil, 0, err
	}

	rows, err := db.Query("SELECT * FROM provider;")
	defer rows.Close()
	if err != nil {
		return nil, 0, err
	}

	providerDb := make([]models.Provider, count)
	toUpdate := false
	for rows.Next() {
		var provider models.Provider
		err := rows.Scan(&provider.ID, &provider.Name, &provider.Contact)
		if err != nil {
			return nil, 0, err
		}

		providerDb = append(providerDb, provider)
		if !toUpdate {
			toUpdate = true
		}
	}

	if toUpdate {
		infrastructure.providerDb = providerDb
		infrastructure.rows = count
		infrastructure.lastUpdate = time.Now()
	}
	return infrastructure.providerDb, infrastructure.rows, nil
}

func (infrastructure *ProvidersInfrastructure) GetSingleProviderById(id int) (models.Provider, error) {
	if infrastructure.isDbUpdated() {
		for _, provider := range infrastructure.providerDb {
			if id == provider.ID {
				return provider, nil
			}
		}
		return models.Provider{}, fmt.Errorf("Provider with ID %d not found.", id)
	}

	db, _ := sql.Open("sqlite", "./db/godrider.db")
	defer db.Close()

	statement, _ := db.Prepare("SELECT * FROM provider WHERE ID = ?;")
	row := statement.QueryRow(id)

	var provider models.Provider
	err := row.Scan(&provider.ID, &provider.Name, &provider.Contact)
	if err != nil {
		return models.Provider{}, err
	}
	return provider, nil
}

func (infrastructure *ProvidersInfrastructure) isDbUpdated() bool {
	isProviderDbSet := infrastructure.rows > 0
	timeNow := time.Now().Unix()
	lastUpdate := infrastructure.lastUpdate.Unix()
	timeAfterLastUpdate := time.Now().Unix() - infrastructure.lastUpdate.Unix()
	fmt.Printf("Now (%d) - Last Update (%d) = %d", timeNow, lastUpdate, timeAfterLastUpdate)
	fmt.Println()
	return isProviderDbSet && (timeAfterLastUpdate > 3600)
}

func (infrastructure *ProvidersInfrastructure) countRegisters(db *sql.DB, count *int) error {
	return db.QueryRow("SELECT COUNT(*) FROM provider;").Scan(count)
}
