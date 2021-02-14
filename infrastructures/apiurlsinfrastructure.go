package infrastructures

import (
	"database/sql"
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
	apiUrlsDb  []models.ApiUrl
	rows       int
	lastUpdate time.Time
}

// GetAllUrls returns all registers from apiurls table and its count as integer
func (istruct *APIUrlsInfrastructure) GetAllUrls() ([]models.ApiUrl, int, error) {
	if istruct.isDbUpdated() {
		return istruct.apiUrlsDb, istruct.rows, nil
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

	rows, err := db.Query("SELECT * FROM apiurl;")
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
	if istruct.isDbUpdated() {
		for _, apiURL := range istruct.apiUrlsDb {
			if key == apiURL.Key {
				return apiURL, nil
			}
		}
		return models.ApiUrl{}, fmt.Errorf("API Url with key '%s' not found", key)
	}

	db, err := sql.Open("sqlite", "./db/godrider.db")
	if err != nil {
		return models.ApiUrl{}, err
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM apiurl WHERE Key = ?;")
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

func (istruct *APIUrlsInfrastructure) isDbUpdated() bool {
	isAPIUrlsDbSet := istruct.rows > 0
	timeAfterLastUpdate := time.Now().Unix() - istruct.lastUpdate.Unix()
	return isAPIUrlsDbSet && (timeAfterLastUpdate <= 3600)
}

func (istruct *APIUrlsInfrastructure) countRegisters(db *sql.DB, count *int) error {
	return db.QueryRow("SELECT COUNT(*) FROM apiurl;").Scan(count)
}
