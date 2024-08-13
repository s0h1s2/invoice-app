package api

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/s0h1s2/invoice-app/internal/config"
	"github.com/s0h1s2/invoice-app/internal/db"
	"github.com/s0h1s2/invoice-app/internal/handlers"
	"github.com/s0h1s2/invoice-app/internal/services"
)

type engine struct {
	engine *gin.Engine
}

func NewEngine() *engine {
	eng := gin.Default()
	eng.Static("/uploads", "../../assets/uploads")
	return &engine{
		engine: eng,
	}
}
func (e *engine) Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load environment variables becuase of %s", err.Error())
	}
	if err = env.Parse(&config.Config); err != nil {
		log.Fatalf("Unable to load env to config becuase of %s", err.Error())
	}
	api := e.engine.Group("/api/v1")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.Db.User, config.Config.Db.Password, config.Config.Db.Host, config.Config.Db.Port, config.Config.Db.Name)

	mysqlStore := db.NewMysqlStore(dsn)

	mysqlStore.Init()

	userService := services.NewUserService(mysqlStore)
	customerService := services.NewCustomerService(mysqlStore)
	userHandler := handlers.NewUserHandler(userService)
	userHandler.RegisterAuthRoutes(api)

	customerHandler := handlers.NewCustomerHandler(customerService)
	customerHandler.RegisterCustomerRoutes(api)

	productHandler := handlers.NewProductHandler(mysqlStore)
	productHandler.RegisterProductRoutes(api)

	supplierHandler := handlers.NewSupplierHandler(mysqlStore)
	supplierHandler.RegisterSupplierRoutes(api)

	productImageUploadHandler := (handlers.NewProductImageHandler(mysqlStore))
	productImageUploadHandler.RegisterProductImageRoutes(api)

	e.engine.Run(":8080")
}
