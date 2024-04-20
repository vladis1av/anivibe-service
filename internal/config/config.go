package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const defaultConfigPath = ".env"

type Config struct {
	HTTPAddr       string        `env:"HTTP_ADDR" env-default:":80"`
	ReadTimeout    time.Duration `env:"READ_TIMEOUT" env-default:"60s"`
	WriteTimeout   time.Duration `env:"WRITE_TIMEOUT" env-default:"60s"`
	IdleTimeout    time.Duration `env:"IDLE_TIMEOUT" env-default:"60s"`
	MaxHeaderBytes int           `env:"MAX_HEADER_BYTES" env-default:"1048576"`
	AllowedOrigins []string      `env:"ALLOWED_ORIGINS" env-separator:"," env-description:"List of allowed sources for CORS"`
}

// LoadConfig загружает конфигурацию из переменных окружения или файла конфигурации.
func LoadConfig() *Config {
	cfg := &Config{}

	// Проверяем значение переменной окружения AMVERA
	amveraVar, exists := os.LookupEnv("AMVERA")

	if exists && amveraVar == "1" {
		// Загрузка конфигурации из переменных окружения (облако Amvera)
		if err := cleanenv.ReadEnv(cfg); err != nil {
			log.Fatalf("Failed to load configuration from environment variables: %v", err)
		}
	} else {
		// Загрузка конфигурации из файла конфигурации и переменных окружения
		configPath := getConfigPath()
		if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
			log.Fatalf("Failed to load configuration from file %s: %v", configPath, err)
		}

		// Загрузка конфигурации из переменных окружения
		if err := cleanenv.ReadEnv(cfg); err != nil {
			log.Fatalf("Failed to load configuration from environment variables: %v", err)
		}
	}

	return cfg
}

// getConfigPath возвращает путь к файлу конфигурации из переменной окружения CONFIG_PATH
// или возвращает путь по умолчанию, если переменная окружения не задана.
func getConfigPath() string {
	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		configPath = defaultConfigPath
	}
	return configPath
}
