package db

import (
	"database/sql"
	"log"
	"srebootcamp/internal/model"

	"github.com/lib/pq"
)

func ConnectDB(cfg pq.Config) *sql.DB {
	c, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}
	db := sql.OpenDB(c)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDB(config model.Config) *sql.DB {
	cfg := pq.Config{
		Host:     config.DBHost,
		Port:     uint16(config.DBPort),
		User:     config.DBUser,
		Password: config.DBPass,
		Database: config.DBName,
		SSLMode:  pq.SSLMode(config.SSLMode),
	}

	db := ConnectDB(cfg)

	return db

}
