package config

import (
	"encoding/json"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

const (
	dbPassEsqSeq = "{password}"
	password     = "calendar-service-password"
)

type IConfig interface {
	GetDbConfig() (*pgxpool.Config, error)
	GetHTTPAddress() string
	GetDataSource() *DataSource
	GetLoggerConfig() *LoggerConfig
}

type LoggerConfig struct {
	LoggerConfig *zerolog.Logger
	Level        string `json:"level"`
}

type DataSource struct {
	Repos string `json:"data_source"`
}

type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Config struct {
	DB         *DB           `json:"db"`
	HTTP       *HTTP         `json:"http"`
	Logger     *LoggerConfig `json:"logger"`
	DataSource *DataSource   `json:"repos"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetDbConfig() (*pgxpool.Config, error) {
	dbDsn := strings.ReplaceAll(c.DB.DSN, dbPassEsqSeq, password)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.DB.MaxOpenConnections

	return poolConfig, nil
}

func (c *Config) GetHTTPAddress() string {
	return net.JoinHostPort(c.HTTP.Host, c.HTTP.Port)
}

func (c *Config) GetLoggerConfig() *LoggerConfig {
	return c.Logger
}

func (c *Config) GetDataSource() *DataSource {
	return c.DataSource
}
