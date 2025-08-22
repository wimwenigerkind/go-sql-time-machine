/*
Copyright © 2025 Wim Wenigerkind <wenigerkind@heptacom.de>
*/
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MySQL MySQLConfig `yaml:"mysql"`
}
type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	ServerId uint32 `yaml:"server_id"`
}

func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := v.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
