package store

import (
	"database/sql"
	"practise/test_task_post/pkg/logger"
	"practise/test_task_post/storage"
)

type Store struct {
	db      *sql.DB
	log     logger.Logger
	openapi *openApiRepo
	// ctx context.Context
}

func New(db *sql.DB, log logger.Logger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}
func (s *Store) OpenApiStore() storage.OpenApiRepo {
	if s.openapi == nil {
		s.openapi = &openApiRepo{store: s}
	}
	return s.openapi
}
