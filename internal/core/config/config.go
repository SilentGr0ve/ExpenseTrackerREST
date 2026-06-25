package config

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   ServerConfig
	Database PostgresConfig
	Logger   LoggerConfig
	JWT      JWTConfig
}

type PostgresConfig struct {
	Host     string        `envconfig:"HOST" 		required:"true"`
	Port     string        `envconfig:"PORT" 		default:"5432"`
	DBName   string        `envconfig:"DB" 			required:"true"`
	User     string        `envconfig:"USER" 		required:"true"`
	Password string        `envconfig:"PASSWORD" 	required:"true"`
	SSLMode  string        `envconfig:"SSLMODE" 	default:"disable"`
	MaxConns int32         `envconfig:"MAXCONNS"    default:"10"`
	Timeout  time.Duration `envconfig:"TIMEOUT" 	required:"true"`
}

type ServerConfig struct {
	Port        string `envconfig:"PORT" 	default:"8080"`
	Environment string `envconfig:"ENV" 	default:"development"`
}

type LoggerConfig struct {
	Level  string `envconfig:"LEVEL" 	default:"DEBUG"`
	Folder string `envconfig:"FOLDER" 	default:"./out/logs"`
}

type JWTConfig struct {
	Secret       string        `envconfig:"SECRET" required:"true"`
	AccessExpiry time.Duration `envconfig:"ACCESS_EXPIRY" default:"15m"`
}

func newConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("POSTGRES", &config.Database); err != nil {
		return Config{}, fmt.Errorf("process postgres config: %w", err)
	}

	if err := envconfig.Process("SERVER", &config.Server); err != nil {
		return Config{}, fmt.Errorf("process server config: %w", err)
	}

	if err := envconfig.Process("LOGGER", &config.Logger); err != nil {
		return Config{}, fmt.Errorf("process logger config: %w", err)
	}

	if err := envconfig.Process("JWT", &config.JWT); err != nil {
		return Config{}, fmt.Errorf("process JWT config: %w", err)
	}

	return config, nil
}

func MustLoad() *Config {
	config, err := newConfig()

	if err != nil {
		log.Fatalf("config: failed to load: %v", err)
	}

	return &config
}
