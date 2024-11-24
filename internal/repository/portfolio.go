package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gowebsite/internal/models"
	"gowebsite/pkg/db/postgres"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"github.com/volatiletech/null/v9"
)

type PortfolioRepository struct {
	*postgres.DB
}

func NewPortfolioRepository(db *postgres.DB) *PortfolioRepository {
	return &PortfolioRepository{db}
}

func (repo *PortfolioRepository) CreateTechonology(ctx context.Context, technology *models.Technology) (int64, error) {
	var resultID int64
	err := sq.Insert("techs").
		Columns("name", "svg").
		Values(technology.Name, technology.Svg).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).
		QueryRow().Scan(&resultID)
	if err != nil {
		return 0, fmt.Errorf("repository.CreateTechnology: %v", err)
	}
	return resultID, nil
}

func (repo *PortfolioRepository) GetTechnology(ctx context.Context, id int64) (*models.Technology, error) {
	var result models.Technology
	err := sq.Select("*").
		From("techs").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).
		QueryRow().Scan(&result.ID, &result.Name, &result.Svg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("repository.GetTechnology: %v", err)
	}
	return &result, nil
}

func (repo *PortfolioRepository) ListTechnologies(ctx context.Context, filter *models.TechnologyFilter) ([]*models.Technology, error) {
	var result []*models.Technology

	query := sq.Select("*").From("techs").PlaceholderFormat(sq.Dollar)
	query = query.Where(sq.Eq{"id": *filter.TechnologiesID})

	if filter.SortField != "" {
		if filter.SortOrder == "" {
			filter.SortOrder = "ASC"
		}
		query = query.OrderBy(fmt.Sprintf("%s %s", filter.SortField, filter.SortOrder),
			"techs.name ASC")
	}
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	rows, err := query.RunWith(repo.DB.DB).Query()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return result, fmt.Errorf("repository.ListTechnologies: %v", err)
	}
	for rows.Next() {
		var technology models.Technology
		if err := rows.Scan(&technology.ID, &technology.Name, &technology.Svg); err != nil {
			return nil, fmt.Errorf("repository.ListTechnologies: %v", err)
		}
		result = append(result, &technology)
	}
	return result, nil
}

func (repo *PortfolioRepository) DeleteTechnology(ctx context.Context, id int64) error {
	_, err := sq.Delete("techs").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).Query()
	if err != nil {

		return fmt.Errorf("repository.DeleteTechnology: %v", err)
	}
	return nil
}

func (repo *PortfolioRepository) CreateProject(ctx context.Context, project *models.Project) (int64, error) {
	var resultID int64
	err := sq.Insert("projects").
		Columns("title", "version", "description", "is_active", "is_archived", "is_developing", "gh_link", "tg_link", "http_link").
		Values(project.Title, project.Version, project.Description, project.IsActive, project.IsArchived, project.IsDeveloping, project.Links).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).
		QueryRow().Scan(&resultID)
	if err != nil {
		return 0, fmt.Errorf("repository.CreateProject: %v", err)
	}

	query := sq.Insert("project_tech").Columns("project_id", "tech_id").PlaceholderFormat(sq.Dollar)

	for _, Technology := range project.Technologies {
		query = query.Values(resultID, Technology.ID)
	}

	_, err = query.RunWith(repo.DB.DB).Exec()
	if err != nil {
		return 0, fmt.Errorf("repository.CreateProject: %v", err)
	}
	return resultID, nil
}

func (repo *PortfolioRepository) GetProject(ctx context.Context, id int64) (*models.Project, error) {

	query := sq.Select("p.id, p.title, p.version, p.description, p.is_active, p.is_archived, p.is_developing, p.links, t.id AS tech_id, t.name AS tech_name, t.svg AS tech_svg").
		From("projects p").
		Join("project_tech pt ON p.id = pt.project_id").
		Join("techs t ON pt.tech_id = t.id").
		Where(sq.Eq{"p.id": id}).
		PlaceholderFormat(sq.Dollar)

	rows, err := query.RunWith(repo.DB.DB).Query()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("repository.GetProject: %v", err)
	}
	var result models.Project
	result.Technologies = []*models.Technology{}
	for rows.Next() {
		var technology models.Technology
		err := rows.Scan(&result.ID, &result.Title, &result.Version, &result.Description, &result.IsActive, &result.IsArchived, &result.IsDeveloping, &result.Links, &technology.ID, &technology.Name, &technology.Svg)
		if err != nil {
			return nil, fmt.Errorf("repository.GetProject: %v", err)
		}
		result.Technologies = append(result.Technologies, &technology)
	}

	return &result, nil
}

