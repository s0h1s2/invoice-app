package api

import (
	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/handlers"
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

	// userService := services.NewUserService()
	userHandler := handlers.NewUserHandler(nil)
	userHandler.RegisterAuthRoutes(api)

	e.engine.Run(":8080")
}
