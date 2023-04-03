package utils

import (
	"daginik-sistemler/server/endpoints"
	"fmt"
	"net"
	"regexp"
)

// Ayarlar
const (
	CON_LISTEN_HOST  = "127.0.0.1" // local
	CON_LISTEN_PORT  = "4444"
	CON_TYPE         = "tcp4"
	CON_MAX_READ_LEN = 1 * 1024 // 1 KB
)

// Gelen bağlantıyı işle ve cevapla
func processRequest(conn net.Conn) {
	for {
		buf := make([]byte, CON_MAX_READ_LEN)
		reqLen, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			return
		}
		go func() {
			req := string(buf[:reqLen])

			type Command struct {
				RegexPattern string
				Func         func(...string) string
			}
			commands := []Command{
				{`^search:([0-9]*):(.+)$`, endpoints.Search},
				{`^list:([0-9]*)$`, endpoints.List},
				{`^delete:([0-9]+)$`, endpoints.Delete},
				{`^upsert:(.+):([0-9.]+):([0-9]+)$`, endpoints.Upsert},
			}

			for _, command := range commands {
				matches := regexp.MustCompile(command.RegexPattern).FindStringSubmatch(req)
				if len(matches) > 0 {
					result := command.Func(matches...)
					conn.Write([]byte(result))
					return
				}
			}

			conn.Write([]byte("Not found"))
		}()
	}
}

// Gelen bağlantıları kabul et ve concurrent işle
func Receive(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processRequest(conn)
	}
}

// TCP socket port oluştur ve dinlemeye başla
func Listen() {
	hostConfig := CON_LISTEN_HOST + ":" + CON_LISTEN_PORT
	listener, err := net.Listen(CON_TYPE, hostConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on " + hostConfig + "...")
	defer listener.Close()

	// Client bağlantıları bekle
	Receive(listener)
}
