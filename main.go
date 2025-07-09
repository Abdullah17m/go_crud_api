package main

import (
	"crud-api/config"
	"crud-api/controller"
	"crud-api/database"
	"crud-api/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo) {
    e.POST("/user", controller.CreateUser)
    e.GET("/users/:id", controller.GetUser)
	e.GET("/users", controller.GetAllUsers)
    e.PUT("/users/:id", controller.UpdateUser)
    e.DELETE("/users/:id", controller.DeleteUser)
}

func main() {
	config.LoadConfig()
	logger.InitLogger()
	database.ConnectMongo()

	e := echo.New()
	e.Use(middleware.Logger())

	RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.AppPort))
}
