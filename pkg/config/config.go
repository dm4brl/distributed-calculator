package config

import (
	"fmt"
	"os"
	"errors"
	"encoding/json"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type Config struct {
	Server struct {
		Address string `mapstructure:"address"`
	} `mapstructure:"server"`

	Storage struct {
		Type string `mapstructure:"type"`
	} `mapstructure:"storage"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("distributed_calculator")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &cfg, nil
}

type Config struct {
	ServerAddress string
	AgentAddress  string
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: ":8080",
		AgentAddress:  ":8081",
	}
}

func (c *Config) Validate() error {
	if c.ServerAddress == "" {
		return fmt.Errorf("ServerAddress is required")
	}
	if c.AgentAddress == "" {
		return fmt.Errorf("AgentAddress is required")
	}
	return nil
}

func (c *Config) LoadFromEnv() {
	c.ServerAddress = os.Getenv("SERVER_ADDRESS")
	c.AgentAddress = os.Getenv("AGENT_ADDRESS")
}



import (
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

type Config struct {
	Server struct {
		Address string
	}
}

func NewConfig(address string) (*Config, error) {
	if address == "" {
		return nil, errors.New("address is required")
	}

	return &Config{
		Server: struct{ Address string }{
			Address: address,
		},
	}, nil
}


type Config struct {
	Server struct {
		Address string `json:"address"`
	} `json:"server"`
}

func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := &Config{}
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return cfg, nil
}
