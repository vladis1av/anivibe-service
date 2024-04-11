package env

import (
	"anivibe-service/internal/config"
	"os"

	"github.com/pkg/errors"
)

var _ config.HTTPConfig = (*httpConfig)(nil)

const (
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	port string
}

func NewHTTPConfig() (*httpConfig, error) {
	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &httpConfig{
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return ":" + cfg.port
}
