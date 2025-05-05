package database

import (
	"fmt"
	"log"

	"github.com/hardzal/portfolio-api-go/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := db.AutoMigrate(
		&models.About{},
		&models.Project{},
		&models.Stack{},
		&models.User{},
		&models.Work{},
	); err != nil {
		log.Fatal("Gagal migrasi table ke database.", err)
		return err
	}

	// stmt := db.Session(&gorm.Session{DryRun: true}).Migrator().CreateTable(&models.About{})
	// fmt.Printf("error gak sih ini: %v", stmt)

	// if !db.Migrator().HasTable(&models.About{}) {
	// 	if err := db.Migrator().CreateTable(&models.About{}); err != nil {
	// 		log.Fatal("CreateTable About:", err)
	// 		return err
	// 	}
	// }

	fmt.Println("DB Migrated")

	return nil
}
