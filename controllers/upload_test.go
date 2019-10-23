package controllers_test

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/upload_images/controllers"
	"github.com/upload_images/db"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	router.LoadHTMLGlob("../public/html/*")
	upload := new(controllers.UploadController)

	router.GET("/", upload.Index)
	router.POST("/upload", upload.Upload)

	return router
}

func TestIndex(t *testing.T) {
	testRouter := SetupRouter()
	req, _ := http.NewRequest("GET", "/", nil)

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

var err error

func TestUpload(t *testing.T) {
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}
	testRouter := SetupRouter()
	path := "../tmp/Screen Shot 2019-09-19 at 11.16.02 PM.png"
	params := map[string]string{"auth": "n6n9Bnee7N"}
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}
