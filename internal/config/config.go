package config

type Config struct {
	Env        string           `koanf:"env"`
	HTTPServer HTTPServerConfig `koanf:"http_server"`
	GRPCServer GRPCServerConfig `koanf:"grpc_server"`
	Repository RepositoryConfig `koanf:"repository"`
}

type HTTPServerConfig struct {
	Port string `koanf:"port"`
}

type GRPCServerConfig struct {
	Port string `koanf:"port"`
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
