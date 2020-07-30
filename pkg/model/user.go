package model

// User is the structure for an user
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// UserEntity is the user entity for the database
type UserEntity struct {
	Entity
	User
}
