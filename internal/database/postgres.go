package database

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDB membuat koneksi database menggunakan GORM.
// databaseURL berasal dari .env.
// Contoh:
// postgresql://user:password@host:port/dbname
// Return-nya *gorm.DB.
// Nanti *gorm.DB ini dipakai oleh repository untuk melakuakan query ke database.

func NewPostgresDB(databaseUrl string) (*gorm.DB, error) {
	if databaseUrl == "" {
		return nil, errors.New("Database url cannot be empty")
	}

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

	if err != nil {
		return  nil, err
	}

	return db, nil
}