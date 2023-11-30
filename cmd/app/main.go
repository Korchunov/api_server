package main

import (
	"fmt"
	"github.com/Korchunov/api_server.git/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	log := setupLogger(cfg.Env)
	log.Info("активирована инфа", slog.String("env", cfg.Env))
	log.Debug("активирован дебаг")
}

func setupLogger(env string) *slog.Logger {
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
