package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldSecret := "CLOUDINARY_API_SECRET"
	cldName := "CLOUDINARY_CLOUD_NAME"
	cldKey := "CLOUDINARY_API_KEY"

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
