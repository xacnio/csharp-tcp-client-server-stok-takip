package main

import (
	"dagitik-sistemler/server/database"
	"dagitik-sistemler/server/socket"
)

// TCP socket dinlemeyi başlat ve client bağlantıları bekle
func main() {
	err := database.CheckConnection()
	if err != nil {
		panic(err)
	}

	socket.Listen()
}
