package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASS"`
	DBName     string `mapstructure:"DB_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	ServerHost string `mapstructure:"SERVER_HOST"`
}

type XenditKey struct {
	WriteKey string `mapstructure:"WRITE_KEY_XENDIT"`
	ReadKey  string `mapstructure:"READ_KEY_XENDIT"`
}

func LoadXenditKey(path string) (key XenditKey, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return XenditKey{}, err
	}

	err = viper.Unmarshal(&key)

	return key, err
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)

	return config, err
}
