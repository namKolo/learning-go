package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Address  string
	Database string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Address:  "localhost:28015",
			Database: "test",
		},
	}
}
