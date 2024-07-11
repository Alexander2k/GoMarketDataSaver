package config

import (
	"github.com/spf13/viper"
	"log"
)

type (
	Config struct {
		PDB   PostgresDB  `mapstructure:"db"`
		Bybit BybitConfig `mapstructure:"bybit"`
	}

	PostgresDB struct {
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
	viper.AddConfigPath("./config/")
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
