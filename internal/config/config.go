package config

import (
	"flag"
	"log"

	"github.com/notnull-co/cfg"
)

var (
	Cfg Config
)

type Config struct {
	Token struct {
		Key string `cfg:"key"`
	} `cfg:"token"`
	CustomerServer struct {
		Port int `cfg:"port"`
	} `cfg:"customerServer"`
	UserServer struct {
		Port int `cfg:"port"`
	} `cfg:"userServer"`
	DB struct {
		ConnectionString string `cfg:"connection_string"`
	} `cfg:"db"`
	Port string `cfg:"port"`
}

func ParseFromFlags() {
	var configDir string

	flag.StringVar(&configDir, "config-dir", "../../internal/config/", "Configuration file directory")
	flag.Parse()

	parse(configDir)
}

func parse(dirs ...string) {
	if err := cfg.Load(&Cfg,
		cfg.Dirs(dirs...),
		cfg.UseEnv("app"),
	); err != nil {
		log.Panic(err)
	}
}

func Get() Config {
	return Cfg
}
