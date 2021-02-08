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
	db, _ := sql.Open("sqlite", "./db/godrider.db")
	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO userprovider (user_id, provider_id, is_active) VALUES (?, ?, ?);")
	result, err := statement.Exec(userProvider.UserId, userProvider.ProviderId, userProvider.IsActive)
	defer statement.Close()

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
