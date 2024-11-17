package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// DataGenerationConfig структура для количества генерируемых данных
type DataGenerationConfig struct {
	Chat         int `yaml:"num_chat"`
	Clients      int `yaml:"num_clients"`
	Orders       int `yaml:"num_orders"`
	Messages     int `yaml:"num_messages"`
	Order        int `yaml:"num_order"`
	Courier      int `yaml:"num_courier"`
	SupportStaff int `yaml:"num_support_staff"`
	Rate         int `yaml:"num_rate"`
}

// Config структура для конфигурационного файла
type Config struct {
	Database struct {
		URL string `yaml:"url"`
	} `yaml:"database"`
	DataGeneration DataGenerationConfig `yaml:"data_generation"`
}

// LoadConfig читает и парсит конфигурационный файл
func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
