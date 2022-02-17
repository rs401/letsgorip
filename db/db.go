package db

import (
	"log"

	authModels "github.com/rs401/letsgorip/auth/models"
	forumsModels "github.com/rs401/letsgorip/forums/models"
	placesModels "github.com/rs401/letsgorip/places/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection interface {
	DB() *gorm.DB
}

type conn struct {
	db *gorm.DB
}

func (c *conn) DB() *gorm.DB {
	return c.db
}

func NewConnection(cfg Config) (Connection, error) {
	dbc, err := gorm.Open(postgres.Open(cfg.ConnStr()), &gorm.Config{})
	if err != nil {
		log.Printf("Error, could not connect to database: %v", err)
		return nil, err
	}
	dbc.AutoMigrate(
		&authModels.User{},
		&forumsModels.Forum{},
		&forumsModels.Thread{},
		&forumsModels.Post{},
		&placesModels.Place{},
	)
	return &conn{db: dbc}, nil
}
