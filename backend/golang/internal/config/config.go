package config

type Config struct {
	Port string `env:"PORT" env-default:"8080"`
	Env  string `env:"ENV" env-default:"development"`
}
