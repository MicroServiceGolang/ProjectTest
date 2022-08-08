package store

import (
	"database/sql"
	"practise/test_task_grud/pkg/logger"
	"practise/test_task_grud/storage"
)

type Store struct {
	db   *sql.DB
	log  logger.Logger
	post *postRepo
	// ctx context.Context
}

func New(db *sql.DB, log logger.Logger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}
func (s *Store) PostStore() storage.PostRepo {
	if s.post == nil {
		s.post = &postRepo{store: s}
	}
	return s.post
}
