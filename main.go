package main

import (
	"flag"
	"log"
	"os"

	"github.com/lazada/grpc-ui/http_server"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "", "http server listening addr")
	flag.Parse()

	if addr == "" {
		flag.Usage()
		os.Exit(2)
	}

	srv := http_server.New(addr)
	if err := srv.Start(); err != nil {
		log.Printf("http server error: %v", err)
		os.Exit(1)
	}
}
