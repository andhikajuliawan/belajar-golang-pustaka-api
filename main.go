package main

import (
	"fmt"
	"log"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "postgres://andhikajuliawan:@localhost:5432/pustaka-api?sslmode=disable&TimeZone=Asia/Jakarta"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	fmt.Println("Database Connection Success")

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/books", handler.QueryHandler)
	v1.POST("books", handler.PostBooksHandler)

	router.Run(":8080")
}
