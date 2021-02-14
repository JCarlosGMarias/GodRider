package infrastructures

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"godrider/infrastructures/models"
)

// ProviderInfrastructurer provides access to all registered webservices' providers
type ProviderInfrastructurer interface {
	// GetAllProviders should return all registers from provider table and its count as integer
	GetAllProviders() ([]models.Provider, int, error)
	// GetSingleProviderByID should return an unique provider model by its ID
	GetSingleProviderByID(id int) (models.Provider, error)
	// GetManyProvidersByIds should return a slice with provider's models by a slice of IDs
	GetManyProvidersByIds(ids []int) ([]models.Provider, error)
}

// ProviderInfrastructure is ProvidersInfrastructurer' implementation struct
type ProviderInfrastructure struct {
	providerDb []models.Provider
	rows       int
	lastUpdate time.Time
}

// ProviderDb is ProvidersInfrastructurer' implementation instance
var ProviderDb ProviderInfrastructurer = &ProviderInfrastructure{}

// GetAllProviders returns all registers from user table and its count as integer
func (istruct *ProviderInfrastructure) GetAllProviders() ([]models.Provider, int, error) {
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

// GetSingleProviderByID returns an unique provider model by its ID
func (istruct *ProviderInfrastructure) GetSingleProviderByID(id int) (models.Provider, error) {
	if istruct.isDbUpdated() {
		for _, provider := range istruct.providerDb {
			if id == provider.ID {
				return provider, nil
			}
		}
		return models.Provider{}, fmt.Errorf("Provider with ID %d not found", id)
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

// GetManyProvidersByIds returns a slice with provider's models by a slice of IDs
func (istruct *ProviderInfrastructure) GetManyProvidersByIds(ids []int) ([]models.Provider, error) {
	idsCount := len(ids)
	if idsCount <= 1 {
		provider, err := istruct.GetSingleProviderByID(ids[0])
		return []models.Provider{provider}, err
	}

	if istruct.isDbUpdated() {
		providers := make([]models.Provider, 0)
		for _, provider := range istruct.providerDb {
			for _, id := range ids {
				if id == provider.ID {
					providers = append(providers, provider)
				}
			}
		}

		if len(providers) > 0 {
			return providers, fmt.Errorf("No providers found")
		}
		return providers, nil
	}

	db, err := sql.Open("sqlite", "./db/godrider.db")
	if err != nil {
		return make([]models.Provider, 0), err
	}
	defer db.Close()

	queryPlaceHolders := make([]string, idsCount)
	queryArgs := make([]interface{}, idsCount)
	for i := 0; i < idsCount; i++ {
		queryPlaceHolders[i] = "?"
		queryArgs[i] = ids[i]
	}

	statement, err := db.Prepare(fmt.Sprintf("SELECT * FROM provider WHERE ID IN (%s);", strings.Join(queryPlaceHolders, ", ")))
	if err != nil {
		return make([]models.Provider, 0), err
	}
	defer statement.Close()

	rows, err := statement.Query(queryArgs...)
	if err != nil {
		return make([]models.Provider, 0), err
	}
	defer rows.Close()

	providers := make([]models.Provider, idsCount)
	for index, provider := range providers {
		if !rows.Next() {
			break
		}

		err := rows.Scan(&provider.ID, &provider.Name, &provider.Contact)
		if err != nil {
			return make([]models.Provider, 0), err
		}

		providers[index] = provider
	}
	return providers, nil
}

func (istruct *ProviderInfrastructure) isDbUpdated() bool {
	isProviderDbSet := istruct.rows > 0
	timeAfterLastUpdate := time.Now().Unix() - istruct.lastUpdate.Unix()
	return isProviderDbSet && (timeAfterLastUpdate <= 3600)
}

func (istruct *ProviderInfrastructure) countRegisters(db *sql.DB, count *int) error {
	return db.QueryRow("SELECT COUNT(*) FROM provider;").Scan(count)
}
