package config

import (
	"flag"
	"os"
)


type Config struct {
	HttpAddr string
	StaticDir string
}


func MustFromFlags(defaults *Config) *Config {
	cfg := defaults

	flag.StringVar(&cfg.HttpAddr, "http", "", "http server listening addr")
	flag.StringVar(&cfg.StaticDir, "static_dir", cfg.StaticDir, "static files directory")

	flag.Parse()

	if cfg.HttpAddr == "" {
		flag.Usage()
		os.Exit(2)
		return nil
	}


	return  cfg
}
