package cmd

import (
	"backend/config"
	"backend/docs"
	"backend/internal/middleware"
	"backend/internal/module"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

//	@title			KLine
//	@version		2.0
//	@description	API for backend KLine
//	@host			localhost:8080
//	@BasePath		/api
func SetupServer(db *gorm.DB, cfg config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	module.RegisterAuthModule(api, db, cfg)

	api.Use(middleware.AuthMiddleware(cfg))
	module.RegisterAccountModule(api, db, cfg)
	module.RegisterBannerModule(api, db, cfg)
	module.RegisterDebitCardModule(api, db, cfg)
	module.RegisterUserModule(api, db, cfg)

	return r
}
