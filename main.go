package main

import (
	"log"

	"github.com/upload_images/controllers"
	"github.com/upload_images/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var err error

func main() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("./public/html/*")
	router.Static("/public", "./public")

	upload := new(controllers.UploadController)

	router.GET("/", upload.Index)
	router.POST("/upload", upload.Upload)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.tmpl", gin.H{})
	})
	router.Run(":9000")
}
