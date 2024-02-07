package config

import (
	"time"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
)

type Configs struct {
	App *AppConfig `mapstructure:"app"`
	DB  *DBConf    `mapstructure:"db"`
}

type AppConfig struct {
	TimeOut time.Duration `mapstructure:"timeout"`
	Port    int           `mapstructure:"port"`
}

type DBConf struct {
	Host     string        `mapstructure:"host"`
	Port     int           `mapstructure:"port"`
	Username string        `mapstructure:"username"`
	Password string        `mapstructure:"password"`
	DBName   string        `mapstructure:"dbname"`
	SSLMode  string        `mapstructure:"sslmode"`
	TimeOut  time.Duration `mapstructure:"timeout"`
}

func New() (*Configs, error) {
	configFile := "config.yaml"
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Configs{}

	if err := defaults.Set(cfg); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
