package config

type Config struct {
	Port int
}

func New() *Config {
	config := &Config{}
	
	return config
}