package main

import (
	"app/pkg/app"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	server := app.NewServer()

	go func() {
		log.Printf("listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to listen and serve: %s\n", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
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
