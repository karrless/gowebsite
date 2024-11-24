package service

import (
	"context"
	"gowebsite/internal/models"
)

type OrderRepo interface {
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

type PortfolioService struct {
	portfolioRepo OrderRepo
}

func NewPortfolioService(repo OrderRepo) *PortfolioService {
	return &PortfolioService{portfolioRepo: repo}
}

func (s *PortfolioService) CreateTechnology(ctx context.Context, technology *models.Technology) (int64, error) {
	return s.portfolioRepo.CreateTechnology(ctx, technology)
}

func (s *PortfolioService) GetTechnology(ctx context.Context, id int64) (*models.Technology, error) {
	return s.portfolioRepo.GetTechnology(ctx, id)
}

func (s *PortfolioService) ListTechnologies(ctx context.Context, filter *models.TechnologyFilter) ([]*models.Technology, error) {
	return s.portfolioRepo.ListTechnologies(ctx, filter)
}

func (s *PortfolioService) DeleteTechnology(ctx context.Context, id int64) error {
	return s.portfolioRepo.DeleteTechnology(ctx, id)
}

func (s *PortfolioService) PatchTechnology(ctx context.Context, technology *models.Technology) error {
	return s.portfolioRepo.PatchTechnology(ctx, technology)
}

func (s *PortfolioService) CreateProject(ctx context.Context, project *models.Project) (int64, error) {

	return s.portfolioRepo.CreateProject(ctx, project)
}

func (s *PortfolioService) GetProject(ctx context.Context, id int64) (*models.Project, error) {
	return s.portfolioRepo.GetProject(ctx, id)
}

func (s *PortfolioService) ListProjects(ctx context.Context, filter *models.ProjectFilter) ([]*models.Project, error) {
	return s.portfolioRepo.ListProjects(ctx, filter)
}

func (s *PortfolioService) DeleteProject(ctx context.Context, id int64) error {
	return s.portfolioRepo.DeleteProject(ctx, id)
}

func (s *PortfolioService) PatchProject(ctx context.Context, project *models.Project, projectUpdate *models.Project) error {
	return s.portfolioRepo.PatchProject(ctx, project, projectUpdate)
}
