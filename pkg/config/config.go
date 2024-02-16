package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
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
