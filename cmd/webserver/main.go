package main

import (
	"fmt"
	"github.com/alefwhite/api-users-go/internal/handler"
	"github.com/alefwhite/api-users-go/internal/repository/categoryrepository"
	"github.com/alefwhite/api-users-go/internal/repository/productrepository"
	"github.com/alefwhite/api-users-go/internal/service/categoryservice"
	"github.com/alefwhite/api-users-go/internal/service/productservice"
	"log/slog"
	"net/http"

	"github.com/alefwhite/api-users-go/config/env"
	"github.com/alefwhite/api-users-go/config/logger"
	_ "github.com/alefwhite/api-users-go/docs"
	"github.com/alefwhite/api-users-go/internal/database"
	"github.com/alefwhite/api-users-go/internal/database/sqlc"
	"github.com/alefwhite/api-users-go/internal/handler/routes"
	"github.com/alefwhite/api-users-go/internal/repository/userrepository"
	"github.com/alefwhite/api-users-go/internal/service/userservice"
	"github.com/go-chi/chi/v5"
)

func main() {
	logger.InitLogger()
	slog.Info("starting api")

	// VsCode env.LoadingConfig("../../")
	// Goland env.LoadingConfig(".")
	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load environment variables", err, slog.String("package", "main"))
		return
	}

	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	queries := sqlc.New(dbConnection)

	// user
	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)

	// category
	categoryRepo := categoryrepository.NewCategoryRepository(dbConnection, queries)
	newCategoryService := categoryservice.NewCategoryService(categoryRepo)

	// product
	productRepo := productrepository.NewProductRepository(dbConnection, queries)
	productService := productservice.NewProductService(productRepo)

	newHandler := handler.NewHandler(newUserService, newCategoryService, productService)

	// init routes
	router := chi.NewRouter()
	routes.InitRoutes(router, newHandler)
	routes.InitDocsRoutes(router)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))

	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}

}
