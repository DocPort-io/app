package main

import (
	"app/pkg/app"
	"app/pkg/storage"
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/docport/")
	viper.AddConfigPath("$HOME/.docport")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("docport")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file, %s", err)
	}

	fileStorage := app.NewFileStorage(storage.Type(viper.GetString("storage.provider")))
	db, queries := app.NewDatabase()
	srv := app.NewServer(db, queries, fileStorage)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(viper.GetString("server.bind"), viper.GetString("server.port")),
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
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		log.Fatalf("%s\n", err)
	}
}
