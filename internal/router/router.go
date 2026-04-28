package router

import (
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vibecart/vibecart-server/docs"
	"github.com/vibecart/vibecart-server/internal/config"
	healthtest "github.com/vibecart/vibecart-server/internal/modules/health"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cld *cloudinary.Cloudinary, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://vibecart.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	healthtest.RegisterRoutes(r, db, cld, cfg)

	return r
}