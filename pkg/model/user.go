package model

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserEntity struct {
	Entity
	User
}
