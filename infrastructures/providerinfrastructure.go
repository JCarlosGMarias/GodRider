package infrastructures

import (
	"fmt"
	"strings"
	"time"

	"godrider/infrastructures/models"
)

// ProviderInfrastructurer provides access to all registered webservices' providers
type ProviderInfrastructurer interface {
	// GetAll should return all registers from provider table and its count as integer
	GetAll() ([]models.Provider, int, error)
	// GetSingleByID should return an unique provider model by its ID
	GetSingleByID(id int) (models.Provider, error)
	// GetManyByIds should return a slice with provider's models by a slice of IDs
	GetManyByIds(ids []int) ([]models.Provider, error)
}

// ProviderInfrastructure is ProvidersInfrastructurer' implementation struct
type ProviderInfrastructure struct {
	tableName  string
	providerDb []models.Provider
	rows       int
	lastUpdate time.Time
}

// GetAll returns all registers from user table and its count as integer
func (istruct *ProviderInfrastructure) GetAll() ([]models.Provider, int, error) {
	if isDbUpdated(istruct.rows, istruct.lastUpdate) {
		return istruct.providerDb, istruct.rows, nil
	}

	db, err := openDb()
	if err != nil {
		return nil, 0, err
	}
	defer db.Close()

	var count int
	if err := countRegisters(db, istruct.tableName, &count); err != nil {
		return nil, 0, err
	}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s;", istruct.tableName))
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

// GetSingleByID returns an unique provider model by its ID
func (istruct *ProviderInfrastructure) GetSingleByID(id int) (models.Provider, error) {
	if isDbUpdated(istruct.rows, istruct.lastUpdate) {
		for _, provider := range istruct.providerDb {
			if id == provider.ID {
				return provider, nil
			}
		}
		return models.Provider{}, fmt.Errorf("Provider with ID %d not found", id)
	}

	db, err := openDb()
	if err != nil {
		return models.Provider{}, err
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s WHERE ID = ?;", istruct.tableName))
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

// GetManyByIds returns a slice with provider's models by a slice of IDs
func (istruct *ProviderInfrastructure) GetManyByIds(ids []int) ([]models.Provider, error) {
	idsCount := len(ids)
	if idsCount <= 1 {
		provider, err := istruct.GetSingleByID(ids[0])
		return []models.Provider{provider}, err
	}

	if isDbUpdated(istruct.rows, istruct.lastUpdate) {
		providers := make([]models.Provider, idsCount)
		for _, provider := range istruct.providerDb {
			for i, id := range ids {
				if id == provider.ID {
					providers[i] = provider
				}
			}
		}

		if len(providers) > 0 {
			return providers, fmt.Errorf("No providers found")
		}
		return providers, nil
	}

	db, err := openDb()
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

	statement, err := db.Prepare(fmt.Sprintf("SELECT * FROM "+istruct.tableName+" WHERE ID IN (%s);", strings.Join(queryPlaceHolders, ", ")))
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

		if err := rows.Scan(&provider.ID, &provider.Name, &provider.Contact); err != nil {
			return make([]models.Provider, 0), err
		}

		providers[index] = provider
	}
	return providers, nil
}

// TableName setter
func (i *ProviderInfrastructure) TableName(name string) {
	if i.tableName == "" {
		i.tableName = name
	}
}
