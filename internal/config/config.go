package config

type jwtConfig struct {
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
}
type databaseConfig struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     uint16 `env:"DB_PORT"`
}
type config struct {
	Jwt jwtConfig
	Db  databaseConfig
}

var Config config
