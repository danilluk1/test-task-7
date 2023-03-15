package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/danilluk1/test-task-7/internal/config"
	"github.com/danilluk1/test-task-7/internal/db/sqlc"
	"github.com/danilluk1/test-task/7/internal/telegram"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err)
	}

	if cfg.AppEnv == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := pgx.Connect(ctx, cfg.PostgresUrl)
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to database:")
	}
	defer conn.Close(ctx)

	store := db.NewStore(conn)

	tg := telegram.NewTelegram(cfg.TelegramToken)
	tg.StartPolling(ctx)

	log.Info().Msg("Started 🚀")
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
	log.Info().Msg("Closing...")
}
