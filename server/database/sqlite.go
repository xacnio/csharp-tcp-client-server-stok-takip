package database

import (
	"dagitik-sistemler/server/models"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	CON_DB_FILE = "stok.db"
)

func SQLiteConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(CON_DB_FILE), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SQLiteMigrate(db *gorm.DB) {
	ifExistStok := db.Migrator().HasTable(&models.Stok{})

	db.AutoMigrate(&models.Stok{})

	if !ifExistStok {
		fmt.Println("Stok tablosu oluşturuldu. Veriler ekleniyor...")
		db.AutoMigrate(&models.Stok{})
		stoks := models.GenerateRandomStok(2000)
		for _, stok := range stoks.Items {
			db.Create(&stok)
		}
	}
}

func CheckConnection() error {
	// Database bağlantısı kontrolü ve migration
	db, err := SQLiteConnection()
	if err != nil {
		return err
	}
	SQLiteMigrate(db)
	return nil
}
