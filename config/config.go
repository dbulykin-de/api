package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	GrpcServer GrpcServer `yaml:"grpc_server"`
	Graceful   Graceful   `yaml:"graceful"`
}

func Instance() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config/config.yaml", instance); err != nil {
			log.Fatalf("read config error: %s", err.Error())
		}
		if err := cleanenv.ReadEnv(instance); err != nil {
			log.Fatalf("read env error: %s", err.Error())
		}
	})
	return instance
}
