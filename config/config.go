package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func ReadConfig() (*Config, error){
	viper.SetConfigName("config")
	
	// by default config is on ~/.ai-commitizen/config.yaml
	viper.AddConfigPath("$HOME/.ai-commitizen")

	// read config
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	// if not found, return empty config
	if _, err := os.Stat(viper.ConfigFileUsed()); os.IsNotExist(err) {
		// show message that config file is not found, please update it
		log.Printf("config file %s is not found, please update it", viper.ConfigFileUsed())

		return &Config{}, nil
	}

	// unmarshal config
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}