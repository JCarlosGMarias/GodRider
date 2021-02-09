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
func (istruct *UsersInfrastructure) GetAllUsers() ([]models.User, int, error) {
	if istruct.isDbUpdated() {
		return istruct.userDb, istruct.rows, nil
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

	rows, err := db.Query("SELECT * FROM user;")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	users := make([]models.User, count)
	toUpdate := false
	for index, user := range users {
		if !rows.Next() {
			break
		}

		err := rows.Scan(&user.ID, &user.Token, &user.User, &user.Password, &user.Name, &user.Surname, &user.Email, &user.Phone, &user.Level)
		if err != nil {
			return nil, 0, err
		}

		users[index] = user
		if !toUpdate {
			toUpdate = true
		}
	}

	if toUpdate {
		istruct.userDb = users
		istruct.rows = count
		istruct.lastUpdate = time.Now()
	}
	return istruct.userDb, istruct.rows, nil
}

func (istruct *UsersInfrastructure) GetSingleUserByUserAndPass(userName string, pass string) (models.User, error) {
	if istruct.isDbUpdated() {
		for _, user := range istruct.userDb {
			if userName == user.User && pass == user.Password {
				return user, nil
			}
		}
		return models.User{}, fmt.Errorf("User %s not found.", userName)
	}

	db, err := sql.Open("sqlite", "./db/godrider.db")
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM user WHERE user = ? AND password = ?;")
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()

	row := statement.QueryRow(userName, pass)

	var user models.User
	err = row.Scan(&user.ID, &user.Token, &user.User, &user.Password, &user.Name, &user.Surname, &user.Email, &user.Phone, &user.Level)
	if err == nil {
		return user, nil
	}
	return models.User{}, err
}

func (istruct *UsersInfrastructure) GetSingleUserByToken(token string) (models.User, error) {
	if istruct.isDbUpdated() {
		for _, user := range istruct.userDb {
			if token == user.Token.String {
				return user, nil
			}
		}
		return models.User{}, fmt.Errorf("User with given token not found.")
	}

	db, err := sql.Open("sqlite", "./db/godrider.db")
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM user WHERE token = ?;")
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()

	row := statement.QueryRow(token)

	var user models.User
	err = row.Scan(&user.ID, &user.Token, &user.User, &user.Password, &user.Name, &user.Surname, &user.Email, &user.Phone, &user.Level)
	if err == nil {
		return user, nil
	}
	return models.User{}, err
}

func (istruct *UsersInfrastructure) isDbUpdated() bool {
	isUserDbSet := istruct.rows > 0
	timeAfterLastUpdate := time.Now().Unix() - istruct.lastUpdate.Unix()
	return isUserDbSet && (timeAfterLastUpdate <= 3600)
}

func (istruct *UsersInfrastructure) countRegisters(db *sql.DB, count *int) error {
	return db.QueryRow("SELECT COUNT(*) FROM user;").Scan(count)
}
