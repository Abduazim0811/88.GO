package api

import (
	"Book/internal/http/api/handler"
	"Book/internal/repository/postgres"
	"Book/internal/service"
	"Book/internal/storage"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := storage.NewPostgres()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	repo := postgres.NewPostgresBookRepository(db)

	service := service.NewBookService(repo)

	handler := handler.NewBookHandler(service)

	router :=gin.Default()

	router.POST("/books", handler.CreateBooks)
	router.GET("/books", handler.GetAllBooks)
	router.GET("/books/:id", handler.GetbyIdBooks)
	router.PUT("/books/:id", handler.UpdateBooks)
	router.DELETE("/books/:id", handler.DeleteBooks)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(os.Getenv("PORT"))
}
