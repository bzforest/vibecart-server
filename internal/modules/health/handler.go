package health

import (
	"context"
	"net/http"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/vibecart/vibecart-server/internal/config"
	"gorm.io/gorm"
)

// @Summary Server Health Check
// @Description Check if backend server is running
// @Tags Health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Vibecart server is running",
		})
	}
}

// @Summary Test Aiven MySQL Connection
// @Description Execute SELECT 1 to verify database connection
// @Tags Test
// @Produce json
// @Success 200 {object} DBTestResponse
// @Failure 500 {object} DBTestErrorResponse
// @Router /test/aiven [get]
func TestAivenHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result int

		err := db.Raw("SELECT 1").Scan(&result).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Aiven connection failed",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Aiven connected successfully",
		})
	}
}

// @Summary Test Cloudinary Connection
// @Description Upload small text file to verify Cloudinary credential
// @Tags Test
// @Produce json
// @Success 200 {object} CloudinaryTestResponse
// @Failure 500 {object} CloudinaryTestErrorResponse
// @Router /test/cloudinary [get]
func TestCloudinaryHandler(cld *cloudinary.Cloudinary, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		_ , err := cld.Upload.Upload(
			context.Background(),
			strings.NewReader("vibecart cloudinary connection test"),
			uploader.UploadParams{
				Folder:   cfg.CloudinaryFolder + "/test",
				PublicID: "connection-test",
			},
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Cloudinary connection failed",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Cloudinary connected successfully",
		})
	}
}

// @Summary API v1 Health Check
// @Description Check if API v1 is running
// @Tags API
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /api/v1/health [get]
func APIHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Vibecart API v1 is running",
	})
}