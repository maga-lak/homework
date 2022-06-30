package main

import (
	"homework/internal/application"
	"homework/internal/configs"
)

func main() {
	//разобратся как прокидывать конфиги из ямл
	// @todo viper
	// енвы и все такое
	// тут же секреты  итд
	cfg := configs.MainConfig{
		Database: configs.DBConfig{
			Dsn: "postgresql://db:testPassQ@192.168.0.1:5432/db?sslmode=disable",
		},
		Auth: configs.AuthConfig{
			SignKey:        []byte("some_key"),
			Alg:            "RS256",
			ExpireDuration: 600,
		},
		Port: 8081,
	}
	app := application.NewApp(&cfg)

	//error or panic?
	app.Run()
}
