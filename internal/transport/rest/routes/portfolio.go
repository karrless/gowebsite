package routes

import (
	"context"
	"gowebsite/internal/repository"
	"gowebsite/internal/service"
	"gowebsite/internal/transport/rest/controllers"
	"gowebsite/pkg/db/postgres"

	"github.com/gin-gonic/gin"
)

func PortfolioRoutes(ctx context.Context, r *gin.RouterGroup, db *postgres.DB) {
	portfolioRepo := repository.NewPortfolioRepository(db)
	portfolioService := service.NewPortfolioService(portfolioRepo)
	portfolioController := controllers.NewPortfolioController(ctx, portfolioService)
	portfolioGroup := r.Group("/portfolio")
	{
		portfolioGroup.GET("/languages", portfolioController.GetListLanguages)
		portfolioGroup.GET("/projects", portfolioController.GetListProjects)

		portfolioGroup.GET("/languages/:id", portfolioController.GetLanguage)
		portfolioGroup.GET("/projects/:id", portfolioController.GetProject)

		portfolioGroup.POST("/languages", portfolioController.CreateLanguage)
		portfolioGroup.POST("/projects", portfolioController.CreateProject)

		portfolioGroup.DELETE("/languages/:id", portfolioController.DeleteLanguage)
		portfolioGroup.DELETE("/projects/:id", portfolioController.DeleteProject)

		portfolioGroup.PUT("/languages/:id", portfolioController.PutLanguage)
		portfolioGroup.PUT("/projects/:id", portfolioController.PutProject)

		portfolioGroup.PATCH("/languages/:id", portfolioController.PatchLanguage)
		portfolioGroup.PATCH("/projects/:id", portfolioController.PatchProject)
	}
}
