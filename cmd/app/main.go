package main

import (
	"app/pkg/app"
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg, err := app.LoadConfig()
	if err != nil {
		log.Fatalf("invalid configuration: %s", err)
	}

	fileStorage := app.NewFileStorage(cfg)
	queries := app.NewDatabase(cfg)
	srv := app.NewServer(queries, fileStorage)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Bind, cfg.Server.Port),
		Handler: srv,
	}

	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to listen and serve: %s\n", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("failed to shutdown app: %s\n", err)
	}
	log.Println("app shutdown complete")

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatalf("%s\n", err)
	}
}
