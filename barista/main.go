package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ileossa/barista/http/handlers"
)

// @title Swagger barista world application
// @version 0.1-beta
// @description This is a poc, to learn go project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.github.com/ileossa
// @contact.email none@email.com

// @license.name GNU AGPLv3
// @license.url https://spdx.org/licenses/AGPL-3.0-or-later.html
func main() {
	router := gin.Default()

	// Templates HTML routes
	router.LoadHTMLGlob("./public/templates/*.html")
	templates := router.Group("/")
	{
		templates.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
		templates.GET("/order", func(c *gin.Context) {
			c.HTML(http.StatusOK, "order.html", nil)
		})
	}

	// API routes
	api := router.Group("/api/v1")
	{
		api.GET("/ping", handlers.Ping)
		beverages := api.Group("/beverages")
		{
			beverages.GET("/", handlers.GetBeverages)
			beverages.GET("/:id", handlers.GetOrderById)
			beverages.POST("/", handlers.NewBeverage)
		}
		order := api.Group("/order")
		{
			order.GET("/:id", handlers.GetOrderById)
			order.POST("/", handlers.PostOrder)
		}

		user := api.Group("/user")
		{
			user.POST("/", handlers.NewUser)
			// user.GET("/:id", handlers.GetUserById)
			// user.GET("/", handlers.GetUsers)
		}
	}

	router.Run("localhost:8080")
}
