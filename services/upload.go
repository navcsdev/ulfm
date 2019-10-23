package services

import (
	"errors"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/upload_images/forms"
	"github.com/upload_images/models"
)

type UploadService struct{}

var imageModel = new(models.ImageModel)

func (service UploadService) Validate(uploadForm forms.UploadForm) (err error) {
	if uploadForm.Auth != os.Getenv("auth") {
		return errors.New("Incorrect authentication")
	}
	maxSizeImage, err := strconv.ParseInt(os.Getenv("max_size_file"), 10, 64)
	if err != nil {
		return err
	}
	if uploadForm.Image.Size > maxSizeImage {
		return errors.New("Can not upload image larger than 8 megabytes")
	}

	return nil
}

func (service UploadService) Execute(c *gin.Context) (err error) {
	var uploadForm forms.UploadForm
	c.Bind(&uploadForm)
	err = service.Validate(uploadForm)
	if err != nil {
		return err
	}
	c.SaveUploadedFile(uploadForm.Image, "./tmp/"+uploadForm.Image.Filename)
	return imageModel.Create(uploadForm)
}
