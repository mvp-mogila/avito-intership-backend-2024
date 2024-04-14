package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const _configpath = "config.yaml"

type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
	Redis    RedisConfig    `yaml:"redis"`
	Auth     AuthConfig     `yaml:"auth"`
}

type AuthConfig struct {
	AdminToken string `yaml:"admin_token"`
	UserToken  string `yaml:"user_token"`
}

type PostgresConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	Sslmode      string `yaml:"sslmode"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleTime  int    `yaml:"max_idle_time"`
}

type RedisConfig struct {
	Host    string        `yaml:"Host"`
	Port    string        `yaml:"port"`
	ExpTime time.Duration `yaml:"expiration_time"`
}

func GetConfig() *Config {
	var cfg Config
	if err := cleanenv.ReadConfig(_configpath, &cfg); err != nil {
		log.Fatal("error while reading config: " + err.Error())
	}
	return &cfg
}
