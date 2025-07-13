package main

import (
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
