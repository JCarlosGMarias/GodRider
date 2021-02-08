package models

import "database/sql"

type User struct {
	ID       int
	Token    sql.NullString
	User     string
	Password string
	Name     string
	Surname  string
	Email    string
	Phone    string
	Level    int
}

type Provider struct {
	ID      int
	Name    string
	Contact string
}

type UserProvider struct {
	UserId     int
	ProviderId int
	IsActive   int
}

type ApiUrl struct {
	Key string
	Url string
}
