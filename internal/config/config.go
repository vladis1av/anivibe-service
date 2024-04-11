package config

import (
	"github.com/joho/godotenv"
)

type HTTPConfig interface {
	Address() string
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
