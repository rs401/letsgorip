// Package db provides access to the database.
package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Config interface defines the methods of the db Config.
type Config interface {
	ConnStr() string
	DbName() string
}

type config struct {
	dbUser  string
	dbPass  string
	dbHost  string
	dbPort  int
	dbName  string
	connStr string
}

// NewConfig constructs and returns a new db Config.
func NewConfig() Config {
	var cfg config
	var err error
	cfg.dbUser = os.Getenv("DB_USER")
	cfg.dbPass = os.Getenv("DB_PASS")
	cfg.dbHost = os.Getenv("DB_HOST")
	cfg.dbName = os.Getenv("DB_NAME")
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error parsing DB_PORT: %v", err)
	}
	cfg.connStr = fmt.Sprintf("host=%s user=%s  password=%s  dbname=%s port=%d",
		cfg.dbHost, cfg.dbUser, cfg.dbPass, cfg.dbName, cfg.dbPort)
	return &cfg
}

// ConnStr returns the corresponding connection string.
func (c *config) ConnStr() string {
	return c.connStr
}

// DbName returns the corresponding database name.
func (c *config) DbName() string {
	return c.dbName
}
