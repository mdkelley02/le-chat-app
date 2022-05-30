package configs

type Settings struct {
	AppRole          string `env:"APP_ROLE,default=user-service"`
	LogLevel         string `env:"LOG_LEVEL,default=info"`
	Host             string `env:"HOST,default=localhost"`
	Port             string `env:"PORT,default=8080"`
	PostgresHost     string `env:"POSTGRES_HOST,default=localhost"`
	PostgresPort     string `env:"POSTGRES_PORT,default=5432"`
	PostgresUser     string `env:"POSTGRES_USER,default=postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,default=postgres"`
	PostgresDB       string `env:"POSTGRES_DB,default=postgres"`
}
