package api

import (
	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/db"
	"github.com/s0h1s2/invoice-app/internal/handlers"
	"github.com/s0h1s2/invoice-app/internal/services"
)

type engine struct {
	engine *gin.Engine
}

func NewEngine() *engine {
	eng := gin.Default()
	return &engine{
		engine: eng,
	}
}
func (e *engine) Start() {
	api := e.engine.Group("/api/v1")
	mysqlStore := db.NewMysqlStore("mydb_bowmeallog:ec10d2384b9154cdf893be3216542975dd632d28@tcp(wk2.h.filess.io:3307)/mydb_bowmeallog")

	mysqlStore.Init()
	userService := services.NewUserService(mysqlStore)

	userHandler := handlers.NewUserHandler(userService)
	userHandler.RegisterAuthRoutes(api)

	e.engine.Run(":8080")
}
