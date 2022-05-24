package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// GetDb connects to the database and returns the orm.
func GetDb() *gorm.DB {
	dsn, ok := os.LookupEnv("DSN")

	if !ok {
		panic("DSN not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

type HexCounter struct {
	Hex   int `gorm:"primaryKey"`
	Count int `gorm:"not null;default:0"`
}
