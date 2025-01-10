package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name        string `mapstructure:"name"`
		Environment string `mapstructure:"environment"`
	} `mapstructure:"app"`

	Logging struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
		File   string `mapstructure:"file"`
	} `mapstructure:"logging"`

	Db struct {
		Uri               string `mapstructure:"uri"`
		Name              string `mapstructure:"database_name"`
		AuthCollection    string `mapstructure:"auth_collection"`
		ProfileCollection string `mapstructure:"profile_collection"`
	} `mapstructure:"db"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("../config/")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}

	return cfg
}
