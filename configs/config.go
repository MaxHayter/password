package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server *Server `yaml:"server"`
}

type Server struct {
	GrpcAddress string `yaml:"grpc_address"`
}

func Configure(fileName string) (*Config, error) {
	var cnf *Config
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cnf)
	if err != nil {
		return nil, err
	}

	return cnf, nil
}
