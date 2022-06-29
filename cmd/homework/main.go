package main

import (
	"homework/internal/application"
	"homework/internal/configs"
)

func main() {
	//разобратся как прокидывать конфиги из ямл
	// как вариант viper
	// енвы и все такое
	// тут же секреты  итд
	cfg := configs.MainConfig{
		configs.DBConfig{
			Dialect:      "postgres",
			Dsn:          "postgresql://ps_report_dev:DbrjoE95x6SrvIxQ@10.48.3.204:5432/ps_report_dev?sslmode=disable",
			DisableLog:   false,
			MaxIdleConn:  10,
			MaxOpenConn:  10,
			ConnLifetime: "1h",
		},
	}
	app := application.NewApp(&cfg)

	//error or panic?
	app.Run()
}
