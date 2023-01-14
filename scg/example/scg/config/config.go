// Code generated by scg 1, .
//
// EchoGame-Go config file
//
// Command:
// scg
//.

package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/ThCompiler/go_game_constractor/scg/example/scg/pkg/logger/prepare"
)

type (
	// Config -.
	Config struct {
		App   App            `yaml:"app"`
		HTTP  HTTP           `yaml:"http"`
		Log   prepare.Config `yaml:"logger"`
		Redis Redis          `yaml:"redis"`
	}

	// App -.
	App struct {
		Name         string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version      string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		ResourcesDir string `env-required:"true" yaml:"resources_dir" env:"RESOURCES_DIR"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Redis -.
	Redis struct {
		URL string `env-required:"true" yaml:"url,omitempty" env:"REDIS_URL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./scg/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
