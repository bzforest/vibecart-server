package cloudinary

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/vibecart/vibecart-server/internal/config"
)

func ConnectCloudinary(cfg *config.Config) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(
		cfg.CloudinaryCloudName,
		cfg.CloudinaryAPIKey,
		cfg.CloudinaryAPISecret,
	)

	if err != nil {
		log.Fatal("Failed to initialize Cloudinary: ", err)
	}

	log.Println("Cloudinary initialized successfully")
	return cld
}