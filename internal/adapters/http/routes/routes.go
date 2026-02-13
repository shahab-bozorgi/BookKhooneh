package routes

import (
	"BookKhoone/infrastructure/config"
	handlers2 "BookKhoone/internal/adapters/http/handlers"
	middlewares2 "BookKhoone/internal/adapters/http/middlewares"
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
		auth.POST("/register", handlers2.RegisterHandler(db, cfg))
		auth.POST("/login", handlers2.LoginHandler(db, cfg))

	}

	users := api.Group("/users")
	{
		users.GET("/:username", handlers2.GetUserHandler(db))
		users.GET("/get_all", middlewares2.AuthMiddleware(db), handlers2.GetAllUsersHandler(db))
	}

	books := api.Group("/books")
	{
		books.POST("/create", middlewares2.AuthMiddleware(db), middlewares2.AdminMiddleware(), handlers2.CreateBookHandler(db))
		books.GET("/get_all", handlers2.GetAllBooksHandler(db))
		books.GET("/get/:id", handlers2.GetBookHandler(db))
		books.GET("/search", middlewares2.AuthMiddleware(db), handlers2.FilterBooksHandler(db))
		books.PATCH("/update/:id", middlewares2.AuthMiddleware(db), middlewares2.AdminMiddleware(), handlers2.UpdateBookHandler(db))
		books.DELETE("/delete/:id", middlewares2.AuthMiddleware(db), middlewares2.AdminMiddleware(), handlers2.DeleteBookHandler(db))
	}
	reviews := api.Group("/reviews")
	{
		reviews.POST("/create", middlewares2.AuthMiddleware(db), handlers2.CreateReviewBookHandler(db))
	}

}
