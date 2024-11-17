package controllers

import (
	"context"
	"gowebsite/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PortfolioService interface {
	ListLanguages(context.Context) ([]*models.Language, error)
	ListProjects(context.Context, *models.ProjectFilter) ([]*models.Project, error)
	CreateLanguage(context.Context, *models.Language) (*models.Language, error)
	CreateProject(context.Context, *models.Project) (*models.Project, error)
	GetLanguage(context.Context, int64) (*models.Language, error)
	GetProject(context.Context, int64) (*models.Project, error)
	UpdateLanguage(context.Context, *models.Language) (*models.Language, error)
	UpdateProject(context.Context, *models.Project) (*models.Project, error)
	PatchLanguage(context.Context, *models.Language, *models.Language) (*models.Language, error)
	PatchProject(context.Context, *models.Project, *models.Project) (*models.Project, error)
	DeleteLanguage(context.Context, int64) error
	DeleteProject(context.Context, int64) error
}

type PortfolioController struct {
	service PortfolioService
	ctx     context.Context
}

func NewPortfolioController(ctx context.Context, service PortfolioService) *PortfolioController {
	return &PortfolioController{service: service, ctx: ctx}
}

func (pc *PortfolioController) GetListLanguages(c *gin.Context) {
	languages, err := pc.service.ListLanguages(pc.ctx)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, languages)
}

func (pc *PortfolioController) GetListProjects(c *gin.Context) {
	filter := &models.ProjectFilter{}

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(400, err)
		return
	}

	projects, err := pc.service.ListProjects(pc.ctx, filter)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, projects)
}

func (pc *PortfolioController) GetLanguage(c *gin.Context) {
	languageId := c.Param("id")

	languageId64, err := strconv.ParseInt(languageId, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "LanguageID is not integer"})
		return
	}

	language, err := pc.service.GetLanguage(pc.ctx, languageId64)
	if err != nil {
		c.JSON(500, err)
		return
	}

	if language == nil {
		c.JSON(404, gin.H{"error": "Language with id " + languageId + " not found"})
		return
	}

	c.JSON(200, language)
}

func (pc *PortfolioController) GetProject(c *gin.Context) {
	projectID := c.Param("id")

	projectID64, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "ProjectID is not integer"})
		return
	}

	project, err := pc.service.GetProject(pc.ctx, projectID64)
	if err != nil {
		c.JSON(500, err)
		return
	}

	if project == nil {
		c.JSON(404, gin.H{"error": "Project with id " + projectID + " not found"})
		return
	}

	c.JSON(200, project)
}

func (pc *PortfolioController) CreateLanguage(c *gin.Context) {
	var language models.Language

	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	languageCreated, err := pc.service.CreateLanguage(pc.ctx, &language)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create language"})
		return
	}

	c.JSON(201, languageCreated)
}

func (pc *PortfolioController) CreateProject(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	projectCreated, err := pc.service.CreateProject(pc.ctx, &project)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(201, projectCreated)
}

func (pc *PortfolioController) DeleteLanguage(c *gin.Context) {
	languageID := c.Param("id")

	languageID64, err := strconv.ParseInt(languageID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "LanguageID is not integer"})
		return
	}

	err = pc.service.DeleteLanguage(pc.ctx, languageID64)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete language"})
		return
	}

	c.JSON(200, gin.H{"message": "Language deleted successfully"})
}

func (pc *PortfolioController) DeleteProject(c *gin.Context) {
	projectID := c.Param("id")

	projectID64, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "ProjectID is not integer"})
		return
	}

	err = pc.service.DeleteProject(pc.ctx, projectID64)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete project"})
		return
	}

	c.JSON(200, gin.H{"message": "Project deleted successfully"})
}

func (pc *PortfolioController) PutLanguage(c *gin.Context) {
	var language models.Language
	languageID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid LanguageID"})
		return
	}

	language.ID = languageID

	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	languageUpdated, err := pc.service.UpdateLanguage(pc.ctx, &language)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update language"})
		return
	}

	c.JSON(200, languageUpdated)
}

func (pc *PortfolioController) PutProject(c *gin.Context) {
	var project models.Project
	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ProjectID"})
		return
	}

	project.ID = projectID

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	projectUpdated, err := pc.service.UpdateProject(pc.ctx, &project)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(200, projectUpdated)
}

func (pc *PortfolioController) PatchLanguage(c *gin.Context) {

	languageID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid LanguageID"})
		return
	}

	language, err := pc.service.GetLanguage(pc.ctx, languageID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get language"})
		return
	}

	if language == nil {
		c.JSON(404, gin.H{"error": "Language with id " + c.Param("id") + " not found"})
		return
	}
	var languageUpdate models.Language

	if err := c.ShouldBindJSON(&languageUpdate); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	languageUpdated, err := pc.service.PatchLanguage(pc.ctx, language, &languageUpdate)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update language"})
		return
	}
	c.JSON(200, languageUpdated)
}

func (pc *PortfolioController) PatchProject(c *gin.Context) {

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid LanguageID"})
		return
	}

	project, err := pc.service.GetProject(pc.ctx, projectID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get language"})
		return
	}

	if project == nil {
		c.JSON(404, gin.H{"error": "Language with id " + c.Param("id") + " not found"})
		return
	}
	var projectUpdate models.Project

	if err := c.ShouldBindJSON(&projectUpdate); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	projectUpdated, err := pc.service.PatchProject(pc.ctx, project, &projectUpdate)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update project"})
		return
	}
	c.JSON(200, projectUpdated)
}
