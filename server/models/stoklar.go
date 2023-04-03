package models

import (
	"encoding/json"
	"math/rand"
)

type Stoklar struct {
	Items        []Stok `json:"items"`
	Total        int64  `json:"total"`
	MaxItemCount int    `json:"max_item_count"`
}

func (s Stoklar) ToJSON() string {
	jsonEncoding, _ := json.Marshal(s)
	return string(jsonEncoding)
}

func GenerateRandomStok(count int) Stoklar {
	urunTurleri := []string{"Çanta", "Ayakkabı", "Cüzdan", "Kol Saati", "Kazak", "Pantolon", "Gömlek", "Şapka", "Kemer", "Küpe", "Kolye", "Bilezik"}
	urunAltTurler := []string{"Erkek", "Kadın", "Çocuk"}
	urunModeller := []string{"Klasik", "Spor", "Günlük"}
	urunRenkler := []string{"Siyah", "Beyaz", "Kırmızı", "Yeşil", "Mavi", "Mor", "Sarı", "Turuncu", "Gri", "Bordo", "Pembe", "Lacivert", "Burgu", "Bej", "Mavi", "Kahverengi", "Gümüş"}

	stoklar := Stoklar{Total: int64(count), Items: make([]Stok, 0)}
	for i := 1; i <= count; i++ {
		stoklar.Items = append(stoklar.Items, Stok{
			Isim:  urunTurleri[rand.Intn(len(urunTurleri))] + " " + urunAltTurler[rand.Intn(len(urunAltTurler))] + " " + urunModeller[rand.Intn(len(urunModeller))] + " " + urunRenkler[rand.Intn(len(urunRenkler))],
			Fiyat: float64(rand.Intn(1400) + 100),
			Sayi:  rand.Intn(100),
		})
	}
	return stoklar
}
