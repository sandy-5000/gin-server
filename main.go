package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"gend.com/gind/initalizer"
	"gend.com/gind/handler"
	"gend.com/gind/database"
)

func main() {
	initalizer.Load()
	mongoURL := os.Getenv("MONGO_URL")
	dbName := "gind"

	err := database.Init(mongoURL, dbName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DB Connected")

	defer func() {
		err := database.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from server",
		})
	})

	r.GET("/products", handler.GetProduct)
	r.GET("/products/:name", handler.GetProductByName)
	r.POST("/products", handler.AddProduct)
	
	r.Run()
}
