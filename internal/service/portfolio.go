package service

import (
	"context"
	"gowebsite/internal/models"
)

type OrderRepo interface {
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

type PortfolioService struct {
	portfolioRepo OrderRepo
}

func NewPortfolioService(repo OrderRepo) *PortfolioService {
	return &PortfolioService{portfolioRepo: repo}
}

func (r *PortfolioService) GetLanguage(ctx context.Context, id int64) (*models.Language, error) {
	return r.portfolioRepo.GetLanguage(ctx, id)
}

func (r *PortfolioService) GetProject(ctx context.Context, id int64) (*models.Project, error) {
	return r.portfolioRepo.GetProject(ctx, id)
}

func (r *PortfolioService) ListLanguages(ctx context.Context) ([]*models.Language, error) {
	return r.portfolioRepo.ListLanguages(ctx)
}

func (r *PortfolioService) ListProjects(ctx context.Context, filter *models.ProjectFilter) ([]*models.Project, error) {
	return r.portfolioRepo.ListProjects(ctx, filter)
}

func (r *PortfolioService) CreateLanguage(ctx context.Context, language *models.Language) (*models.Language, error) {
	return r.portfolioRepo.CreateLanguage(ctx, language)
}

func (r *PortfolioService) CreateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	return r.portfolioRepo.CreateProject(ctx, project)
}

func (r *PortfolioService) UpdateLanguage(ctx context.Context, language *models.Language) (*models.Language, error) {
	return r.portfolioRepo.UpdateLanguage(ctx, language)
}

func (r *PortfolioService) UpdateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	return r.portfolioRepo.UpdateProject(ctx, project)
}

func (r *PortfolioService) PatchLanguage(ctx context.Context, language *models.Language, languageUpdate *models.Language) (*models.Language, error) {
	return r.portfolioRepo.PatchLanguage(ctx, language, languageUpdate)
}

func (r *PortfolioService) PatchProject(ctx context.Context, project *models.Project, projectUpdate *models.Project) (*models.Project, error) {
	return r.portfolioRepo.PatchProject(ctx, project, projectUpdate)
}

func (r *PortfolioService) DeleteLanguage(ctx context.Context, id int64) error {
	return r.portfolioRepo.DeleteLanguage(ctx, id)
}

func (r *PortfolioService) DeleteProject(ctx context.Context, id int64) error {
	return r.portfolioRepo.DeleteProject(ctx, id)
}
