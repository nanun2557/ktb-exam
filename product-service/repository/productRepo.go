package repository

import (
	"product-service/models"

	"database/sql"
)

type ProductRepo interface {
	Get(tx *sql.Tx) ([]models.Product, error)
	DeleteById(tx *sql.Tx, id int) error
}

type ProductRepoContext struct {
	*sql.DB
}

func (s *ProductRepoContext) Get(tx *sql.Tx) ([]models.Product, error) {
	products := make([]models.Product, 0)

	query := "SELECT id, name, author_id from products"

	rows, err := s.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b models.Product
		err = rows.Scan(&b.ID, &b.Name, &b.Price)
		products = append(products, b)
	}

	return products, nil
}

func (s *ProductRepoContext) DeleteById(tx *sql.Tx, id int) error {
	query := "DELETE FROM products WHERE products.id = $1 RETURNING products.id"

	row, err := s.Exec(query, id)
	if err != nil {
		return err
	}

	if r, err := row.RowsAffected(); err != nil {
		return err
	} else if r == 0 {
		return sql.ErrNoRows
	}

	return nil
}
