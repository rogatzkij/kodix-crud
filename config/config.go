package config

import "github.com/kelseyhightower/envconfig"

// Настройки сервиса
type Config struct {
	Port  int    `default:"8080"`                      // порт, на котором будет запущен сервис
	Mongo string `default:"mongodb://127.0.0.1:27017"` // адрес монги
}

// Считываем настройки из переменных окружения
func GetConfigFromEnv() (*Config, error) {
	conf := &Config{}
	err := envconfig.Process("KODIX", conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
