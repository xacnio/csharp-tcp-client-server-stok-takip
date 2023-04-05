package endpoints

import (
	"dagitik-sistemler/server/database"
	"dagitik-sistemler/server/models"
	"strconv"
)

func Delete(args ...string) string {
	if len(args) < 2 {
		return ""
	}
	id, _ := strconv.Atoi(args[1])

	db, err := database.SQLiteConnection()
	if err != nil {
		return err.Error()
	}

	stok := models.Stok{
		Id: id,
	}
	e := db.Delete(&stok).Error
	if e != nil {
		return e.Error()
	}

	return "OK"
}
