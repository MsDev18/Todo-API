package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	// TODO - review
	Name     string `koanf:"name"`
	Host     string `koanf:"host"`
	Port     uint   `koanf:"port"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}

type MySQLDB struct {
	config Config
	Db     *sql.DB
}

func New(cfg Config) *MySQLDB {
	// Generate DataSource
	// DataSource structer
	// user:password@(host:port)/dbname?param...
	dataSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(fmt.Sprintf("panic on connection to databse => %s", err))
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{
		Db: db,
	}
}
