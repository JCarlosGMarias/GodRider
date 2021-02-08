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

	rows, err := db.Query("SELECT * FROM user;")
	defer rows.Close()
	if err != nil {
		return nil, 0, err
	}

	userDb := make([]models.User, count)
	toUpdate := false
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
		return models.User{}, fmt.Errorf("User %s not found.", userName)
	}

	db, _ := sql.Open("sqlite", "./db/godrider.db")
	defer db.Close()

	statement, _ := db.Prepare("SELECT * FROM user WHERE user = ? AND password = ?;")
	row := statement.QueryRow(userName, pass)
	defer statement.Close()

	var user models.User
	err := row.Scan(&user.ID, &user.Token, &user.User, &user.Password, &user.Name, &user.Surname, &user.Email, &user.Phone, &user.Level)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (infrastructure *UsersInfrastructure) GetSingleUserByToken(token string) (models.User, error) {
	if infrastructure.isDbUpdated() {
		for _, user := range infrastructure.userDb {
			if token == user.Token.String {
				return user, nil
			}
		}
		return models.User{}, fmt.Errorf("User with given token not found.")
	}

	db, _ := sql.Open("sqlite", "./db/godrider.db")
	defer db.Close()

	statement, _ := db.Prepare("SELECT * FROM user WHERE token = ?;")
	row := statement.QueryRow(token)
	defer statement.Close()

	var user models.User
	err := row.Scan(&user.ID, &user.Token, &user.User, &user.Password, &user.Name, &user.Surname, &user.Email, &user.Phone, &user.Level)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (infrastructure *UsersInfrastructure) isDbUpdated() bool {
	isUserDbSet := infrastructure.rows > 0
	timeNow := time.Now().Unix()
	lastUpdate := infrastructure.lastUpdate.Unix()
	timeAfterLastUpdate := time.Now().Unix() - infrastructure.lastUpdate.Unix()
	fmt.Printf("Now (%d) - Last Update (%d) = %d\n", timeNow, lastUpdate, timeAfterLastUpdate)
	return isUserDbSet && (timeAfterLastUpdate <= 3600)
}

func (infrastructure *UsersInfrastructure) countRegisters(db *sql.DB, count *int) error {
	return db.QueryRow("SELECT COUNT(*) FROM user;").Scan(count)
}
