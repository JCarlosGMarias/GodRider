package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"godrider/controllers"
	"godrider/infrastructures/models"

	_ "modernc.org/sqlite"
)

func CreateDb() {
	db, _ := sql.Open("sqlite", "./db/godrider.db")
	defer db.Close()

	createQuery := "DROP TABLE IF EXISTS user; CREATE TABLE user (" +
		"ID INTEGER PRIMARY KEY, " +
		"Token TEXT, " +
		"user TEXT, password TEXT, name TEXT, surname TEXT, email TEXT, phone TEXT, level INTEGER);"
	statement, _ := db.Prepare(createQuery)
	statement.Exec()
	statement.Close()

	addUserQuery := "INSERT INTO user (user, password, name, surname, email, phone, level) " +
		`VALUES ("John", "Salchichon", "John", "Salchichon", "asdf@omg.god", "666333987", 1);`
	statement, _ = db.Prepare(addUserQuery)
	_, err := statement.Exec()

	if err == nil {
		rows, _ := db.Query("SELECT * FROM user;")
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.ID, &user.Token, &user.User, &user.Password, &user.Name, &user.Surname, &user.Email, &user.Phone, &user.Level); err != nil {
				fmt.Println("oops")
			} else {
				fmt.Println(user)
			}
		}
	}
	statement.Close()
}

func main() {
	routes := controllers.GetRoutes()
	//CreateDb()

	http.HandleFunc(routes["LoginUrl"], controllers.Login)
	// Endpoints
	http.HandleFunc(routes["GetApiUrlsUrl"], controllers.GetApiUrls)
	// Providers
	http.HandleFunc(routes["GetProvidersUrl"], controllers.GetProviders)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
