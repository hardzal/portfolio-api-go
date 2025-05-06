package utils

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/hardzal/portfolio-api-go/config"
)

func UploadToCloudinary(file *multipart.FileHeader, fileFolder string) (string, error) {
	ctx := context.Background()
	cld, err := config.SetupCloudinary()
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	uploadParams := uploader.UploadParams{
		PublicID: file.Filename,
		Folder:   fileFolder,
	}

	result, err := cld.Upload.Upload(ctx, src, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}
