package storage

import (
	"database/sql"
	"fmt"
)

type PostsStorage struct {
	db *sql.DB
}

func (p *PostsStorage) Create() (error, int) {
	fmt.Println("Creating new post...")
	return nil, 0
}
