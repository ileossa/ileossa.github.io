package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// order represents data about a order
type order struct {
	ID   string `json:"id"`
	Menu int    `json:"menu"`
}

var orders = []order{}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range beverages {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "beverage not found"})
}

// postOrder adds an new order from a JSON request body
func PostOrder(c *gin.Context) {
	var newOrder order

	//Call BindJSON to bind the received JSON to new beverage's order
	if err := c.BindJSON(&newOrder); err != nil {
		return
	}
	//Add the new order to the slice and return them
	orders = append(orders, newOrder)
	c.IndentedJSON(http.StatusCreated, orders)
}
