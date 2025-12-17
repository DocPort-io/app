package app

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config represents the application configuration loaded via Viper.
type Config struct {
	Database struct {
		Driver string `mapstructure:"driver" validate:"required"`
		URL    string `mapstructure:"url" validate:"required"`
	} `mapstructure:"database"`

	Server struct {
		Host string `mapstructure:"host"`
		Bind string `mapstructure:"bind" validate:"required"`
		Port string `mapstructure:"port" validate:"required"`
	} `mapstructure:"server"`

	Storage struct {
		Provider string `mapstructure:"provider" validate:"required"`
		Path     string `mapstructure:"path" validate:"required"`
	} `mapstructure:"storage"`
}

// LoadConfig reads configuration using Viper, unmarshals into a struct, and validates it.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/docport/")
	viper.AddConfigPath("$HOME/.docport")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("docport")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

//

// validateConfig performs semantic validation beyond struct tags.
func validateConfig(cfg *Config) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(cfg); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// Validate server port is a valid TCP port number and non-empty
	if _, err := parsePort(cfg.Server.Port); err != nil {
		return fmt.Errorf("server.port invalid: %w", err)
	}

	// Validate storage.path exists and is a directory (only for filesystem provider)
	if cfg.Storage.Provider == "filesystem" {
		if err := ensureDirExists(cfg.Storage.Path); err != nil {
			return fmt.Errorf("storage.path invalid: %w", err)
		}
	}

	// Validate database DSN by attempting to open and ping
	if err := validateDatabaseDSN(cfg.Database.Driver, cfg.Database.URL); err != nil {
		return fmt.Errorf("database URL invalid: %w", err)
	}

	return nil
}

func parsePort(p string) (int, error) {
	if p == "" {
		return 0, errors.New("empty port")
	}
	// net.JoinHostPort validates when combining, use that
	if _, err := net.ResolveTCPAddr("tcp", net.JoinHostPort("127.0.0.1", p)); err != nil {
		return 0, fmt.Errorf("not a valid port: %w", err)
	}
	// successful resolution implies numeric 1..65535; return dummy value
	return 0, nil
}

func ensureDirExists(path string) error {
	if path == "" {
		return errors.New("path is empty")
	}
	abs := path
	if !filepath.IsAbs(path) {
		if a, err := filepath.Abs(path); err == nil {
			abs = a
		}
	}
	info, err := os.Stat(abs)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("directory does not exist: %s", abs)
		}
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("not a directory: %s", abs)
	}
	return nil
}

func validateDatabaseDSN(driver, dsn string) error {
	if driver == "" || dsn == "" {
		return errors.New("driver or dsn is empty")
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Ping()
}
