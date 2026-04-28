package main

import (
	"log"

	"github.com/vibecart/vibecart-server/internal/config"
	dbpkg "github.com/vibecart/vibecart-server/internal/database"
	routerpkg "github.com/vibecart/vibecart-server/internal/router"
	pkgcloudinary "github.com/vibecart/vibecart-server/pkg/cloudinary"
)

// @title Vibecart API
// @version 1.0
// @description Backend API documentation for Vibecart project.
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.LoadConfig()

	db := dbpkg.ConnectMySQL(cfg)

	cld := pkgcloudinary.ConnectCloudinary(cfg)

	r := routerpkg.SetupRouter(db, cld, cfg)

	log.Println("Server running on port", cfg.AppPort)

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}