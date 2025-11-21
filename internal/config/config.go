package config

type Config struct {
	Env        string           `koanf:"env"`
	Repository RepositoryConfig `koanf:"repository"`
}

type PostgresConfig struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	DBName   string `koanf:"dbname"`
}

type RepositoryConfig struct {
	Postgres PostgresConfig `koanf:"postgres"`
}
