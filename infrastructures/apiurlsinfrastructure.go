package infrastructures

import (
	"database/sql"
	"fmt"
	"time"

	"godrider/infrastructures/models"

	_ "modernc.org/sqlite"
)

type ApiUrlsInfrastructure struct {
	apiUrlsDb  []models.ApiUrl
	rows       int
	lastUpdate time.Time
}

var ApiUrlsDb = ApiUrlsInfrastructure{}

// GetAllUrls returns all registers from apiurls table and its count as integer
func (istruct *ApiUrlsInfrastructure) GetAllUrls() ([]models.ApiUrl, int, error) {
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

	rows, err := db.Query("SELECT * FROM apiurls;")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	apiUrls := make([]models.ApiUrl, count)
	toUpdate := false
	for index, apiUrl := range apiUrls {
		if !rows.Next() {
			break
		}

		if err := rows.Scan(&apiUrl.Key, &apiUrl.Url); err != nil {
			return nil, 0, err
		}

		apiUrls[index] = apiUrl
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

func (istruct *ApiUrlsInfrastructure) GetSingleUrl(key string) (models.ApiUrl, error) {
	if istruct.isDbUpdated() {
		for _, apiUrl := range istruct.apiUrlsDb {
			if key == apiUrl.Key {
				return apiUrl, nil
			}
		}
		return models.ApiUrl{}, fmt.Errorf("API Url with key '%s' not found.", key)
	}

	db, err := sql.Open("sqlite", "./db/godrider.db")
	if err != nil {
		return models.ApiUrl{}, err
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM apiurls WHERE Key = ?;")
	if err != nil {
		return models.ApiUrl{}, err
	}
	defer statement.Close()

	row := statement.QueryRow(key)

	var apiUrl models.ApiUrl
	err = row.Scan(&apiUrl.Key, &apiUrl.Url)
	if err == nil {
		return apiUrl, nil
	}
	return models.ApiUrl{}, err
}

func (istruct *ApiUrlsInfrastructure) isDbUpdated() bool {
	isApiUrlsDbSet := istruct.rows > 0
	timeAfterLastUpdate := time.Now().Unix() - istruct.lastUpdate.Unix()
	return isApiUrlsDbSet && (timeAfterLastUpdate <= 3600)
}

func (istruct *ApiUrlsInfrastructure) countRegisters(db *sql.DB, count *int) error {
	return db.QueryRow("SELECT COUNT(*) FROM apiurls;").Scan(count)
}
