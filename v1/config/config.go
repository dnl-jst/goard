package config

import (
	"code.google.com/p/gcfg"
	"flag"
)

var Std *Config

type Config struct {
	Http struct {
		Listen        string
	}

	Goard struct {
		Workspace      string
	}
}

func init() {
	ConfigPath := flag.String("conf", "./config.gcfg", "Specify a file containing the configuration or \"./config.gcfg\" is used")
	flag.Parse()
	Std = &Config{}
	gcfg.ReadFileInto(Std, *ConfigPath)
}
