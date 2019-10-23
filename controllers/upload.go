package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/upload_images/services"
)

type UploadController struct{}

func (ctrl UploadController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Upload Form",
		"auth":  os.Getenv("auth"),
	})
}

func (ctrl UploadController) Upload(c *gin.Context) {
	service := new(services.UploadService)
	err := service.Execute(c)
	if err != nil {
		c.HTML(http.StatusForbidden, "index.tmpl", gin.H{
			"title":   "Upload Form",
			"auth":    os.Getenv("auth"),
			"message": err.Error(),
		})
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "Upload Form",
			"auth":    os.Getenv("auth"),
			"message": "Upload file success!",
		})
	}
}
