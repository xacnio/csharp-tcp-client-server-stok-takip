package main

import (
	"dagitik-sistemler/server/database"
	"dagitik-sistemler/server/utils"
)

// TCP socket dinlemeyi başlat ve client bağlantıları bekle
func main() {
	err := database.CheckConnection()
	if err != nil {
		panic(err)
	}

	utils.Listen()
}
