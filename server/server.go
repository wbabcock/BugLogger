package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/wbabcock/BugLogger/configuration"
	"github.com/wbabcock/BugLogger/routes"
)

func main() {
	// Load config file
	configLocation := "config.json"

	// No config file was passed used the default in the same directory
	if len(os.Args) > 1 {
		args := os.Args[1]
		configLocation = fmt.Sprintf("%s", args)
	}

	// Load the configuration file
	config, _ := configuration.LoadConfig(configLocation)

	fmt.Println("Bug Logger API")

	// Set flag for Production
	gin.SetMode(gin.ReleaseMode)

	// API router
	router := gin.Default()
	
	routerConfig := cors.DefaultConfig()
	routerConfig.AllowAllOrigins = config.API.AllowAllOrigins
	routerConfig.AllowOrigins = config.API.AllowOrigins
	routerConfig.AllowMethods = config.API.AllowMethods
	routerConfig.AllowHeaders = config.API.AllowHeaders
	router.Use(cors.New(routerConfig))

	// Import all the routes that will be used in the API
	routes.InitBugLogger(router)

	log.WithFields(log.Fields{"port": config.API.Port}).Info("Server running on localhost...")
	router.Run(":" + fmt.Sprintf("%d", config.API.Port))
}
