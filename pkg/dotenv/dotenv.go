package dotenv

import "github.com/spf13/viper"

func Viper() error {
	// viper.SetConfigFile(".env")
	viper.SetConfigFile("/app/.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	return err
}
