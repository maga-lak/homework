package application

import (
	"homework/internal/service"

	"homework/internal/configs"
	"homework/internal/infrastructure/repository/postgre"
)

type Application struct {
	cfg *configs.MainConfig
}

func NewApp(cfg *configs.MainConfig) *Application {
	return &Application{
		cfg,
	}
}

func (a *Application) Halt() {

}

func (a *Application) Shutdown() {

}

func (a *Application) Run() {
	conn, err := postgre.GetConnection(a.cfg.Database.Dsn, "ps-report-go")
	if err != nil {
		panic(err)
	}

	userRepo := postgre.NewUserRepository(conn)

	authService := service.NewAuthService(userRepo)

	//
	//cnt := &container.AppContainer{
	//	ReportSrv: reportSrv,
	//}
	//srv := server.NewServers(sugar, cnt, a.cfg.Report.Port)
	//err = srv.Run()
	//if err != nil {
	//	panic(err)
	//}
}
