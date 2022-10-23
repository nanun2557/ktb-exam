package repository

import (
	"log"
	"product-service/models"

	"database/sql"
)

type ProductRepo interface {
	Get(tx *sql.Tx) ([]models.Product, error)
	Create(tx *sql.Tx, product *models.Product) (int64, error)
	UpdateById(tx *sql.Tx, product *models.Product) (int64, error)
	DeleteById(tx *sql.Tx, id int) error
}

type ProductRepoContext struct {
	db *sql.DB
}

func (s *ProductRepoContext) Get(tx *sql.Tx) ([]models.Product, error) {
	products := make([]models.Product, 0)

	query := "SELECT id, name, price from products"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b models.Product
		err = rows.Scan(&b.ID, &b.Name, &b.Price)
		if err != nil {
			log.Println(err)
			return products, err
		}
		products = append(products, b)
	}

	return products, nil
}

func (s *ProductRepoContext) Create(tx *sql.Tx, product *models.Product) (int64, error) {
	var err error

	query := "INSERT INTO products (name, price) VALUES ($1, $2)"
	if err != nil {
		return 0, err
	}

	var id int64

	err = s.db.QueryRow(query, product.Name, product.Price).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (s *ProductRepoContext) UpdateById(tx *sql.Tx, product *models.Product) (int64, error) {
	query, err := s.db.Prepare("UPDATE products SET name = $1, price = $2 WHERE product.id = $3 RETURNING id")
	if err != nil {
		return 0, err
	}

	var id int64

	err = query.QueryRow(product.Name, product.Price, product.ID).Scan(&id)
	if id == 0 {
		return 0, sql.ErrNoRows
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ProductRepoContext) DeleteById(tx *sql.Tx, id int) error {
	query := "DELETE FROM products WHERE products.id = $1 RETURNING products.id"

	row, err := s.db.Exec(query, id)
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
