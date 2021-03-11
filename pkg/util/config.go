package util

import "github.com/spf13/viper"

// Config holds all the configuration settings
// of the application which are read by viper
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	Addr     string `mapstructure:"ADDR"`
	YtKey    string `mapstructure:"YT_KEY"`
	IgKey    string `mapstructure:"IG_KEY"`
	TtKey    string `mapstructure:"TT_KEY"`
}

// LoadConfig reads a config file from the path provided
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
