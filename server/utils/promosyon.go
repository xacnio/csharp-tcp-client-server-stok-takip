package utils

import (
	"math"
	"strconv"
)

const OGRENCI_NO string = "201913172017"

func CalculatePromoPrice(price float64) float64 {
	//-promosyon kodu (öğrenci numarası verilerek: ABCD1317EFGH hesaplanacak)
	//Promosyon(ABCD1317EFGH)
	//2023-ABCD=yıl=(1,2,3,4,5,6)
	//Promosyon = (E*x^3+F*x^2+G*x+H)/x^4 =   x=yıl   (en küçük)
	//Promosyonlu fiyat = (1-promosyon)*fiyat
	yilOgr, _ := strconv.Atoi(OGRENCI_NO[0:4])
	yil := 2023 - yilOgr
	E, _ := strconv.ParseFloat(string(OGRENCI_NO[0]), 64)
	F, _ := strconv.ParseFloat(string(OGRENCI_NO[1]), 64)
	G, _ := strconv.ParseFloat(string(OGRENCI_NO[2]), 64)
	H, _ := strconv.ParseFloat(string(OGRENCI_NO[3]), 64)
	promosyon := (E*float64(yil)*math.Pow(float64(yil), 2) + F*float64(yil)*float64(yil) + G*float64(yil) + H) / math.Pow(float64(yil), 4)
	return (1 - promosyon) * price
}
