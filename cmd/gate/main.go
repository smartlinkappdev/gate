package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"cmd/gate/main.go/internal/app"
	"cmd/gate/main.go/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// получение конфига
	cfg := config.New()

	// установка логгера

	log := setupLogger(cfg.Env)
	log.Info("starting gate")

	// запуск приложения в горутине
	application := app.New(log)
	application.Init(cfg)
	go application.Start()

	// обработка остановки
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	s := <-stop
	log.Info("stopping application", slog.String("signal", s.String()))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelDebug,
				ReplaceAttr: nil,
			}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelDebug,
				ReplaceAttr: nil,
			}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelInfo,
				ReplaceAttr: nil,
			}),
		)
	}

	return log
}
