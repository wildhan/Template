package main

import (
	"fmt"
	"net/http"
	"os"

	"template/config/database"
	"template/lib/log"

	userHandler "template/package/user/handler"
	userRepository "template/package/user/repository"
	userUsecase "template/package/user/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Error(fmt.Sprintf("Failed load .env: %v", err.Error()))
		os.Exit(2)
	}

	dbConn := database.CreateConnection()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Template")
	})

	userRepo := userRepository.NewUserRepo(dbConn)
	userUC := userUsecase.NewUserUsecase(userRepo)
	userHandler.NewUserHandler(userUC).Mount(e.Group("/user"))

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Error(fmt.Sprintf("Failed start echo: %v", err.Error()))
	}
}
