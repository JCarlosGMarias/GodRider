package infrastructures

import (
	"fmt"
	"time"

	"godrider/infrastructures/models"
)

// UserProviderInfrastructurer provides edition access to the userprovider's table for rlations between users and providers
type UserProviderInfrastructurer interface {
	// InsertSingle should create a new connection row
	InsertSingle(userProvider *models.UserProvider) error
	// UpdateSingle should edit a connection's active field; its used to activate/deactivate an userprovider connection
	UpdateSingle(userProvider *models.UserProvider) error
}

// UserProviderInfrastructure is UserProviderInfrastructurer's implementation struct
type UserProviderInfrastructure struct {
	userProviderDb []models.UserProvider
	tableName      string
	rows           int
	lastUpdate     time.Time
}

// InsertSingle creates a new connection row
func (istruct *UserProviderInfrastructure) InsertSingle(userProvider *models.UserProvider) error {
	db, err := openDb()
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

	if count, err := result.RowsAffected(); err != nil || count == 0 {
		return fmt.Errorf("No new rows created")
	}

	istruct.userProviderDb = append(istruct.userProviderDb, *userProvider)
	istruct.rows++
	istruct.lastUpdate = time.Now()
	return nil
}

// UpdateSingle edit a connection's active field; its used to activate/deactivate an userprovider connection
func (istruct *UserProviderInfrastructure) UpdateSingle(userProvider *models.UserProvider) error {
	db, err := openDb()
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

	if count, err := result.RowsAffected(); err != nil || count == 0 {
		return fmt.Errorf("No rows updated")
	}

	for index, row := range istruct.userProviderDb {
		if row.UserId == userProvider.UserId && row.ProviderId == userProvider.ProviderId {
			istruct.userProviderDb[index].IsActive = userProvider.IsActive
			break
		}
	}

	istruct.lastUpdate = time.Now()
	return nil
}

// TableName setter
func (i *UserProviderInfrastructure) TableName(name string) {
	if i.tableName == "" {
		i.tableName = name
	}
}
