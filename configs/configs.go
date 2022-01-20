package configs

import (
	"time"

	"github.com/spf13/viper"
)

var (
	DbConfig DBConfig
	RdConfig RedisConfig
)

type DBConfig struct {
	Host        string
	Port        string
	User        string
	Pass        string
	Dbname      string
	MaxPool     int
	IdlePool    int
	MaxLifetime time.Duration
}

type RedisConfig struct {
	Host string
	Port string
	Pass string
}

func initializeConfigProp() {
	DbConfig.Host = viper.GetString("db.host")
	DbConfig.Port = viper.GetString("db.port")
	DbConfig.User = viper.GetString("db.user")
	DbConfig.Pass = viper.GetString("db.pass")
	DbConfig.Dbname = viper.GetString("db.dbname")
	DbConfig.MaxPool = viper.GetInt("db.pool.max")
	DbConfig.IdlePool = viper.GetInt("db.pool.idle")
	DbConfig.MaxLifetime = viper.GetDuration("db.pool.lifetime")
	RdConfig.Host = viper.GetString("")
	RdConfig.Port = viper.GetString("")
	RdConfig.Pass = viper.GetString("")
}

func init() {
	initializeConfigProp()
}
