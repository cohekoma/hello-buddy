package storage

import "database/sql"

type Storage struct {
	PostsStorage interface {
		Create() (err error, postId int)
	}

	UsersStorage interface {
		Create() (err error, userId int)
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		PostsStorage: &PostsStorage{db},
		UsersStorage: &UsersStorage{db},
	}
}
