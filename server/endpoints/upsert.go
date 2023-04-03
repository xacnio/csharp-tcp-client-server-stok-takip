package endpoints

import (
	"daginik-sistemler/server/database"
	"daginik-sistemler/server/models"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func Upsert(args ...string) string {
	if len(args) < 4 {
		return ""
	}
	name := args[1]
	price, e := strconv.ParseFloat(args[2], 64)
	if e != nil {
		return e.Error()
	}
	sayi, e := strconv.Atoi(args[3])
	if e != nil {
		return e.Error()
	}

	db, err := database.SQLiteConnection()
	if err != nil {
		return err.Error()
	}

	// Check if the item exists
	stok := models.Stok{}
	e = db.Model(models.Stok{}).Where("isim = ?", name).First(&stok).Error
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
