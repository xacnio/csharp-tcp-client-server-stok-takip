package endpoints

import (
	"dagitik-sistemler/server/database"
	"dagitik-sistemler/server/models"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func Upsert(args ...string) string {
	if len(args) < 5 {
		return ""
	}
	id, e := strconv.Atoi(args[1])
	if e != nil {
		return e.Error()
	}
	name := args[2]
	price, e := strconv.ParseFloat(args[3], 64)
	if e != nil {
		return e.Error()
	}
	sayi, e := strconv.Atoi(args[4])
	if e != nil {
		return e.Error()
	}

	db, err := database.SQLiteConnection()
	if err != nil {
		return err.Error()
	}

	// Check if the item exists
	stok := models.Stok{}
	e = db.Model(models.Stok{}).Where("id = ?", id).First(&stok).Error
	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			stok := models.Stok{
				Isim:  name,
				Fiyat: price,
				Sayi:  sayi,
			}
			e := db.Create(&stok).Error
			if e != nil {
				return e.Error()
			}
		} else {
			return e.Error()
		}
	} else {
		stok.Sayi = sayi
		stok.Fiyat = price
		e := db.Save(&stok).Error
		if e != nil {
			return e.Error()
		}
	}

	return "OK"
}
