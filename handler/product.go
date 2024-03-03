package handler

import (
	"net/http"
	"gend.com/gind/database"
	"gend.com/gind/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetProduct(c *gin.Context) {
	cursor, err := database.Products.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch"})
		return
	}
	
	var products []model.Product
	if err = cursor.All(c, &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch"})
		return
	}

	c.JSON(http.StatusOK, products) 
}

func GetProductByName(c *gin.Context) {
	name := c.Param("name")

	var product model.Product
	err := database.Products.FindOne(c, bson.M{"name": name}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusOK, gin.H{"message": "Product not found"})
            return
        }
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func AddProduct(c *gin.Context) {
	var body model.CreateProduct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	res, err := database.Products.InsertOne(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Insert"})
		return
	}

	product := model.Product{
		ID:			res.InsertedID.(primitive.ObjectID),
		Name: 		body.Name,
		Category: 	body.Category,
		Price: 		body.Price,
		Stock:		body.Stock,
	}

	c.JSON(http.StatusCreated, product)
}
