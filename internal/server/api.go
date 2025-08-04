package server

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"wakeapi/internal/config"
	"wakeapi/internal/controllers"
	"wakeapi/internal/routes"
)

func Start() error {
	// create new echo server
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())

	_ = godotenv.Load()

	// get config file path
	filePath, exists := os.LookupEnv("CONFIG_FILE_PATH")
	if !exists {
		filePath = config.DEFAULT_CONFIG_FILE_PATH
	}

	// load/print devices
	deviceManager := config.NewDeviceManager()
	err := deviceManager.Load(filePath)
	if err != nil {
		return fmt.Errorf("error loading config from file: %v", err)
	}
	deviceManager.PrintDevices()

	// controller ctx with device manager
	controllerCtx := controllers.NewContext(deviceManager)

	// routes
	apiGrp := e.Group("/api/v1")
	routes.Register(apiGrp, controllerCtx)

	// port
	port, exists := os.LookupEnv("HTTP_PORT")
	if !exists {
		port = "8080"
	}
	// start server
	e.Logger.Fatal(e.Start(":" + port))

	return nil
}
