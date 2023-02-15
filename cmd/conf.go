package cmd

import "github.com/spf13/viper"

var conf = struct {
	Debug      bool   `mapstructure:"DEBUG"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	HttpListen string `mapstructure:"HTTP_LISTEN"`
}{}

func confLoad() {
	viper.SetDefault("HTTP_LISTEN", ":8080")

	viper.SetConfigFile("conf.yml")
	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&conf)
}
