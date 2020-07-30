package model

import (
	"database/sql"
	"time"
)

// Entity contains the basic fields for database entities
type Entity struct {
	ID        string       `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
