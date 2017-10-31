package config

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)


type Config struct {
	HttpAddr string
	TargetAddr string
	StaticDir string
}


func MustFromFlags(defaults *Config) *Config {
	cfg := defaults

	flag.StringVar(&cfg.HttpAddr, "http", "", "http server listening addr")
	flag.StringVar(&cfg.TargetAddr, "target", "", "grpc server with reflection addr")
	flag.StringVar(&cfg.StaticDir, "static_dir", cfg.StaticDir, "static files directory")

	flag.Parse()

	if cfg.HttpAddr == "" {
		flag.Usage()
		os.Exit(2)
		return nil
	}

	if cfg.TargetAddr == "" {
		flag.Usage()
		os.Exit(2)
		return nil
	}

	if err := cfg.validate(); err != nil {
		log.Printf("invalid config: %v", err)
		os.Exit(2)
	}

	return  cfg
}

func (c *Config) validate() error {
	host, _, err := net.SplitHostPort(c.TargetAddr)
	if err != nil {
		return fmt.Errorf("invalid addr: %v", c.TargetAddr)
	}
	if host == "" {
		return fmt.Errorf("empty host: %v", c.TargetAddr)

	}
	return nil
}