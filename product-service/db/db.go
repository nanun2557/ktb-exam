package db

import (
	"database/sql"
	"fmt"
	"product-service/configs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// use Mysql for database
func NewMySql(c configs.MySqlConfiguration) (*sql.DB, error) {

	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Host, c.Port, c.DbName)
	// fmt.Println("mysql-uri: ", uri)

	db, err := sql.Open("mysql", uri)
	if err != nil {
		// return nil, err
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
