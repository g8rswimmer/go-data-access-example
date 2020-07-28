package database

const user = `
CREATE TABLE IF NOT EXISTS user (
	id CHAR(36) NOT NULL,
	fist_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP, 
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	deleted_at DATETIME,
	PRIMARY KEY (id)
)
`
