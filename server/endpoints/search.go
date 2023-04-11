package endpoints

import (
	"dagitik-sistemler/server/database"
	"dagitik-sistemler/server/models"
	"dagitik-sistemler/server/utils"
	"strconv"
	"strings"
)

func Search(args ...string) string {
	page := 1
	if len(args) != 3 {
		return "Geçersiz istek"
	}
	page, _ = strconv.Atoi(args[1])
	if page < 1 {
		page = 1
	}
	search := args[2]

	offset := (page - 1) * MAX_PAGE_SIZE
	limit := MAX_PAGE_SIZE

	db, err := database.SQLiteConnection()
	if err != nil {
		return err.Error()
	}

	stoklar := models.Stoklar{}
	totalCount := int64(0)
	isimColumn := "replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(lower(isim), 'Ç', 'c'), 'Ğ', 'g'), 'Ş', 's'), 'ç','c'), 'ğ','g'), 'İ','i'),'ş','s'),'Ö','o'),'ö','o'),'Ü','u'),'ü','u'),'ı','i')"
	search = utils.TurkishToEnglish(strings.ToLower(search))
	db.Model(models.Stok{}).Where(isimColumn+" LIKE ?", "%"+search+"%").Select("COUNT(*)").Count(&totalCount)

	e := db.Model(models.Stok{}).Where(isimColumn+" LIKE ?", "%"+search+"%").Offset(offset).Limit(limit).Order("id DESC").Find(&stoklar.Items).Error
	if e != nil {
		return e.Error()
	}
	stoklar.MaxItemCount = limit
	stoklar.Total = totalCount

	return stoklar.ToJSON()
}
