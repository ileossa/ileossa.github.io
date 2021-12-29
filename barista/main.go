package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ileossa/barista/http/handlers"
)

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
	api := router.Group("/api")
	{
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
