package routes

import (
	"BookKhoone/internal/config"
	"BookKhoone/internal/handlers"
	"BookKhoone/internal/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	api := r.Group("/api")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterHandler(db, cfg))
		auth.POST("/login", handlers.LoginHandler(db, cfg))

	}

	users := api.Group("/users")
	{
		users.GET("/:username", handlers.GetUserHandler(db))
		users.GET("/get_all", middlewares.AuthMiddleware(db), handlers.GetAllUsersHandler(db))
	}

	books := api.Group("/books")
	{
		books.POST("/create", middlewares.AuthMiddleware(db), middlewares.AdminMiddleware(), handlers.CreateBookHandler(db))
		books.GET("/get_all", handlers.GetAllBooksHandler(db))
		books.GET("/get/:title", handlers.GetBookHandler(db))
		books.PATCH("/update/:id", middlewares.AuthMiddleware(db), middlewares.AdminMiddleware(), handlers.UpdateBookHandler(db))
		books.DELETE("/delete/:id", middlewares.AuthMiddleware(db), middlewares.AdminMiddleware(), handlers.DeleteBookHandler(db))
	}

}
