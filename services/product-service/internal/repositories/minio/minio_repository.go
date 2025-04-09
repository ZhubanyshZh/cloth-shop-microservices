package minio

import (
	"context"
	"fmt"
	minio2 "github.com/ZhubanyshZh/go-project-service/internal/config/minio"
	"github.com/ZhubanyshZh/go-project-service/internal/models"
	"github.com/minio/minio-go/v7"
	"os"
	"time"
)

func UploadFile(product *models.ProductCreate) ([]string, error) {
	if product.Images == nil || len(product.Images) == 0 {
		return nil, nil
	}

	var uploadedImages []string

	for _, image := range product.Images {
		imageFile, err := image.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open image file: %w", err)
		}

		objectName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), image.Filename)
		bucket := os.Getenv("MINIO_BUCKET")
		contentType := image.Header.Get("Content-Type")

		_, err = minio2.Client.PutObject(
			context.Background(),
			bucket,
			objectName,
			imageFile,
			image.Size,
			minio.PutObjectOptions{ContentType: contentType},
		)
		imageFile.Close()

		if err != nil {
			return nil, fmt.Errorf("failed to upload image to MinIO: %w", err)
		}

		imageURL := fmt.Sprintf("http://%s/%s/%s",
			os.Getenv("MINIO_ENDPOINT"), bucket, objectName)
		uploadedImages = append(uploadedImages, imageURL)
	}

	return uploadedImages, nil
}
