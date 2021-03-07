package config

import (
	"github.com/go-ini/ini"
)

type Database struct {
	IP       string `ini:"ip"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
}

type IniConfig struct {
	Database `ini:"database"`
}

var Config = &IniConfig{}

func Init() error {
	path := "articles/singleton-database/conf/database.conf"
	return InitFromFile(path)
}

func InitFromFile(path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}

	return cfg.MapTo(Config)
}