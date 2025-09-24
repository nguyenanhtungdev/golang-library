package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`

	ServerPort string `mapstructure:"SERVER_PORT"`
	GinMode    string `mapstructure:"GIN_MODE"`

	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiration string `mapstructure:"JWT_EXPIRATION"`

	LogLevel    string `mapstructure:"LOG_LEVEL"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func Load() *Config {
	config := &Config{}

	viper.AutomaticEnv()

	if os.Getenv("RENDER") == "" {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		_ = viper.ReadInConfig()
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatal().Err(err).Msg("Failed to unmarshal config")
	}

	return config
}
