package database

import (
	"fmt"

	"github.com/hardzal/portfolio-api-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.ConfigDB) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	// Coba koneksi ke database target
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Fallback ke database default
		fallbackDSN := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s sslmode=disable",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPassword,
		)
		db, err = gorm.Open(postgres.Open(fallbackDSN), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		// Create database jika tidak ada
		db.Exec("CREATE DATABASE " + cfg.DBName)
	}
	return db, nil
}
