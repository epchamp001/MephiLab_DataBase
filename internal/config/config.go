package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// DataGenerationConfig структура для количества генерируемых данных
type DataGenerationConfig struct {
	Chat         int `yaml:"num_chats"`
	Clients      int `yaml:"num_clients"`
	Orders       int `yaml:"num_orders"`
	Messages     int `yaml:"num_messages"`
	PromoCodes   int `yaml:"num_promo_codes"`
	Courier      int `yaml:"num_couriers"`
	SupportStaff int `yaml:"num_support_staffs"`
	Rate         int `yaml:"num_rates"`
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
