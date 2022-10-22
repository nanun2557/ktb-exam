package repository

import (
	"database/sql"
)

type Repository struct {
	DB          *sql.DB
	ProductRepo ProductRepo
}

func New(db *sql.DB) *Repository {
	return &Repository{
		DB:          db,
		ProductRepo: &ProductRepoContext{db},
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
