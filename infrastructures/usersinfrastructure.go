package infrastructures

import (
	"database/sql"
	"fmt"
	"time"

	"godrider/infrastructures/models"

	_ "modernc.org/sqlite"
)

type UsersInfrastructure struct {
	userDb     []models.User
	rows       int
	lastUpdate time.Time
}

var UsersDb = UsersInfrastructure{}

// GetAllUsers returns all registers from user table and its count as integer
func (infrastructure *UsersInfrastructure) GetAllUsers() ([]models.User, int, error) {
	if infrastructure.isDbUpdated() {
		return infrastructure.userDb, infrastructure.rows, nil
	}

	db, _ := sql.Open("sqlite", "./db/godrider.db")
	defer db.Close()

	var count int
	if err := infrastructure.countRegisters(db, &count); err != nil {
		return nil, 0, err
	}

	userDb := make([]models.User, count)
	toUpdate := false
	rows, _ := db.Query("SELECT * FROM user;")
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Token, &user.User, &user.Password, &user.Name, &user.Surname, &user.Email, &user.Phone, &user.Level)
		if err != nil {
			return nil, 0, err
		}

		userDb = append(userDb, user)
		if !toUpdate {
			toUpdate = true
		}
	}

	if toUpdate {
		infrastructure.userDb = userDb
		infrastructure.rows = count
		infrastructure.lastUpdate = time.Now()
	}
	return infrastructure.userDb, infrastructure.rows, nil
}

func (infrastructure *UsersInfrastructure) GetSingleUserByUserAndPass(userName string, pass string) (models.User, error) {
	if infrastructure.isDbUpdated() {
		for _, user := range infrastructure.userDb {
			if userName == user.User && pass == user.Password {
				return user, nil
			}
		}
		return models.User{}, nil
	} else {
		db, _ := sql.Open("sqlite", "./db/godrider.db")
		defer db.Close()

		query := fmt.Sprintf("SELECT * FROM user WHERE user = '%s' AND password = '%s';", userName, pass)
		statement, _ := db.Prepare(query)
		row := statement.QueryRow()

		var user models.User
		err := row.Scan(&user.ID, &user.Token, &user.User, &user.Password, &user.Name, &user.Surname, &user.Email, &user.Phone, &user.Level)
		if err != nil {
			return models.User{}, err
		}
		return user, nil
	}
}

func (infrastructure *UsersInfrastructure) isDbUpdated() bool {
	isUserDbSet := infrastructure.userDb != nil
	timeAfterLastUpdate := time.Now().Unix() - infrastructure.lastUpdate.Unix()
	return isUserDbSet && (timeAfterLastUpdate > 3600)
}

func (infrastructure *UsersInfrastructure) countRegisters(db *sql.DB, count *int) error {
	return db.QueryRow("SELECT COUNT(*) FROM user;").Scan(count)
}
