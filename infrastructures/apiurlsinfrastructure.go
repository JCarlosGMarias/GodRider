package infrastructures

import (
	"fmt"
	"time"

	"godrider/infrastructures/models"
)

// APIUrlsInfrastructurer provides an access to many API configurations, for instance, the API endpoints
type APIUrlsInfrastructurer interface {
	// GetAllUrls should return all registers from apiurls table and its count as integer
	GetAllUrls() ([]models.ApiUrl, int, error)
	// GetSingleURL should return an unique url based on its key
	GetSingleURL(key string) (models.ApiUrl, error)
}

// APIUrlsInfrastructure is APIUrlsInfrastructurer's implementation struct
type APIUrlsInfrastructure struct {
	tableName  string
	apiUrlsDb  []models.ApiUrl
	rows       int
	lastUpdate time.Time
}

// GetAllUrls returns all registers from apiurls table and its count as integer
func (istruct *APIUrlsInfrastructure) GetAllUrls() ([]models.ApiUrl, int, error) {
	if isDbUpdated(istruct.rows, istruct.lastUpdate) {
		return istruct.apiUrlsDb, istruct.rows, nil
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

	apiUrls := make([]models.ApiUrl, count)
	toUpdate := false
	for index, apiURL := range apiUrls {
		if !rows.Next() {
			break
		}

		if err := rows.Scan(&apiURL.Key, &apiURL.Url); err != nil {
			return nil, 0, err
		}

		apiUrls[index] = apiURL
		if !toUpdate {
			toUpdate = true
		}
	}

	if toUpdate {
		istruct.apiUrlsDb = apiUrls
		istruct.rows = count
		istruct.lastUpdate = time.Now()
	}
	return istruct.apiUrlsDb, istruct.rows, nil
}

// GetSingleURL returns an unique url based on its key
func (istruct *APIUrlsInfrastructure) GetSingleURL(key string) (models.ApiUrl, error) {
	if isDbUpdated(istruct.rows, istruct.lastUpdate) {
		for _, apiURL := range istruct.apiUrlsDb {
			if key == apiURL.Key {
				return apiURL, nil
			}
		}
		return models.ApiUrl{}, fmt.Errorf("API Url with key '%s' not found", key)
	}

	db, err := openDb()
	if err != nil {
		return models.ApiUrl{}, err
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s WHERE Key = ?;", istruct.tableName))
	if err != nil {
		return models.ApiUrl{}, err
	}
	defer statement.Close()

	row := statement.QueryRow(key)

	var apiURL models.ApiUrl
	err = row.Scan(&apiURL.Key, &apiURL.Url)
	if err == nil {
		return apiURL, nil
	}
	return models.ApiUrl{}, err
}

// TableName setter
func (i *APIUrlsInfrastructure) TableName(name string) {
	if i.tableName == "" {
		i.tableName = name
	}
}
