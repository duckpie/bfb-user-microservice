package config

type Config struct {
	Services ServicesConfigs `toml:"services"`
}

type ServicesConfigs struct {
	Server ServerConfig   `toml:"server"`
	DB     DatabaseConfig `toml:"database"`
	Redis  RedisConfig    `toml:"redis"`
}

type ServerConfig struct {
	Port int64 `toml:"PORT"`
}

type DatabaseConfig struct {
	DbUrl string `toml:"DB_URL"`
}

type RedisConfig struct {
	Host string `toml:"HOST"`
	Port int64  `toml:"PORT"`
	DB   int8   `toml:"DB"`
}

func NewConfig() *Config {
	return &Config{}
}
