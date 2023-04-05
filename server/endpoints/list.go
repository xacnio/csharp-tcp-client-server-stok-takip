package endpoints

import (
	"dagitik-sistemler/server/database"
	"dagitik-sistemler/server/models"
	"strconv"
)

const (
	MAX_PAGE_SIZE = 100
)

func List(args ...string) string {
	page := 1
	if len(args) > 1 {
		page, _ = strconv.Atoi(args[1])
	}
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * MAX_PAGE_SIZE
	limit := MAX_PAGE_SIZE

	db, err := database.SQLiteConnection()
	if err != nil {
		return err.Error()
	}

	stoklar := models.Stoklar{}

	totalCount := int64(0)
	db.Model(models.Stok{}).Select("COUNT(*)").Count(&totalCount)

	e := db.Model(models.Stok{}).Offset(offset).Limit(limit).Order("id DESC").Find(&stoklar.Items).Error
	if e != nil {
		return e.Error()
	}
	stoklar.MaxItemCount = limit
	stoklar.Total = totalCount

	return stoklar.ToJSON()
}
