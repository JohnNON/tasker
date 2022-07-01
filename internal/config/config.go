package config

// Config - содержит конфигурацию для запуска сервера
type Config struct {
	BindAddr   string
	EndPoint   string
	DBURL      string
	DBName     string
	Collection string
}

// NewConfig - инициализация конфига по умолчанию
func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		EndPoint:   "/api/v1",
		DBURL:      "mongodb://docker:mongopw@localhost:49153",
		DBName:     "tasks",
		Collection: "tasks",
	}
}
