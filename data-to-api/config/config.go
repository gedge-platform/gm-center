package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "192.168.150.102",
			Port:     12003,
			Username: "username",
			Password: "password",
			Name:     "gedge",
			Charset:  "utf8",
		},
	}
}
