// @title           BookKhoone API
// @version         1.0
// @description     BookKhoone backend with Gin
// @host            localhost:8080
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	_ "BookKhoone/docs"
	"BookKhoone/internal/config"
	"BookKhoone/internal/database"
	"BookKhoone/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	r := gin.Default()

	routes.SetupRoutes(r, db, cfg)
	r.Run(":" + cfg.Port)
}
