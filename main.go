package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"client-api/pkg/api"
)

func main() {

	router := gin.Default()
	// router.GET("/cars", getCars)
	// router.GET("/cars/:id", getCarByID)
	// router.POST("/cars", createCar)
	// router.DELETE("/cars/:id", deleteCar)

	router.POST("/orders", createOrderHandler)

	// router.POST("/products", createProductHandler)
	router.GET("/products", getProductsHandler)

	router.Run("localhost:8080")
}

func createOrderHandler(c *gin.Context) {
	var order api.Orders
	if err := c.BindJSON(&order); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create an order"})
		return
	}

	orderResponse := createOrder(&order)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": orderResponse})
}

// func createProductHandler(c *gin.Context) {
// 	var product api.Products
// 	if err := c.BindJSON(&product); err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create a product"})
// 		return
// 	}

// 	productResponse := createProduct(&product)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"message": productResponse})
// }

func getProductsHandler(c *gin.Context) {
	productResponse := getProducts()

	c.IndentedJSON(http.StatusOK, gin.H{"message": productResponse})
}
