package models

import (
	"dagitik-sistemler/server/utils"
	"gorm.io/gorm"
)

type Stok struct {
	Id               int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Isim             string  `gorm:"type:varchar(100);index;unique" json:"isim"`
	Fiyat            float64 `gorm:"type:decimal(10,2)" json:"fiyat"`
	PromosyonluFiyat float64 `gorm:"-" json:"promosyonlu_fiyat"`
	Sayi             int     `gorm:"type:int" json:"sayi"`
	CreatedAt        int64   `gorm:"autoCreateTime" json:"created_at"`
}

func (s *Stok) AfterCreate(tx *gorm.DB) (err error) {
	if s.Id > 0 {
		s.PromosyonluFiyat = utils.CalculatePromoPrice(s.Fiyat)
	}
	return
}

func (s *Stok) AfterFind(tx *gorm.DB) (err error) {
	if s.Id > 0 {
		s.PromosyonluFiyat = utils.CalculatePromoPrice(s.Fiyat)
	}
	return
}
