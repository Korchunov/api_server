package main

import (
	"fmt"
	"github.com/Korchunov/api_server.git/internal/config"
	"github.com/Korchunov/api_server.git/internal/lib/logger/sl"
	"github.com/Korchunov/api_server.git/internal/storage/sqlite"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	if err := godotenv.Load("enviroment.env"); os.IsNotExist(err) {
		log.Fatal("фатальный лог")
	}
	cfg := config.MustLoad()
	//Отладочный Print
	fmt.Println("это cfg:\t", cfg)
	log := setupLogger(cfg.Env)
	//Отладочный Print
	fmt.Println("это log\t", log)
	log.Info("активирована инфа", slog.String("env", cfg.Env))
	log.Debug("активирован дебаг", slog.String("env", cfg.Env))

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
}

func setupLogger(env string) *slog.Logger {
	//Пакет log: logg := log.New(os.Stdout, "ERR:\t", log.LstdFlags|log.Lshortfile)
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	return log
}
