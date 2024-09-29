package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env              string `yaml:"env" env-default:"local"`
	HTTPServerConfig `yaml:"http_server"`
	DBConfig         `yaml:"db"`
}

type HTTPServerConfig struct {
	Address     string        `yaml:"address" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idleTimeout" env-default:"120s"`
}

type DBConfig struct {
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"test"`
	DBName   string `yaml:"dbName" env-default:"hw6"`
	Password string `yaml:"password" env-default:"12345"`
	SslMode  string `yaml:"sslMode" env-default:"disable"`
}

func DefaultConfig() Config {
	configPath := "../../config/config.yaml"

	fmt.Println(configPath)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist")
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	return config
}
