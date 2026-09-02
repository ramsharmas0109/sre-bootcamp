package db

import (
	"database/sql"
	"log"
	"srebootcamp/internal/config"

	"github.com/lib/pq" // To register the driver.
)

func ConnectDB(cfg pq.Config) *sql.DB {
	c, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Create connection pool.
	db := sql.OpenDB(c)
	// defer db.Close()

	// Make sure it works.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDB() *sql.DB {
	config := config.LoadConfig()

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

// user -> config -> connector -> db connect -> db pointer -> db query
