package database

import (
	"fmt"
	"os"

	"github.com/nguyenanhtungdev/golang-library/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Option func(*gorm.Config)

func WithLogger(l logger.Interface) Option {
	return func(cfg *gorm.Config) {
		cfg.Logger = l
	}
}

func ConnectDatabase(cfg *config.Config, opts ...Option) (*gorm.DB, error) {
	if envURL := os.Getenv("DATABASE_URL"); envURL != "" {
		cfg.DatabaseURL = envURL
	}

	gormConfig := &gorm.Config{}

	if cfg.LogLevel == "debug" && gormConfig.Logger == nil {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else if gormConfig.Logger == nil {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	for _, opt := range opts {
		opt(gormConfig)
	}

	log.Debug().Msg("Connecting to: " + cfg.DatabaseURL)

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
