package app

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"payment-gateway/internal/controller"
	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"
	"payment-gateway/internal/routes"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type App struct {
	Router *gin.Engine
}

var databaseInstance *gorm.DB

func NewApp() *App {
	router := gin.Default()

	db := InitDb()
	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	ctrl := controller.NewContoller(srv)

	setupRoutes(router, ctrl)

	return &App{
		Router: router,
	}
}

func InitDb() *gorm.DB {
	var err error
	databaseInstance, err = connectDb()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	err = performMigration()
	if err != nil {
		log.Fatalf("Could not auto migrate: %v", err)
	}
	return databaseInstance
}

func connectDb() (*gorm.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDatabase, err := databaseConnection.DB()
	if err != nil {
		return nil, err
	}

	idleConns, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	openConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	connLifetime, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))

	sqlDatabase.SetMaxIdleConns(idleConns)
	sqlDatabase.SetMaxOpenConns(openConns)
	sqlDatabase.SetConnMaxLifetime(time.Duration(connLifetime) * time.Second)

	return databaseConnection, nil
}

func performMigration() error {
	err := databaseInstance.AutoMigrate(
		&model.User{},
		&model.Store{},
		&model.Product{},
		&model.Payment{},
		&model.Order{},
		&model.OrderItem{},
	)
	if err != nil {
		return err
	}
	return nil
}

func setupRoutes(router *gin.Engine, ctrl controller.Controller) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "Hello world"})
	})

	routes.UserRoutes(router, ctrl)
	routes.ProductRoutes(router, ctrl)
	routes.OrderRoutes(router, ctrl)
	routes.OrderItemRoutes(router, ctrl)
	routes.PaymentRoutes(router, ctrl)
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}
