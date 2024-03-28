package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rusMatryoska/cpanel-go/internal/config"
	"github.com/rusMatryoska/cpanel-go/internal/storage/postgresql"

	"golang.org/x/exp/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	ctx := context.Background()

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)
	log.Debug("debug messages are enabled")

	/////////////////////////////////////////////////////////

	DBItem := &postgresql.Database{
		DBConnURL: fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.Storage.User, "pgpwd4habr", cfg.Storage.Address, cfg.Storage.DBName),
	}

	pool, err := DBItem.GetDBConnection(ctx)

	if err != "" {
		log.Error(err)
	}

	defer pool.Close()

	// DBItem.ConnPool = pool
	// DBItem.DBErrorConnect = dbErrorConnect

	// st = storage.Storage(DBItem)

	// TODO: init router

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
