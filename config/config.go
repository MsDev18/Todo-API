package config

import (
	"todo-api/repository/mysql"
)

type AppConfig struct {
	Port uint   `koanf:"port"`
	Host string `koanf:"host"`
}

type Config struct {
	App AppConfig    `koanf:"app"`
	DB  mysql.Config `koanf:"db"`
}
