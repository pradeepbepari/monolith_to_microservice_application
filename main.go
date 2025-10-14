package main

import (
	"fmt"
	"log"
	"pov_golang/database"
	"pov_golang/database/migrations"
	"pov_golang/handlers"
	"pov_golang/logger"
	"pov_golang/repository"
	"pov_golang/routes"
	"pov_golang/service"

	_ "pov_golang/docs"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	AppPort            string `envconfig:"APP_PORT" `
	DBName             string `envconfig:"DB_NAME" `
	DBUser             string `envconfig:"DB_USER"`
	DBPassword         string `envconfig:"DB_PASSWORD"`
	DBHost             string `envconfig:"DB_HOST" `
	DBPort             string `envconfig:"DB_PORT" `
	NewRelicLisenceKey string `envconfig:"NEW_RELIC_LICENSE_KEY"`
	NewRelicAppName    string `envconfig:"NEW_RELIC_APP_NAME"`
	MigrationPath      string `envconfig:"MIGRATIONS_PATH" `
}

// @title           pov_Golang API
// @version         1.0
// @description     This is a sample server for a WORK store.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	// app, err := newrelic.NewApplication(
	// 	newrelic.ConfigAppName(config.NewRelicAppName),
	// 	newrelic.ConfigLicense(config.NewRelicLisenceKey),
	// 	newrelic.ConfigDistributedTracerEnabled(true),
	// 	newrelic.ConfigAppLogForwardingEnabled(true),
	// )
	if err != nil {
		log.Fatal("Could not start New Relic", err)
	}
	// _logger := logger.NewLogger(app)
	_logger := logger.NewLogger(nil)

	db, err := database.ConnectDB(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	if err := migrations.RunMigration(db, config.MigrationPath); err != nil {
		log.Fatal("Could not run the migrations", err)
	}
	repo := repository.NewRepository(db, _logger)
	service := service.NewService(repo)
	handler := handlers.NewHandler(service)
	di := routes.Dependencies{
		UserHandler: handler,
	}
	router := gin.Default()
	routes.ApiRoutes(di, router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	defer database.CloseDB(db)
	router.Run(fmt.Sprintf(":%s", config.AppPort))

}
