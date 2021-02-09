package infrastructures

import (
	"database/sql"
	"fmt"
	"time"

	"godrider/infrastructures/models"

	_ "modernc.org/sqlite"
)

type UserProviderInfrastructure struct {
	userProviderDb []models.UserProvider
	tableName      string
	rows           int
	lastUpdate     time.Time
}

var UserProviderDb = UserProviderInfrastructure{tableName: "userprovider"}

func (istruct *UserProviderInfrastructure) InsertSingle(userProvider *models.UserProvider) error {
	db, err := sql.Open("sqlite", "./db/godrider.db")
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("INSERT INTO %s (user_id, provider_id, is_active) VALUES (?, ?, ?);", istruct.tableName))
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(userProvider.UserId, userProvider.ProviderId, userProvider.IsActive)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil || count == 0 {
		return fmt.Errorf("No new rows created!")
	}

	istruct.userProviderDb = append(istruct.userProviderDb, *userProvider)
	istruct.rows++
	istruct.lastUpdate = time.Now()
	return nil
}

func (istruct *UserProviderInfrastructure) UpdateSingle(userProvider *models.UserProvider) error {
	db, err := sql.Open("sqlite", "./db/godrider.db")
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("UPDATE %s SET is_active = ? WHERE user_id = ? AND provider_id = ?;", istruct.tableName))
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(userProvider.IsActive, userProvider.UserId, userProvider.ProviderId)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil || count == 0 {
		return fmt.Errorf("No rows updated!")
	}

	for _, row := range istruct.userProviderDb {
		if row.UserId == userProvider.UserId && row.ProviderId == userProvider.ProviderId {
			localUpdate := &row
			localUpdate.IsActive = userProvider.IsActive
			break
		}
	}

	istruct.lastUpdate = time.Now()
	return nil
}
