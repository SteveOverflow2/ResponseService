package config

import (
	"log"
	"strconv"
	"time"
)

type MySQLConfig struct {
	Database string `env:"MYSQL_DATABASE"`
	Username string `env:"MYSQL_USERNAME"`
	Password string `env:"MYSQL_PASSWORD"`
	Port     int    `env:"MYSQL_PORT"`
	Host     string `env:"MYSQL_HOST"`

	ConnMaxLifeTime time.Duration `env:"CONN_MAX_LIFE_TIME"`
	MaxIdleConns    int           `env:"MAX_IDLE_CONNS"`
	MaxOpenConns    int           `env:"MAX_OPEN_CONNS"`
}

const mysqlLog string = "[MYSQL - Config]: "

func (cfg *Config) loadMySQLConfig() MySQLConfig {
	envFields := cfg.loadEnvFields(MySQLConfig{})
	const invalidErr = "%s is not a valid entry"

	connMaxLifeTime, err := time.ParseDuration(envFields["ConnMaxLifeTime"])
	if err != nil {
		log.Fatalf(mysqlLog+invalidErr, envFields["ConnMaxLifeTime"])
	}
	MaxIdleConns, err := strconv.Atoi(envFields["MaxIdleConns"])
	if err != nil {
		log.Fatalf(mysqlLog+invalidErr, envFields["MaxIdleConns"])
	}
	MaxOpenConns, err := strconv.Atoi(envFields["MaxOpenConns"])
	if err != nil {
		log.Fatalf(mysqlLog+invalidErr, envFields["MaxOpenConns"])
	}

	port, err := strconv.Atoi(envFields["Port"])
	if err != nil {
		log.Fatalf(mysqlLog+invalidErr, envFields["Port"])
	}

	return MySQLConfig{
		Database:        envFields["Database"],
		Username:        envFields["Username"],
		Password:        envFields["Password"],
		Port:            port,
		Host:            envFields["Host"],
		ConnMaxLifeTime: connMaxLifeTime,
		MaxIdleConns:    MaxIdleConns,
		MaxOpenConns:    MaxOpenConns,
	}
}
