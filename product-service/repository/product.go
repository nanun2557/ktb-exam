package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"product-service/models"
	"time"

	"database/sql"

	"github.com/go-redis/redis/v8"
)

type ProductRepo interface {
	Get(tx *sql.Tx) ([]models.Product, error)
	Create(tx *sql.Tx, product *models.Product) (int64, error)
	UpdateById(tx *sql.Tx, product *models.Product) error
	DeleteById(tx *sql.Tx, id int) error

	SetCache(ctx context.Context, key string, value interface{}, expireTime time.Duration) error
	GetCache(ctx context.Context, key string, dest interface{}) error
}

type ProductRepoContext struct {
	db    *sql.DB
	cache *redis.Client
}

func (r *ProductRepoContext) Get(tx *sql.Tx) ([]models.Product, error) {
	products := make([]models.Product, 0)

	query := "SELECT id, name, price from products"
	rows, err := r.db.Query(query)
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

func (r *ProductRepoContext) Create(tx *sql.Tx, product *models.Product) (int64, error) {
	var err error

	query := "INSERT INTO products (name, price) VALUES (?, ?)"
	if err != nil {
		return 0, err
	}

	var id int64

	res, err := r.db.Exec(query, product.Name, product.Price)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (r *ProductRepoContext) UpdateById(tx *sql.Tx, product *models.Product) error {
	query := fmt.Sprintf("UPDATE products SET name = '%s', price = %d WHERE id = %d", product.Name, product.Price, product.ID)
	log.Println("query: ", query)

	res, err := r.db.Exec(query)
	if err != nil {
		log.Println(err)
		return err
	}

	if r, err := res.RowsAffected(); err != nil {
		log.Println(err)
		return err
	} else if r == 0 {
		log.Println(err)
		return sql.ErrNoRows
	}

	return nil
}

func (r *ProductRepoContext) DeleteById(tx *sql.Tx, id int) error {
	query := fmt.Sprintf("DELETE FROM products WHERE products.id = %d", id)
	log.Println("query: ", query)

	row, err := r.db.Exec(query)
	if err != nil {
		log.Println(err)
		return err
	}

	if r, err := row.RowsAffected(); err != nil {
		log.Println(err)
		return err
	} else if r == 0 {
		log.Println(err)
		return sql.ErrNoRows
	}

	return nil
}

func (r *ProductRepoContext) SetCache(ctx context.Context, key string, value interface{}, expireTime time.Duration) error {
	value, err := json.Marshal(value)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return err
	}
	err = r.cache.Set(ctx, key, value, expireTime).Err()
	if err != nil {
		fmt.Println("redis set error:", err)
		return err
	}
	return nil
}

func (r *ProductRepoContext) GetCache(ctx context.Context, key string, dest interface{}) error {
	result := r.cache.Get(ctx, key)
	if result.Val() == redis.Nil.Error() {
		return errors.New("not found key")
	}

	data, err := result.Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &dest)
	if err != nil {
		return err
	}

	return nil
}
