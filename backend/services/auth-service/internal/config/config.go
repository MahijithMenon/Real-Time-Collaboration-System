package config

import "os"

type Config struct {
	Port  string
	DBUrl string
}

func Load(defaultPort string) Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		// Temporary local default. Change later.
		dbUrl = "postgres://postgres:postgres@localhost:5432/auth_service?sslmode=disable"
	}

	return Config{
		Port:  port,
		DBUrl: dbUrl,
	}
}
