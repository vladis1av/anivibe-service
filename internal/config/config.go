package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPAddr       string        `env:"HTTP_ADDR" env-default:":8080"`
	ReadTimeout    time.Duration `env:"READ_TIMEOUT" env-default:"5s"`
	WriteTimeout   time.Duration `env:"WRITE_TIMEOUT" env-default:"10s"`
	IdleTimeout    time.Duration `env:"IDLE_TIMEOUT" env-default:"60s"`
	MaxHeaderBytes int           `env:"MAX_HEADER_BYTES" env-default:"1048576"`

	ProxyHTTPAddr       string        `env:"PROXY_HTTP_ADDR" env-default:":8081"`
	ProxyReadTimeout    time.Duration `env:"PROXY_READ_TIMEOUT" env-default:"30s"`
	ProxyWriteTimeout   time.Duration `env:"PROXY_WRITE_TIMEOUT" env-default:"30s"`
	ProxyIdleTimeout    time.Duration `env:"PROXY_IDLE_TIMEOUT" env-default:"60s"`
	ProxyMaxHeaderBytes int           `env:"PROXY_MAX_HEADER_BYTES" env-default:"1048576"`
}

var (
	instance *Config
	once     sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		cfg := &Config{}
		err := cleanenv.ReadEnv(cfg)
		if err != nil {
			panic(err)
		}
		instance = cfg
	})
	return instance
}
