package main

import (
	"backend/internal/handler"
	"backend/internal/handler/middleware"
	"backend/internal/handler/router"
	"backend/internal/infra/ent"
	"backend/internal/infra/repository"
	"backend/internal/usecase"
	"backend/pkg/config"
	"log/slog"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func setupHandler(cfg *config.Config) http.Handler {
	client, err := ent.Open("mysql", cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed opening connection to mysql", "error", err)
	}

	userRepo := repository.NewUserRepository(client)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	fileHandler := handler.NewFileHandler()

	swaggerHandler := handler.NewSwaggerHandler()

	mux := router.SetupRoutes(cfg, &router.Handlers{
		UserHandler:    userHandler,
		FileHandler:    fileHandler,
		SwaggerHandler: swaggerHandler,
	})

	handler := middleware.CORS(mux)

	if cfg.Env == "production" {
		handler = middleware.BasicAuth(cfg.BasicAuthUser, cfg.BasicAuthPass, handler)
	}

	return handler
}
