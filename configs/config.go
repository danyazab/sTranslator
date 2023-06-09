package configs

type Config struct {
	BindAddr string `env:"BIND_ADDR"`
	LogLevel string `env:"LOG_LEVEL"`

	Url      string `env:"URL"`
	BotToken string `env:"BOT_TOKEN"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8088",
		LogLevel: "debug",

		Url:      "6017029566:AAGBxkHgOY3CQc8Q2QT6LJgFCy7XC0Qx2Os",
		BotToken: "6005524984:AAG3gykz5gU75bhwRF1_cdm9iQJgB039FEo",
	}
}
