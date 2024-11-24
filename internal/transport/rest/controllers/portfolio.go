package controllers

import (
	"context"
	"fmt"
	"gowebsite/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PortfolioService interface {
	CreateTechnology(ctx context.Context, technology *models.Technology) (int64, error)
	GetTechnology(ctx context.Context, id int64) (*models.Technology, error)
	ListTechnologies(ctx context.Context, filter *models.TechnologyFilter) ([]*models.Technology, error)
	DeleteTechnology(ctx context.Context, id int64) error
	PatchTechnology(ctx context.Context, technology *models.Technology) error
	CreateProject(ctx context.Context, project *models.Project) (int64, error)
	GetProject(ctx context.Context, id int64) (*models.Project, error)
	ListProjects(ctx context.Context, filter *models.ProjectFilter) ([]*models.Project, error)
	DeleteProject(ctx context.Context, id int64) error
	PatchProject(ctx context.Context, project *models.Project, projectUpdate *models.Project) error
}

type PortfolioController struct {
	service PortfolioService
	ctx     context.Context
}

func NewPortfolioController(ctx context.Context, service PortfolioService) *PortfolioController {
	return &PortfolioController{service: service, ctx: ctx}
}

// @Summary Technology list
// @Description Get technology list
// @Tags Portfolio
// @Accept json
// @Param tech_id query []int64 false "Technology ID"
// @Param sort_field query string false "Sort field"
// @Param sort_order query string false "Sort order"
// @Param limit query int false "Limit of projects"
// @Param Offset query int false "Offset of projects"
// @Produce json
// @Success 200 {array} models.Technology "Technology"
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal error"
// @Router /portfolio/techs [get]
func (pc *PortfolioController) GetListTechnologies(c *gin.Context) {
	filter := &models.TechnologyFilter{}

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(400, err)
		return
	}

	languages, err := pc.service.ListTechnologies(pc.ctx, filter)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, languages)
}

// @Summary Project list
// @Description Get project list
// @Tags Portfolio
// @Accept json
// @Param tech_id query []int64 false "Language ID"
// @Param is_active query bool false "Is active"
// @Param is_archived query bool false "Is archived"
// @Param is_developing query bool false "Is developing"
// @Param sort_field query string false "Sort field"
// @Param sort_order query string false "Sort order"
// @Param limit query int false "Limit of projects"
// @Param Offset query int false "Offset of projects"
// @Produce json
// @Success 200 {array} models.Project "Project"
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal error"
// @Router /portfolio/projects [get]
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

// @Summary Technology
// @Description Get technology
// @Tags Portfolio
// @Accept json
// @Param id path int true "Technology ID"
// @Produce json
// @Success 200 {object} models.Technology "Technology"
// @Failure 404 {object} error "Technology not found"
// @Failure 500 {object} error "Internal error"
// @Failure 400 {object} error "Bad request"
// @Router /portfolio/techs/{id} [get]
func (pc *PortfolioController) GetTechnology(c *gin.Context) {
	technologyID := c.Param("id")

	technologyID64, err := strconv.ParseInt(technologyID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Technology ID is not integer"})
		return
	}

	technology, err := pc.service.GetTechnology(pc.ctx, technologyID64)
	if err != nil {
		c.JSON(500, err)
		return
	}

	if technology == nil {
		c.JSON(404, gin.H{"error": "Technology with id " + technologyID + " not found"})
		return
	}

	c.JSON(200, technology)
}

// @Summary Project
// @Description Get project
// @Tags Portfolio
// @Accept json
// @Param id path int true "Project ID"
// @Produce json
// @Success 200 {object} models.Project "Project"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Project nor found"
// @Failure 500 {object} error "Internal error"
// @Router /portfolio/projects/{id} [get]
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

