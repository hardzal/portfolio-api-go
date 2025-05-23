package utils

import (
	"context"
	"errors"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/hardzal/portfolio-api-go/config"
)

var allowedFile = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

func UploadToCloudinary(file *multipart.FileHeader, fileFolder string) (string, error) {
	ctx := context.Background()

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedFile[ext] {
		return "", errors.New("only jpg/png images are allowed")
	}

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
		PublicID:     strings.TrimSuffix(file.Filename, ext),
		Folder:       "portfolio/" + fileFolder,
		ResourceType: "image",
	}

	result, err := cld.Upload.Upload(ctx, src, uploadParams)
	if err != nil {
		log.Printf("Cloudinary error: %v", err)
		return "", err
	} else {
		log.Printf("Cloudinary response: %+v", result) // <-- Ini yang penting
	}

	// Cek URL
	if result != nil && result.SecureURL == "" {
		log.Println("WARNING: Upload success but no URL generated")
	}

	imageUrl := result.SecureURL

	return imageUrl, nil
}
