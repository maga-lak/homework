package application

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"homework/internal/configs"
	v1 "homework/internal/controller/http/v1"
	"homework/internal/infrastructure/repository/postgre"
	"homework/internal/service"
	"net/http"
)

type Application struct {
	cfg        *configs.MainConfig
	httpServer *http.Server
}

func NewApp(cfg *configs.MainConfig) *Application {
	return &Application{
		cfg: cfg,
	}
}

func (a *Application) Halt() {

}

func (a *Application) Shutdown() {

}

func (a *Application) Run() {
	sugar := initLogger()
	sugar.Info("start app")

	conn, err := postgre.GetConnection(a.cfg.Database.Dsn, "ps-report-go")
	if err != nil {
		sugar.Fatalf("db connection error: %s", err.Error())
		panic(err)
	}
	sugar.Info("init db")

	userRepo := postgre.NewUserRepository(conn)

	authService := service.NewAuthService(userRepo, a.cfg.Auth)

	a.initServer(sugar, authService)
}

func (a *Application) initServer(sugar *zap.SugaredLogger, authService *service.AuthService) {
	handler := v1.NewServer(authService, sugar)

	router := chi.NewMux()
	router.Post("/api/v1/auth", handler.Auth)

	a.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: router,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			sugar.Fatalf("Failed to start http server:  %s", err.Error())
			panic(err)
		}
	}()
}

func initLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}
