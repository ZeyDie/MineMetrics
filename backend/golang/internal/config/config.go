package config

type Config struct {
	Port string `env:"PORT" env-default:"8080"`
}

type SQLConfig struct {
	Driver string `env:"SQL_DRIVER" env-default:"mysql"`

	Host     string `env:"SQL_HOST" env-default:"localhost"`
	Port     string `env:"SQL_PORT" env-default:"3306"`
	DBName   string `env:"SQL_DBNAME" env-default:"database"`
	User     string `env:"SQL_USER" env-default:"database_user"`
	Password string `env:"SQL_PASSWORD" env-default:"*********"`
	SSL      string `env:"SQL_SSL" env-default:"disable"`
	TimeZone string `env:"SQL_TIMEZONE" env-default:"UTC"`

	MaxOpenConns    int    `env:"SQL_MAX_OPEN_CONNS" env-default:"10"`
	MaxIdleConns    int    `env:"SQL_MAX_IDLE_CONNS" env-default:"5"`
	MaxConnLifetime string `env:"SQL_MAX_CONN_LIFETIME" env-default:"30m"`
}
