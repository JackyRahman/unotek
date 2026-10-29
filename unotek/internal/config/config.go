package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Security SecurityConfig
}

type AppConfig struct {
	Name               string
	Env                string
	Port               string
	Debug              bool
	URL                string
	EnableAutoMigrate  bool
	EnableAutoSeeder   bool
}

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	SSLMode         string
	Timezone        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type JWTConfig struct {
	Secret                string
	ExpirationHours       int
	RefreshExpirationDays int
}

type SecurityConfig struct {
	RateLimitRPS   float64
	RateLimitBurst int
}

func LoadConfig() *Config {
	// Load .env file if it exists, otherwise use OS env
	if err := godotenv.Load(); err != nil {
		log.Println("[CONFIG] .env file not found, loading from system environment variables")
	}

	return &Config{
		App: AppConfig{
			Name:              getEnv("APP_NAME", "unotek-backend"),
			Env:               getEnv("APP_ENV", "development"),
			Port:              getEnv("APP_PORT", "8080"),
			Debug:             getEnvAsBool("APP_DEBUG", true),
			URL:               getEnv("APP_URL", "http://localhost:8080"),
			EnableAutoMigrate: getEnvAsBool("ENABLE_AUTO_MIGRATION", true),
			EnableAutoSeeder:  getEnvAsBool("ENABLE_AUTO_SEEDER", true),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "5432"),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", "postgres"),
			Name:            getEnv("DB_NAME", "unotek_company_db"),
			SSLMode:         getEnv("DB_SSLMODE", "disable"),
			Timezone:        getEnv("DB_TIMEZONE", "Asia/Jakarta"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 50),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
			ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", 60*time.Minute),
		},
		JWT: JWTConfig{
			Secret:                getEnv("JWT_SECRET", "super-secret-key-change-this-in-production-min-32-chars-long!"),
			ExpirationHours:       getEnvAsInt("JWT_EXPIRATION_HOURS", 24),
			RefreshExpirationDays: getEnvAsInt("JWT_REFRESH_EXPIRATION_DAYS", 7),
		},
		Security: SecurityConfig{
			RateLimitRPS:   getEnvAsFloat64("RATE_LIMIT_RPS", 20.0),
			RateLimitBurst: getEnvAsInt("RATE_LIMIT_BURST", 40),
		},
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return fallback
}

func getEnvAsFloat64(key string, fallback float64) float64 {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		if fVal, err := strconv.ParseFloat(val, 64); err == nil {
			return fVal
		}
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		if boolVal, err := strconv.ParseBool(val); err == nil {
			return boolVal
		}
	}
	return fallback
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		if duration, err := time.ParseDuration(val); err == nil {
			return duration
		}
	}
	return fallback
}
