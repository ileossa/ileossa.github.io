package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Ping
// @Description We can check if application is up
// @Produce json
// @Success 200 {object} string "return UP if application is ready"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "UP")
}
