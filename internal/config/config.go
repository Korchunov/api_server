package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type conf struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_Timeout" env-default:"60s"`
}

// Must... означает, что процедура будет возвращать не ошибку, а Fatal.
// Это как раз актуально при первом запуске приложения
func MustLoad() *conf {
	//Проверяем, есть ли у переменной окружения какое-либо значение
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	// Проверяем, существует ли файл вообще
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config fatal %s does not exist", configPath)
	}
	//Читаем файл конфигурации и маппим со структурой
	var cfg conf
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("cannot read config: %s", err)
	}
	return &cfg
}
