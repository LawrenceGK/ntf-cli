package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	DefaultTopic   string `json:"default_topic"`
	DefaultMessage string `json:"default_message"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func SaveConfig(filePath string, config *Config) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(config)
}

// GetConfigPath 返回配置文件的完整路径
func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// 根据操作系统确定配置文件路径
	var configPath string
	if runtime.GOOS == "windows" {
		configPath = filepath.Join(homeDir, "AppData", "Local", "ntf-cli", "config.json")
	} else {
		configPath = filepath.Join(homeDir, ".config", "ntf-cli", "config.json")
	}

	// 确保配置文件目录存在
	configDir := filepath.Dir(configPath)
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return "", err
	}

	return configPath, nil
}
