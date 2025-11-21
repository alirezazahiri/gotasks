package config

const EnvPrefix = "GOTASKS_"

var defaultConfig = map[string]any{
	"env": "development",
	"repository.postgres.username": "",
	"repository.postgres.password": "",
	"repository.postgres.host": "",
	"repository.postgres.port": "",
	"repository.postgres.dbname": "",
}
