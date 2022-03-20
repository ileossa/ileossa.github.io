package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// the breverage represents data about a drink available in the menu.
type beverage struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
}

var beverages = []beverage{
	{ID: 0, Name: "Café Filtre", Description: "Café de torréfaction moyenne, doux, équilibré et riche en saveur.", Ingredients: ingredients},
	{ID: 1, Name: "Americano", Description: "Shots d'espresso couverts d'eau chaude, générant une fine couche crémeuse.", Ingredients: ingredients},
	{ID: 2, Name: "Mocha", Description: "Un espresso riche et corsé, mélangé avec du sirop saveur chocolat et du lait chauffé à la vapeur.", Ingredients: ingredients},
}

// @BasePath /api/v1
// GetBeverages godoc
// @Summary GetBeverages responds with the list of all beverages as JSON
// @Schemes
// @Description list of all beverages
// @Tags Beverages
// @Accept json
// @Produce json
// @Success 200 {struct} beverage
// @Router /api/v1 [get]
func GetBeverages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, beverages)
}

// @BasePath /api/v1/beverages
// GetBeveragesById godoc
// @Summary GetBeverages responds with the beverage by Id passing as JSON
// @Schemes
// @Description list of beverage selected
// @Tags Beverages
// @Accept json
// @Produce json
// @Success 200 {struct} beverage
// @Router /api/v1 [get]
func GetBeverageById(c *gin.Context) {

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "beverage not found"})
}

// @BasePath /api/v1/beverages
// NewBeverage godoc
// @Summary GetBeverages responds with the beverage by Id passing as JSON
// @Schemes
// @Description Add new beverage on the list of selection
// @Tags Beverages
// @Accept json
// @Produce json
// @Success 200 {struct} beverage
// @Router /api/v1 [post]
func NewBeverage(c *gin.Context) {
	var beverage beverage

	if err := c.BindJSON(&beverage); err != nil {
		return
	}
	//input id
	beverage.ID = int(beverages[len(beverages)-1].ID) + 1
	// Add the new order to the slice and return them
	beverages = append(beverages, beverage)
	c.IndentedJSON(http.StatusCreated, beverages)
}
