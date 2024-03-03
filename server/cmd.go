package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/OLIENTTECH/backend-challenges/infrastructure/datastore"
	"github.com/OLIENTTECH/backend-challenges/infrastructure/external/db/postgres"
	"github.com/OLIENTTECH/backend-challenges/pkg/log"
	"github.com/OLIENTTECH/backend-challenges/ui"
	"github.com/OLIENTTECH/backend-challenges/usecase"
)

const (
	httpAddr = ":8080"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "OLIENTTECH backend challenges server",
	}
	cmd.RunE = func(_ *cobra.Command, _ []string) error {
		return run()
	}

	return cmd
}

func run() error {
	logger := log.New()

	db, err := postgres.GetDBConnection()
	if err != nil {
		log.Panic("server: failed to connect to DB", log.Ferror(err))
	}
	dbClient := postgres.NewClient(db)
	txManager := postgres.NewTxManager(db)

	datastore := datastore.NewDataStore(dbClient)
	usecase := usecase.NewUsecase(txManager, datastore, logger)
	handler := ui.NewHandler(usecase, logger)
	server := newEchoServer(handler)

	// Start server
	go func() {
		if err := server.Start(httpAddr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server: failed to start HTTP server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 30 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) //nolint:gomnd
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error("server: failed to shutdown server gracefully", zap.Error(err))
	}

	return nil
}
