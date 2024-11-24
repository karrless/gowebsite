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
		portfolioGroup.GET("/techs", portfolioController.GetListTechnologies)
		portfolioGroup.GET("/projects", portfolioController.GetListProjects)

		portfolioGroup.GET("/techs/:id", portfolioController.GetTechnology)
		portfolioGroup.GET("/projects/:id", portfolioController.GetProject)

		portfolioGroup.POST("/techs", portfolioController.CreateTechnology)
		portfolioGroup.POST("/projects", portfolioController.CreateProject)

		portfolioGroup.DELETE("/techs/:id", portfolioController.DeleteTechnology)
		portfolioGroup.DELETE("/projects/:id", portfolioController.DeleteProject)

		portfolioGroup.PATCH("/techs/:id", portfolioController.PatchTechnology)
		portfolioGroup.PATCH("/projects/:id", portfolioController.PatchProject)
	}
}
