package config

import "time"

type GrpcServer struct {
	Host              string        `yaml:"host"`
	Port              string        `yaml:"port"`
	MaxConnectionIdle time.Duration `yaml:"max_connection_idle"`
	MaxConnectionAge  time.Duration `yaml:"max_connection_age"`
	Timeout           time.Duration `yaml:"timeout"`
	Time              time.Duration `yaml:"time"`
}

type Graceful struct {
	Timeout time.Duration `yaml:"timeout"`
}
