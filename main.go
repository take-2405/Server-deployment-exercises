  
package main

import (
	"Attendance/pkg/server"
	"flag"
	"log"
	"os"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	port := os.Getenv("PORT")
	port =":"+port
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	flag.StringVar(&addr, "addr", port, "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Server.Run(addr)
}