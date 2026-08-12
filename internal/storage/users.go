package storage

import (
	"database/sql"
	"fmt"
)

type UsersStorage struct {
	db *sql.DB
}

func (p *UsersStorage) Create() (error, int) {
	fmt.Println("Creating new user...")
	fmt.Println("Done - New user is created!")
	return nil, 0
}
