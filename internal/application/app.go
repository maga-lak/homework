package application

import (
	"go.uber.org/zap"

	"homework/internal/configs"
	"homework/internal/datasource/postgre"
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
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()

	conn, err := postgre.GetConnection(a.cfg.Database.Dsn, "ps-report-go")
	if err != nil {
		panic(err)
	}

	userRepo := postgre.NewUserRepository(conn)
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
