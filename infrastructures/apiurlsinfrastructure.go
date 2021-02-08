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
func (infrastructure *ApiUrlsInfrastructure) GetAllUrls() ([]models.ApiUrl, int, error) {
	if infrastructure.isDbUpdated() {
		return infrastructure.apiUrlsDb, infrastructure.rows, nil
	}

	db, _ := sql.Open("sqlite", "./db/godrider.db")
	defer db.Close()

	var count int
	if err := infrastructure.countRegisters(db, &count); err != nil {
		return nil, 0, err
	}

	rows, err := db.Query("SELECT * FROM apiurls;")
	defer rows.Close()

	if err != nil {
		return nil, 0, err
	}

	apiUrlDb := make([]models.ApiUrl, count)
	toUpdate := false
	for rows.Next() {
		var apiUrl models.ApiUrl
		if err := rows.Scan(&apiUrl.Key, &apiUrl.Url); err != nil {
			return nil, 0, err
		}

		apiUrlDb = append(apiUrlDb, apiUrl)
		if !toUpdate {
			toUpdate = true
		}
	}

	if toUpdate {
		infrastructure.apiUrlsDb = apiUrlDb
		infrastructure.rows = count
		infrastructure.lastUpdate = time.Now()
	}
	return infrastructure.apiUrlsDb, infrastructure.rows, nil
}

func (infrastructure *ApiUrlsInfrastructure) GetSingleUrl(key string) (models.ApiUrl, error) {
	if infrastructure.isDbUpdated() {
		for _, apiUrl := range infrastructure.apiUrlsDb {
			if key == apiUrl.Key {
				return apiUrl, nil
			}
		}
		return models.ApiUrl{}, fmt.Errorf("API Url with key '%s' not found.", key)
	}

	db, _ := sql.Open("sqlite", "./db/godrider.db")
	defer db.Close()

	statement, _ := db.Prepare("SELECT * FROM apiurls WHERE Key = ?;")
	row := statement.QueryRow(key)

	var apiUrl models.ApiUrl
	if err := row.Scan(&apiUrl.Key, &apiUrl.Url); err != nil {
		return models.ApiUrl{}, err
	}

	return apiUrl, nil
}

func (infrastructure *ApiUrlsInfrastructure) isDbUpdated() bool {
	isApiUrlsDbSet := infrastructure.rows > 0
	timeNow := time.Now().Unix()
	lastUpdate := infrastructure.lastUpdate.Unix()
	timeAfterLastUpdate := time.Now().Unix() - infrastructure.lastUpdate.Unix()
	fmt.Printf("Now (%d) - Last Update (%d) = %d\n", timeNow, lastUpdate, timeAfterLastUpdate)
	return isApiUrlsDbSet && (timeAfterLastUpdate <= 3600)
}

func (infrastructure *ApiUrlsInfrastructure) countRegisters(db *sql.DB, count *int) error {
	return db.QueryRow("SELECT COUNT(*) FROM apiurls;").Scan(count)
}