func (repo *PortfolioRepository) ListProjects(ctx context.Context, filter *models.ProjectFilter) ([]*models.Project, error) {
	var result []*models.Project
	query := sq.Select("p.id, p.title, p.version, p.description, p.is_active, p.is_archived, p.is_developing, p.links, t.id AS tech_id, t.name AS tech_name, t.svg AS tech_svg").
		From("projects p").
		Join("project_tech pt ON p.id = pt.project_id").
		Join("techs t ON pt.tech_id = t.id").
		PlaceholderFormat(sq.Dollar)

	if filter.TechnologiesID != nil {
		query = query.Where(sq.Eq{"t.id": filter.TechnologiesID})
	}
	if filter.IsActive != nil {
		query = query.Where(sq.Eq{"is_active": *filter.IsActive})
	}
	if filter.IsArchived != nil {
		query = query.Where(sq.Eq{"is_archived": *filter.IsArchived})
	}
	if filter.IsDeveloping != nil {
		query = query.Where(sq.Eq{"is_developing": *filter.IsDeveloping})
	}
	query = query.PlaceholderFormat(sq.Dollar)

	if filter.SortField != "" {
		if filter.SortOrder == "" {
			filter.SortOrder = "ASC"
		}
		query = query.OrderBy(
			fmt.Sprintf("%s %s", filter.SortField, filter.SortOrder),
			"projects.title ASC",
		)
	}

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}
	rows, err := query.RunWith(repo.DB.DB).Query()
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return nil, fmt.Errorf("repository.ListProjects: %v", err)
	}

	var currentProject *models.Project
	for rows.Next() {
		var projectID, technologyID int64
		var projectTitle, projectVersion, projectDescription, techName string
		var projectIsActive, projectIsArchived, projectIsDeveloping null.Bool
		var projectLinks pq.StringArray
		var techSvg null.String

		err := rows.Scan(&projectID, &projectTitle, &projectVersion, &projectDescription, &projectIsActive, &projectIsArchived, &projectIsDeveloping, &projectLinks, &technologyID, &techName, &techSvg)
		if err != nil {
			return nil, fmt.Errorf("repository.ListProjects: %v", err)
		}

		if currentProject == nil || currentProject.ID != projectID {
			if currentProject != nil {
				result = append(result, currentProject)
			}
			currentProject = &models.Project{
				ID:           projectID,
				Title:        projectTitle,
				Version:      projectVersion,
				Description:  projectDescription,
				IsActive:     projectIsActive,
				IsArchived:   projectIsArchived,
				IsDeveloping: projectIsDeveloping,
				Links:        projectLinks,
			}
		}

		currentProject.Technologies = append(currentProject.Technologies,
			&models.Technology{
				ID:   technologyID,
				Name: techName,
				Svg:  techSvg,
			})
	}
	result = append(result, currentProject)
	return result, nil
}

func (repo *PortfolioRepository) DeleteProject(ctx context.Context, id int64) error {
	_, err := sq.Delete("projects").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).
		Query()
	if err != nil {
		return fmt.Errorf("repository.DeleteProject: %v", err)
	}
	return nil
}

func (repo *PortfolioRepository) PatchTechnology(ctx context.Context, technology *models.Technology) error {

	query := sq.Update("techs").Where(sq.Eq{"id": technology.ID}).PlaceholderFormat(sq.Dollar)

	if technology.Name != "" {
		query = query.Set("name", technology.Name)
	}

	if technology.Svg.Valid {
		query = query.Set("svg", technology.Svg)
	}

	_, err := query.RunWith(repo.DB.DB).Exec()
	if err != nil {
		return fmt.Errorf("repository.PatchTechnology: %v", err)
	}
	return nil
}

// TODO:
func (repo *PortfolioRepository) PatchProject(ctx context.Context, project *models.Project, projectUpdate *models.Project) error {
	query := sq.Update("projects").Where(sq.Eq{"id": project.ID}).PlaceholderFormat(sq.Dollar)

	if projectUpdate.Title != "" {
		query = query.Set("title", projectUpdate.Title)
	}

	if projectUpdate.Version != "" {
		query = query.Set("version", projectUpdate.Version)
	}

	if projectUpdate.Description != "" {
		query = query.Set("description", projectUpdate.Description)
	}
	if projectUpdate.IsActive.Valid {
		query = query.Set("is_active", projectUpdate.IsActive)
	}

	if projectUpdate.IsArchived.Valid {
		query = query.Set("is_archived", projectUpdate.IsArchived)
	}

	if projectUpdate.IsDeveloping.Valid {
		query = query.Set("is_developing", projectUpdate.IsDeveloping)
	}

	if projectUpdate.Links != nil {
		query = query.Set("gh_link", projectUpdate.Links)
	}

	_, err := query.RunWith(repo.DB.DB).Exec()

	if err != nil {
		return fmt.Errorf("repository.UpdateProject: %v", err)
	}
	deleteQuery := sq.Delete("project_tech").Where(sq.Eq{"project_id": project.ID}).PlaceholderFormat(sq.Dollar)
	createQuery := sq.Insert("project_tech").Columns("project_id", "tech_id").PlaceholderFormat(sq.Dollar)
	if projectUpdate.Technologies != nil {
		for _, technology := range projectUpdate.Technologies {
			createQuery = createQuery.Values(project.ID, technology.ID)
		}
		_, err := deleteQuery.RunWith(repo.DB.DB).Exec()
		if err != nil {
			return fmt.Errorf("repository.UpdateProject: %v", err)
		}
		_, err = createQuery.RunWith(repo.DB.DB).Exec()
		if err != nil {
			return fmt.Errorf("repository.UpdateProject: %v", err)
		}
	}

	return nil
}
