package dal

// UserTable defines the user table for the dal
const UserTable = `
CREATE TABLE IF NOT EXISTS user (
	id CHAR(36) NOT NULL,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP, 
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	deleted_at DATETIME,
	PRIMARY KEY (id)
)
`
