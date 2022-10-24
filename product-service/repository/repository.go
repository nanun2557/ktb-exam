package repository

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	DB          *sql.DB
	Cache       *redis.Client
	ProductRepo ProductRepo
}

func New(db *sql.DB, cache *redis.Client) *Repository {
	return &Repository{
		DB:          db,
		Cache:       cache,
		ProductRepo: &ProductRepoContext{db, cache},
	}
}

func (s *Repository) Begin() (*sql.Tx, error) {
	return s.DB.Begin()
}

func (s *Repository) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (s *Repository) RollBack(tx *sql.Tx) error {
	return tx.Rollback()
}