// @Summary Create Technology
// @Description Create technology and write to database
// @Tags Portfolio
// @Accept json
// @Param name body string true "Technology name"
// @Param svg body string flase "Technology svg"
// @Produce json
// @Success 200 {object} int64 "Technology ID"
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal error"
// @Router /portfolio/techs [post]
func (pc *PortfolioController) CreateTechnology(c *gin.Context) {
	var technology models.Technology

	if err := c.ShouldBindJSON(&technology); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	technologyID, err := pc.service.CreateTechnology(pc.ctx, &technology)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create technology"})
		return
	}

	c.JSON(201, technologyID)
}

// @Summary Create Project
// @Description Create project and write to database
// @Tags Portfolio
// @Accept json
// @Param title body string true "Project title"
// @Param version body string true "Project version"
// @Param description body string true "Project description"
// @Param tech_id body []int64 true "Technology ID"
// @Param isActive body bool true "Is active"
// @Param isArchived body bool true "Is archived"
// @Param isDeveloping body bool true "Is developing"
// @Param links body []string false "Links"
// @Produce json
// @Success 200 {object} int64 "Project ID"
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal error"
// @Router /portfolio/projects [post]
func (pc *PortfolioController) CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	fmt.Printf("%+v\n", project)

	projectID, err := pc.service.CreateProject(pc.ctx, &project)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(201, projectID)
}

// @Summary Delete Technology
// @Description Delete technology
// @Tags Portfolio
// @Accept json
// @Param id path int true "Technology ID"
// @Produce json
// @Success 200 {object} map[string]any "Message"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Technology not found"
// @Failure 500 {object} error "Internal error"
// @Router /portfolio/techs/{id} [delete]
func (pc *PortfolioController) DeleteTechnology(c *gin.Context) {
	technologyID := c.Param("id")

	technologyID64, err := strconv.ParseInt(technologyID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "LanguageID is not integer"})
		return
	}

	err = pc.service.DeleteTechnology(pc.ctx, technologyID64)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete language"})
		return
	}

	c.JSON(200, gin.H{"message": "Language deleted successfully"})
}

// @Summary Delete Project
// @Description Delete project
// @Tags Portfolio
// @Accept json
// @Param id path int true "Project ID"
// @Produce json
// @Success 200 {object} map[string]any "Message"
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal error"
// @Router /portfolio/projects/{id} [delete]
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

// @Summary Update Technology
// @Description Update technology
// @Tags Portfolio
// @Accept json
// @Param id path int true "Technology ID"
// @Param name body string false "Technology name"
// @Param svg body string flase "Technology svg"
// @Produce json
// @Success 200 {object} map[string]any "Message"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Technology not found"
// @Failure 500 {object} error "Internal error"
// @Router /portfolio/techs/{id} [patch]
func (pc *PortfolioController) PatchTechnology(c *gin.Context) {
	technologyID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid technologyID"})
		return
	}

	var technologyUpdate models.Technology

	if err := c.ShouldBindJSON(&technologyUpdate); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	technologyUpdate.ID = technologyID
	err = pc.service.PatchTechnology(pc.ctx, &technologyUpdate)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update technology"})
		return
	}
	c.JSON(200, gin.H{"message": "Technology updated successfully"})
}

// @Summary Update Project
// @Description Update project
// @Tags Portfolio
// @Accept json
// @Param id path int true "Project ID"
// @Param title body string false "Project title"
// @Param version body string false "Project version"
// @Param description body string false "Project description"
// @Param language_id body int false "Language ID"
// @Param isActive body bool false "Is active"
// @Param isArchived body bool false "Is archived"
// @Param isDeveloping body bool false "Is developing"
// @Param links body string false "Links"
// @Produce json
// @Success 200 {object} map[string]any "Message"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Technology not found"
// @Failure 500 {object} error "Internal error"
// @Router /portfolio/projects/{id} [patch]
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

	err = pc.service.PatchProject(pc.ctx, project, &projectUpdate)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update project"})
		return
	}
	c.JSON(200, gin.H{"message": "Project updated successfully"})
}
