package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    ServerPort  int    `mapstructure:"SERVER_PORT"`
    DBHost      string `mapstructure:"DB_HOST"`
    DBPort      int    `mapstructure:"DB_PORT"`
    DBUser      string `mapstructure:"DB_USER"`
    DBPassword  string `mapstructure:"DB_PASSWORD"`
    DBName      string `mapstructure:"DB_NAME"`
}

func LoadConfig() (Config, error) {
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return Config{}, err
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return Config{}, err
    }

    return config, nil
}