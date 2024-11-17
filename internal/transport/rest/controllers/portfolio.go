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

// @Summary Список ЯП
// @Description Получить список всех языков программирования из базы данных
// @Tags Portfolio
// @Accept json
// @
// @Produce json
// @Success 200 {array} models.Language "Информация о языке"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
// @Router /portfolio/languages [get]
func (pc *PortfolioController) GetListLanguages(c *gin.Context) {
	languages, err := pc.service.ListLanguages(pc.ctx)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, languages)
}

// @Summary Список проектов
// @Description Получить список всех проектов из базы данных
// @Tags Portfolio
// @Accept json
// @Param lang_id query int false "Language ID"
// @Param is_active query bool false "Is active"
// @Param is_archived query bool false "Is archived"
// @Param is_developing query bool false "Is developing"
// @Param sort_field query string false "Sort field"
// @Param sort_order query string false "Sort order"
// @Param limit query int false "Limit of projects"
// @Param Offset query int false "Offset of projects"
// @Produce json
// @Success 200 {array} models.Project "Информация о проекте"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
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

// @Summary Язык программирования
// @Description Получить язык программирования из базы данных
// @Tags Portfolio
// @Accept json
// @Param id path int true "Language ID"
// @Produce json
// @Success 200 {object} models.Language "Информация о языке"
// @Failure 404 {object} error "Язык программирования не найден"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
// @Failure 400 {object} error "Невалидные данные"
// @Router /portfolio/languages/{id} [get]
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

// @Summary Проект
// @Description Получить проект из базы данных
// @Tags Portfolio
// @Accept json
// @Param id path int true "Project ID"
// @Produce json
// @Success 200 {object} models.Project "Информация о проекте"
// @Failure 400 {object} error "Невалидные данные"
// @Failure 404 {object} error "Проект не найден"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
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

// @Summary Создать ЯП
// @Description Создать язык программирования и записать в базу данных
// @Tags Portfolio
// @Accept json
// @Param name body string true "Language name"
// @Param svg body string flase "Language svg"
// @Produce json
// @Success 200 {object} models.Language "Информация о языке"
// @Failure 400 {object} error "Невалидные данные"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
// @Router /portfolio/languages [post]
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

// @Summary Создать Проект
// @Description Создать проект и записать в базу данных
// @Tags Portfolio
// @Accept json
// @Param title body string true "Project title"
// @Param version body string true "Project version"
// @Param description body string true "Project description"
// @Param language_id body int true "Language ID"
// @Param isActive body bool true "Is active"
// @Param isArchived body bool true "Is archived"
// @Param isDeveloping body bool true "Is developing"
// @Param GHLink body string false "GitHub link"
// @Param TGLink body string false "Telegram link"
// @Param HTTPLink body string false "HTTP link"
// @Produce json
// @Success 200 {object} models.Project "Информация о проекте"
// @Failure 400 {object} error "Невалидные данные"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
// @Router /portfolio/projects [post]
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

// @Summary Удалить ЯП
// @Description Удалить язык программирования из базы данных
// @Tags Portfolio
// @Accept json
// @Param id path int true "Language ID"
// @Produce json
// @Success 200 {object} map[string]any "Сообщение о успешном удалении языка"
// @Failure 400 {object} error "Невалидные данные"
// @Failure 404 {object} error "ЯП не найден"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
// @Router /portfolio/languages/{id} [delete]
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

// @Summary Удалить Проект
// @Description Удалить проект из базы данных
// @Tags Portfolio
// @Accept json
// @Param id path int true "Project ID"
// @Produce json
// @Success 200 {object} map[string]any "Сообщение о успешном удалении проекта"
// @Failure 400 {object} error "Невалидные данные"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
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

// @Summary Заменить ЯП
// @Description Заменить язык программирования и записать в базу данных
// @Tags Portfolio
// @Accept json
// @Param id path int true "Language ID"
// @Param name body string true "Language name"
// @Param svg body string flase "Language svg"
// @Produce json
// @Success 200 {object} models.Language "Информация о языке"
// @Failure 400 {object} error "Невалидные данные"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
// @Router /portfolio/languages/{id} [put]
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

// @Summary Заменить Проект
// @Description Заменить проект и записать в базу данных
// @Tags Portfolio
// @Accept json
// @Param id path int true "Project ID"
// @Param title body string true "Project title"
// @Param version body string true "Project version"
// @Param description body string true "Project description"
// @Param language_id body int true "Language ID"
// @Param isActive body bool true "Is active"
// @Param isArchived body bool true "Is archived"
// @Param isDeveloping body bool true "Is developing"
// @Param GHLink body string false "GitHub link"
// @Param TGLink body string false "Telegram link"
// @Param HTTPLink body string false "HTTP link"
// @Produce json
// @Success 200 {object} models.Project "Информация о проекте"
// @Failure 400 {object} error "Невалидные данные"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
// @Router /portfolio/projects/{id} [put]
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

// @Summary Обновить ЯП
// @Description Обновить язык программирования и записать в базу данных
// @Tags Portfolio
// @Accept json
// @Param id path int true "Language ID"
// @Param name body string false "Language name"
// @Param svg body string flase "Language svg"
// @Produce json
// @Success 200 {object} models.Language "Информация о языке"
// @Failure 400 {object} error "Невалидные данные"
// @Failure 404 {object} error "ЯП не найден"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
// @Router /portfolio/languages/{id} [patch]
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

// @Summary Обновить Проект
// @Description Обновить проект и записать в базу данных
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
// @Param GHLink body string false "GitHub link"
// @Param TGLink body string false "Telegram link"
// @Param HTTPLink body string false "HTTP link"
// @Produce json
// @Success 200 {object} models.Project "Информация о проекте"
// @Failure 400 {object} error "Невалидные данные"
// @Failure 404 {object} error "Проект не найден"
// @Failure 500 {object} error "Внутренняя ошибка сервера"
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

	projectUpdated, err := pc.service.PatchProject(pc.ctx, project, &projectUpdate)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update project"})
		return
	}
	c.JSON(200, projectUpdated)
}
