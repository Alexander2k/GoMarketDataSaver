package config

import (
	"github.com/spf13/viper"
	"log"
)

type (
	Config struct {
		PDB     PostgresDB   `mapstructure:"db"`
		Bybit   BybitConfig  `mapstructure:"bybit"`
		ClickDB ClickHouseDB `mapstructure:"clickDB"`
	}

	PostgresDB struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"db_name"`
		SSLMode  string `mapstructure:"ssl"`
	}

	ClickHouseDB struct {
		Host                 string `mapstructure:"host"`
		Port                 string `mapstructure:"port"`
		Database             string `mapstructure:"database"`
		Username             string `mapstructure:"username"`
		Password             string `mapstructure:"password"`
		DialTimeout          int    `mapstructure:"dial_timeout"`
		MaxOpenConns         int    `mapstructure:"max_open_conns"`
		MaxIdleConns         int    `mapstructure:"max_idle"`
		BlockBufferSize      int    `mapstructure:"block_buffer_size"`
		MaxCompressionBuffer int    `mapstructure:"max_compression_buffer"`
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
