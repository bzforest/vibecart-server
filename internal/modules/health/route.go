package health

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/vibecart/vibecart-server/internal/config"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, cld *cloudinary.Cloudinary, cfg *config.Config) {
	r.GET("/health", HealthHandler(cfg))
	r.GET("/test/aiven", TestAivenHandler(db))
	r.GET("/test/cloudinary", TestCloudinaryHandler(cld, cfg))

	api := r.Group("/api/v1")
	{
		api.GET("/health", APIHealthHandler)
	}
}