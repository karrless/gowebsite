package repository

import (
	"context"
	"fmt"
	"gowebsite/internal/models"
	"gowebsite/pkg/db/postgres"

	sq "github.com/Masterminds/squirrel"
)

type PortfolioRepository struct {
	*postgres.DB
}

func NewPortfolioRepository(db *postgres.DB) *PortfolioRepository {
	return &PortfolioRepository{db}
}

func (repo *PortfolioRepository) CreateLanguage(ctx context.Context, language *models.Language) (*models.Language, error) {
	var result models.Language
	err := sq.Insert("languages").
		Columns("name", "svg").
		Values(language.Name, language.Svg).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).
		QueryRow().Scan(&result.ID, &result.Name, &result.Svg)
	if err != nil {
		return nil, fmt.Errorf("repository.CreateLanguage: %v", err)
	}
	return &result, nil
}

func (repo *PortfolioRepository) GetLanguage(ctx context.Context, id int64) (*models.Language, error) {
	var result models.Language
	err := sq.Select("*").
		From("languages").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).
		QueryRow().Scan(&result.ID, &result.Name, &result.Svg)
	if err != nil {
		return nil, fmt.Errorf("repository.GetLanguage: %v", err)
	}
	return &result, nil
}

func (repo *PortfolioRepository) ListLanguages(ctx context.Context) ([]*models.Language, error) {
	var result []*models.Language
	rows, err := sq.Select("*").
		From("languages").
		RunWith(repo.DB.DB).Query()
	if err != nil {
		return nil, fmt.Errorf("repository.GetLanguages: %v", err)
	}
	for rows.Next() {
		var language models.Language
		if err := rows.Scan(&language.ID, &language.Name, &language.Svg); err != nil {
			return nil, fmt.Errorf("repository.GetLanguages: %v", err)
		}
		result = append(result, &language)
	}
	return result, nil
}

func (repo *PortfolioRepository) UpdateLanguage(ctx context.Context, language *models.Language) (*models.Language, error) {
	var result models.Language
	err := sq.Update("languages").
		Set("name", language.Name).
		Set("svg", language.Svg).
		Where(sq.Eq{"id": language.ID}).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).
		QueryRow().Scan(&result.ID, &result.Name, &result.Svg)
	if err != nil {
		return nil, fmt.Errorf("repository.UpdateLanguage: %v", err)
	}
	return &result, nil
}

func (repo *PortfolioRepository) DeleteLanguage(ctx context.Context, id int64) error {
	_, err := sq.Delete("languages").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).Query()
	if err != nil {
		return fmt.Errorf("repository.DeleteLanguage: %v", err)
	}
	return nil
}

func (repo *PortfolioRepository) CreateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	var result models.Project
	err := sq.Insert("projects").
		Columns("title", "version", "description", "is_active", "is_archived", "is_developing", "gh_link", "tg_link", "http_link", "lang_id").
		Values(project.Title, project.Version, project.Description, project.IsActive, project.IsArchived, project.IsDeveloping, project.GHLink, project.TGLink, project.HTTPLink, project.LanguageID).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).
		QueryRow().Scan(&result.ID, &result.Title, &result.Version, &result.Description, &result.LanguageID, &result.IsActive, &result.IsArchived, &result.IsDeveloping, &result.GHLink, &result.TGLink, &result.HTTPLink)
	if err != nil {
		return nil, fmt.Errorf("repository.CreateProject: %v", err)
	}
	lang, err := repo.GetLanguage(ctx, result.LanguageID)
	if err != nil {
		return nil, fmt.Errorf("repository.CreateProject: %v", err)
	}
	result.Language = lang
	return &result, nil
}

func (repo *PortfolioRepository) GetProject(ctx context.Context, id int64) (*models.Project, error) {
	var result models.Project
	var lang models.Language
	err := sq.Select("*").
		From("projects").
		Where(sq.Eq{"projects.id": id}).
		PlaceholderFormat(sq.Dollar).
		InnerJoin("languages ON projects.lang_id = languages.id").
		RunWith(repo.DB.DB).
		QueryRow().
		Scan(&result.ID, &result.Title, &result.Version, &result.Description, &result.LanguageID, &result.IsActive, &result.IsArchived, &result.IsDeveloping, &result.GHLink, &result.TGLink, &result.HTTPLink, &lang.ID, &lang.Name, &lang.Svg)
	if err != nil {
		return nil, fmt.Errorf("repository.GetProject: %v", err)
	}
	result.Language = &lang
	return &result, nil
}

func (repo *PortfolioRepository) ListProjects(ctx context.Context, filter *models.ProjectFilter) ([]*models.Project, error) {
	var result []*models.Project
	query := sq.Select("*").From("projects").InnerJoin("languages ON projects.lang_id = languages.id")

	if filter.LanguageID != nil {
		query = query.Where(sq.Eq{"lang_id": *filter.LanguageID})
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
		return nil, fmt.Errorf("repository.ListProjects: %v", err)
	}
	for rows.Next() {
		var project models.Project
		var lang models.Language
		if err := rows.Scan(&project.ID, &project.Title, &project.Version, &project.Description, &project.LanguageID, &project.IsActive, &project.IsArchived, &project.IsDeveloping, &project.GHLink, &project.TGLink, &project.HTTPLink, &lang.ID, &lang.Name, &lang.Svg); err != nil {
			return nil, fmt.Errorf("repository.ListProjects: %v", err)
		}

		project.Language = &lang
		result = append(result, &project)
	}
	return result, nil
}

func (repo *PortfolioRepository) UpdateProject(ctx context.Context, project *models.Project) error {
	_, err := sq.Update("projects").
		Set("title", project.Title).
		Set("version", project.Version).
		Set("description", project.Description).
		Set("is_active", project.IsActive).
		Set("is_archived", project.IsArchived).
		Set("is_developing", project.IsDeveloping).
		Set("gh_link", project.GHLink).
		Set("tg_link", project.TGLink).
		Set("http_link", project.HTTPLink).
		Where(sq.Eq{"id": project.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB.DB).
		Exec()
	if err != nil {
		return fmt.Errorf("repository.UpdateProject: %v", err)
	}
	return nil
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
