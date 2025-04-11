package main

import (
	"go-bookstore/internal/controller"
	"go-bookstore/internal/repository"
	"go-bookstore/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize router
	router := gin.Default()

	// Initialize repository, service and controller

	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	// Routes definition
	v1 := router.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.GET("", bookController.GetAllBooks)
			books.GET("/:id", bookController.GetBookByID)
			books.POST("", bookController.CreateBook)
			books.PUT("/:id", bookController.UpdateBook)
			books.DELETE("/:id", bookController.DeleteBook)
		}
	}

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Xin chào từ server!")
	})

	// Start the server
	log.Println("Starting server on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
