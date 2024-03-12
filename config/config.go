package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type (
	Config struct {
		App      App `yaml:"app"`
		Database `yaml:"database"`
		SMTP
		MinIO MinIOCredentials `yaml:"minio"`
	}
	Database struct {
		DSN string `yaml:"dsn"`
	}
	App struct {
		Name   string `yaml:"name"`
		HTTP   HTTP   `yaml:"http"`
		Server Server `yaml:"server"`
		Auth   Auth   `yaml:"auth"`
	}
	SMTP struct {
		From     string `yaml:"from"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	}
	HTTP struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	Auth struct {
		SecretKey string `yaml:"secret_key"`
	}
	Server struct {
		CORS CORS `yaml:"cors"`
	}
	CORS struct {
		AllowOrigins string `yaml:"allow_origins"`
	}
	MinIOCredentials struct {
		Endpoint  string `yaml:"endpoint"`
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
	}
)

var ProjectRootPath = ConfigsDirPath() + "/../"

func ConfigsDirPath() string {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error in generating env dir")
	}

	return filepath.Dir(f)
}

func Read() *Config {
	// Load YAML configs
	var config *Config

	data, readErr := os.ReadFile(ConfigsDirPath() + "/configs.yml")
	if readErr != nil {
		panic(readErr)
	}

	parseErr := yaml.Unmarshal(data, &config)
	if parseErr != nil {
		panic(parseErr)
	}

	// Load JwtSecret keys
	secretData, secretErr := os.ReadFile(ConfigsDirPath() + "/jwt_secret.pem")
	if secretErr != nil {
		panic(secretErr)
	}

	config.App.Auth.SecretKey = strings.TrimSpace(string(secretData))

	return config
}
