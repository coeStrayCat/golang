package db

import (
	"database/sql"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt sql.NullTime
}
