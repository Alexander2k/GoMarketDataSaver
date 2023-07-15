package config

import "log"
import "github.com/spf13/viper"

type (
	Config struct {
		Db    DbConfig    `mapstructure:"db"`
		Bybit BybitConfig `mapstructure:"bybit"`
	}

	DbConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"db_name"`
		SSLMode  string `mapstructure:"ssl"`
	}

	BybitConfig struct {
		UrlPerpetual   string   `mapstructure:"url_perpetual"`
		UrlSpot        string   `mapstructure:"url_spot"`
		SpotPairs      []string `mapstructure:"spot_pairs"`
		PerpetualPairs []string `mapstructure:"perpetual_pairs"`
		PingTimeout    int64    `mapstructure:"ping_timeout"`
	}
)

func NewConfig() (c *Config, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./app/config/")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Cant read config %s", err)
		return
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Printf("Cant unmarshal config %s", err)
		return
	}
	return
}
