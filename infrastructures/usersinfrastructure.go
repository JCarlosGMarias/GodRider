package infrastructures

import (
	"fmt"
	"time"

	"godrider/infrastructures/models"
)

// UserInfrastructurer provides an access to many users related db interactions
type UserInfrastructurer interface {
	// GetAll should return all registers from user table and its count as integer
	GetAll() ([]models.User, int, error)
	// GetSingleByUserAndPass should return an unique user model by its user and password
	GetSingleByUserAndPass(userName string, pass string) (models.User, error)
	// GetSingleByToken should return an unique user model by its token
	GetSingleByToken(token string) (models.User, error)
}

// UserInfrastructure is UserInfrastructurer's implementation struct
type UserInfrastructure struct {
	tableName  string
	userDb     []models.User
	rows       int
	lastUpdate time.Time
}

// GetAll returns all registers from user table and its count as integer
func (istruct *UserInfrastructure) GetAll() ([]models.User, int, error) {
	if isDbUpdated(istruct.rows, istruct.lastUpdate) {
		return istruct.userDb, istruct.rows, nil
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

// GetSingleByUserAndPass returns an unique user model by its user and password
func (istruct *UserInfrastructure) GetSingleByUserAndPass(userName string, pass string) (models.User, error) {
	if isDbUpdated(istruct.rows, istruct.lastUpdate) {
		for _, user := range istruct.userDb {
			if userName == user.User && pass == user.Password {
				return user, nil
			}
		}
		return models.User{}, fmt.Errorf("User %s not found", userName)
	}

	db, err := openDb()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s WHERE user = ? AND password = ?;", istruct.tableName))
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

// GetSingleByToken returns an unique user model by its token
func (istruct *UserInfrastructure) GetSingleByToken(token string) (models.User, error) {
	if isDbUpdated(istruct.rows, istruct.lastUpdate) {
		for _, user := range istruct.userDb {
			if token == user.Token.String {
				return user, nil
			}
		}
		return models.User{}, fmt.Errorf("User with given token not found")
	}

	db, err := openDb()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s WHERE token = ?;", istruct.tableName))
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

// TableName setter
func (i *UserInfrastructure) TableName(name string) {
	if i.tableName == "" {
		i.tableName = name
	}
}
