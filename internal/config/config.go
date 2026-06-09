package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Theme             string   `mapstructure:"theme" yaml:"theme"`
	PagerEnabled      bool     `mapstructure:"pager_enabled" yaml:"pager_enabled"`
	SensitivePatterns []string `mapstructure:"sensitive_patterns" yaml:"sensitive_patterns"`
}

func LoadConfig(cfgFile string) *Config {
	// Default configuration
	cfg := &Config{
		Theme:        "fruity",
		PagerEnabled: true,
		SensitivePatterns: []string{},
	}

	// Set up viper
	v := viper.New()
	
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		// Search for config in home directory
		home, err := os.UserHomeDir()
		if err == nil {
			v.AddConfigPath(home)
			v.SetConfigName(".config-read")
			v.SetConfigType("yaml")
		}
	}

	// Environment variables override
	v.SetEnvPrefix("CONFIG_READ")
	v.AutomaticEnv()

	// Read config file if it exists
	if err := v.ReadInConfig(); err == nil {
		if err := v.Unmarshal(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to parse config: %v\n", err)
		}
	}

	return cfg
}

func SetConfigValue(cfgFile string, key string, value string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot find home directory: %w", err)
	}

	configPath := cfgFile
	if configPath == "" {
		configPath = filepath.Join(home, ".config-read.yaml")
	}

	// Read existing config or create new
	cfg := make(map[string]interface{})
	data, err := os.ReadFile(configPath)
	if err == nil {
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return fmt.Errorf("failed to parse config: %w", err)
		}
	}

	// Set value
	cfg[key] = value

	// Write back
	out, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.WriteFile(configPath, out, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	fmt.Printf("Configuration updated: %s = %s\n", key, value)
	return nil
}