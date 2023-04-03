package main

import (
	"daginik-sistemler/server/database"
	"daginik-sistemler/server/utils"
)

// TCP socket dinlemeyi başlat ve client bağlantıları bekle
func main() {
	err := database.CheckConnection()
	if err != nil {
		panic(err)
	}

	utils.Listen()
}
