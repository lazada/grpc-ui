package main

import (
	"github.com/lazada/grpc-ui/config"
	"github.com/lazada/grpc-ui/http_server"
	"log"
	"os"
)

func main() {
	cfg := config.MustFromFlags(&config.Config{
		StaticDir: "static",
	})

	srv := http_server.New(cfg.HttpAddr, cfg.StaticDir, cfg.TargetAddr)
	if err := srv.Start(); err != nil {
		log.Printf("http server error: %v", err)
		os.Exit(1)
	}
}

