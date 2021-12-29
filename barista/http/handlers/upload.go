package handlers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {

	//source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	rd := uuid.NewString()
	rd = strings.ReplaceAll(rd, "-", "")
	last := file.Filename[strings.LastIndex(file.Filename, ".")+1:]

	filename := filepath.Base(rd + "." + last)
	dist := "./uploaded/" + filename
	if err := c.SaveUploadedFile(file, dist); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "File"+file.Filename+" uploaded on "+dist)
}
