package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type ConfigDB struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	SSLMode  string
}

func NewMysqldb(cfg ConfigDB) (*sql.DB, error) {
	param := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)

	db, err := sql.Open("mysql", param)

	if err != nil {
		return nil, err
	}

	return db, nil
}
