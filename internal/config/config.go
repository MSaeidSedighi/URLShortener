package config

import "github.com/spf13/viper"

type Config struct {
	// manual assignment not needed if names match completely
	Port      string `mapstructure:"PORT"`
	DbUrl     string `mapstructure:"DB_URL"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
