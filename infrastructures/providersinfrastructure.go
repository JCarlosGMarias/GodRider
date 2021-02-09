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
func (istruct *ProvidersInfrastructure) GetAllProviders() ([]models.Provider, int, error) {
	if istruct.isDbUpdated() {
		return istruct.providerDb, istruct.rows, nil
	}

	db, err := sql.Open("sqlite", "./db/godrider.db")
	if err != nil {
		return nil, 0, err
	}
	defer db.Close()

	var count int
	if err := istruct.countRegisters(db, &count); err != nil {
		return nil, 0, err
	}

	rows, err := db.Query("SELECT * FROM provider;")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	providers := make([]models.Provider, count)
	toUpdate := false
	for index, provider := range providers {
		if !rows.Next() {
			break
		}

		err := rows.Scan(&provider.ID, &provider.Name, &provider.Contact)
		if err != nil {
			return nil, 0, err
		}

		providers[index] = provider
		if !toUpdate {
			toUpdate = true
		}
	}

	if toUpdate {
		istruct.providerDb = providers
		istruct.rows = count
		istruct.lastUpdate = time.Now()
	}
	return istruct.providerDb, istruct.rows, nil
}

func (istruct *ProvidersInfrastructure) GetSingleProviderById(id int) (models.Provider, error) {
	if istruct.isDbUpdated() {
		for _, provider := range istruct.providerDb {
			if id == provider.ID {
				return provider, nil
			}
		}
		return models.Provider{}, fmt.Errorf("Provider with ID %d not found.", id)
	}

	db, err := sql.Open("sqlite", "./db/godrider.db")
	if err != nil {
		return models.Provider{}, err
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM provider WHERE ID = ?;")
	if err != nil {
		return models.Provider{}, err
	}
	defer statement.Close()

	row := statement.QueryRow(id)

	var provider models.Provider
	err = row.Scan(&provider.ID, &provider.Name, &provider.Contact)
	if err == nil {
		return provider, nil
	}
	return models.Provider{}, err
}

func (istruct *ProvidersInfrastructure) isDbUpdated() bool {
	isProviderDbSet := istruct.rows > 0
	timeAfterLastUpdate := time.Now().Unix() - istruct.lastUpdate.Unix()
	return isProviderDbSet && (timeAfterLastUpdate <= 3600)
}

func (istruct *ProvidersInfrastructure) countRegisters(db *sql.DB, count *int) error {
	return db.QueryRow("SELECT COUNT(*) FROM provider;").Scan(count)
}
