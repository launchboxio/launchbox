package server

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

// TODO: Support loading a bunch of configuration from
// environment variables with the appropriate name
const (
	DatabaseDsn          = "DATABASE_DSN"
	RedisUrl             = "REDIS_URL"
	RedisPassword        = "REDIS_PASSWORD"
	CentrifugoUrl        = "CENTRIFUGO_URL"
	PrometheusUrl        = "PROMETHEUS_URL"
	LokiUrl              = "LOKI_URL"
	VaultUrl             = "VAULT_URL"
	HTTPClientCert       = "LAUNCHBOX_CLIENT_CERT"
	HTTPClientKey        = "LAUNCHBOX_CLIENT_KEY"
	HTTPSSLVerifyEnvName = "LAUNCHBOX_HTTP_SSL_VERIFY"
)

type CorsConfig struct {
	AllowOrigins     []string      `yaml:"allowed_origins,omitempty"`
	AllowMethods     []string      `yaml:"allowed_methods,omitempty"`
	AllowHeaders     []string      `yaml:"allowed_headers,omitempty"`
	ExposeHeaders    []string      `yaml:"expose_headers,omitempty"`
	AllowCredentials bool          `yaml:"allow_credentials"`
	MaxAge           time.Duration `yaml:"max_age"`
}

type DatabaseConfig struct {
	Dsn string `yaml:"dsn,omitempty"`
}

type RedisConfig struct {
	Url      string `yaml:"url,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type CentrifugoConfig struct {
	Enabled bool   `yaml:"enabled"`
	Url     string `yaml:"url"`
}

type PrometheusConfig struct {
	Enabled bool   `yaml:"enabled"`
	Url     string `yaml:"url"`
}

type LokiConfig struct {
	Enabled bool   `yaml:"enabled"`
	Url     string `yaml:"url"`
}

type TaskWorker struct {
	ConsumerTag string `yaml:"consumer_tag"`
}

type Config struct {
	Database   DatabaseConfig   `yaml:"database,omitempty"`
	Cors       CorsConfig       `yaml:"cors,omitempty"`
	Redis      RedisConfig      `yaml:"redis,omitempty"`
	Centrifugo CentrifugoConfig `yaml:"centrifugo,omitempty"`
	Prometheus PrometheusConfig `yaml:"prometheus,omitempty"`
	Loki       LokiConfig       `yaml:"loki,omitempty"`
	Worker     TaskWorker       `yaml:"worker"`
}

func LoadDefaultConfig(filePath string) (*Config, error) {
	config := &Config{}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return nil, err
	}

	return config, nil
}
