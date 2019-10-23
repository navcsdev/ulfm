package models

import (
	"context"
	"time"

	"github.com/upload_images/db"
	"github.com/upload_images/forms"
)

const CollectionNameImage = "image"

type Image struct {
	FileName    string `bson:file_name`
	ContentType string `bson:content_type`
	Size        int64
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type ImageModel struct{}

func (m ImageModel) Create(form forms.UploadForm) error {
	collection := db.GetDB().Collection(CollectionNameImage)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	img := Image{
		FileName:    form.Image.Filename,
		ContentType: form.Image.Header["Content-Type"][0],
		Size:        form.Image.Size,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := collection.InsertOne(ctx, img)
	return err
}
