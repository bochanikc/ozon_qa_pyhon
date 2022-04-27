//go:build integration
// +build integration

package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host string `envconfig:"GRPC_HOST" default:"localhost:50051"`
}

func ProcessConfig() Config {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatalf("envconfig.Process() err: %v", err)
	}
	return c
}
