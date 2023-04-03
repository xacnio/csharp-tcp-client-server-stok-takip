package models

type Stok struct {
	Id        int     `gorm:"primaryKey" json:"id"`
	Isim      string  `gorm:"type:varchar(100);unique" json:"isim"`
	Fiyat     float64 `gorm:"type:decimal(10,2)" json:"fiyat"`
	Sayi      int     `gorm:"type:int" json:"sayi"`
	CreatedAt int64   `gorm:"autoCreateTime" json:"created_at"`
}
