package routes

import (
	"BookKhoone/internal/config"
	"BookKhoone/internal/handlers"
	"BookKhoone/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterHandler(db, cfg))
		auth.POST("/login", handlers.LoginHandler(db, cfg))

	}
	books := api.Group("/books")
	{
		books.POST("/create", middlewares.AuthMiddleware(), handlers.CreateBookHandler(db))
		books.GET("/get_all", handlers.GetAllBooksHandler(db))
		books.GET("/get/:name", handlers.GetBookHandler(db))
	}

}
