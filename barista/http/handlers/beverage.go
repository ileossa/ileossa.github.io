package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// the breverage represents data about a drink available in the menu.
type beverage struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
}

var beverages = []beverage{
	{ID: "0", Name: "Café Filtre", Description: "Café de torréfaction moyenne, doux, équilibré et riche en saveur.", Ingredients: ingredients},
	{ID: "1", Name: "Americano", Description: "Shots d'espresso couverts d'eau chaude, générant une fine couche crémeuse.", Ingredients: ingredients},
	{ID: "2", Name: "Mocha", Description: "Un espresso riche et corsé, mélangé avec du sirop saveur chocolat et du lait chauffé à la vapeur.", Ingredients: ingredients},
}

// getBeverages responds with the list of all beverages as JSON
func GetBeverages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, beverages)
}
