package main

import (
	"os"

	"github.com/asofyana/school-adm/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	router.Static("/assets", "./assets")

	// Initialize the routes
	routes.InitializeRoutes(router)

	// Start serving the application
	port := os.Getenv("PORT")
	router.Run(":" + port)

}
