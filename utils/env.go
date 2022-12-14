package utils

import "github.com/spf13/viper"

func ReadENV(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return viper.GetString(key)
}