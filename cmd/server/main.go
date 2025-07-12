package main

import (
	"BookKhoone/internal/config"
	"BookKhoone/internal/database"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)

	fmt.Println("Connected to database:", cfg.DBName)
	fmt.Println("Server running on port:", cfg.Port)

	// Gin server will be added later
	_ = db // just to avoid unused warning for now
}
