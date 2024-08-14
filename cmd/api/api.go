package api

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/s0h1s2/invoice-app/internal/config"
	mysqlstore "github.com/s0h1s2/invoice-app/internal/db/mysqlStore"
	"github.com/s0h1s2/invoice-app/internal/handlers"
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

	db := mysqlstore.NewMysqlStore(dsn)

	db.Init()
	userStore := mysqlstore.NewMysqlUserStore(db)
	productStore := mysqlstore.NewMysqlProductStore(db)
	customerStore := mysqlstore.NewMysqlCustomerStore(db)
	supplierStore := mysqlstore.NewMysqlSupplierStore(db)
	productImageStore := mysqlstore.NewProductImageStore(db)

	userHandler := handlers.NewUserHandler(userStore)
	userHandler.RegisterAuthRoutes(api)

	customerHandler := handlers.NewCustomerHandler(customerStore)
	customerHandler.RegisterCustomerRoutes(api)
	productHandler := handlers.NewProductHandler(productStore, supplierStore)
	productHandler.RegisterProductRoutes(api)

	supplierHandler := handlers.NewSupplierHandler(supplierStore)
	supplierHandler.RegisterSupplierRoutes(api)

	productImageUploadHandler := handlers.NewProductImageHandler(productImageStore, productStore)
	productImageUploadHandler.RegisterProductImageRoutes(api)

	e.engine.Run(":8080")
}
