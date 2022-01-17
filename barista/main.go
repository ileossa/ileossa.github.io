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
		api.GET("/menus", handlers.GetBeverages)
		api.GET("/order/:id", handlers.GetOrderById)
		api.POST("/order", handlers.PostOrder)
	}

	// Multipart
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", handlers.UploadFile)

	router.Run("localhost:8080")
}
