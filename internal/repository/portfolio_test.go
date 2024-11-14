package repository

import (
	"context"
	"gowebsite/internal/config"
	"gowebsite/internal/models"
	"gowebsite/pkg/db/postgres"
	"testing"

	"github.com/guregu/null/v5"
)

func TestNewPortfolioRepository(t *testing.T) {
	cfg := config.New("../../configs/.env")
	t.Logf("%+v\n", cfg)
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
}

func TestPortfolioRepository_ListLanguages(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	languages, err := repo.ListLanguages(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, l := range languages {
		t.Logf("%+v\n", l)
	}
}

func TestPortfolioRepository_GetLanguage(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	l, err := repo.GetLanguage(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if l == nil {
		t.Failed()
	}
	t.Logf("%+v\n", l)
	if l.ID != 1 {
		t.Fail()
	}
	l, err = repo.GetLanguage(context.Background(), 100)
	if err == nil {
		t.Fail()
	}
}

func TestPortfolioRepository_DeleteLanguage(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	err = repo.DeleteLanguage(context.Background(), 3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPortfolioRepository_CreateLanguage(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	language := &models.Language{
		Name: "name",
		Svg:  null.StringFrom("svg"),
	}
	l, err := repo.CreateLanguage(context.Background(), language)
	if err != nil {
		t.Fatal(err)
	}
	if l == nil {
		t.Failed()
	}
	t.Logf("%+v\n", l)
}

func TestPortfolioRepository_UpdateLanguage(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	language := &models.Language{
		ID:   1,
		Name: "name",
		Svg:  null.StringFrom("svg"),
	}
	l, err := repo.UpdateLanguage(context.Background(), language)
	if err != nil {
		t.Fatal(err)
	}
	if l.Name != "name" {
		t.Fail()
	}
	l, err = repo.GetLanguage(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if l.Name != "name" {
		t.Fail()
	}
	t.Logf("%+v\n", l)
}

func TestPortfolioRepository_CreateProject(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	ghLink := "test"
	project := &models.Project{
		Title:        "title",
		Version:      "version",
		Description:  "description",
		IsActive:     true,
		IsArchived:   true,
		IsDeveloping: true,
		GHLink:       null.StringFrom(ghLink),
		LanguageID:   1,
	}

	p, err := repo.CreateProject(context.Background(), project)
	if err != nil {
		t.Fatal(err)
	}
	if p == nil {
		t.Failed()
	}
	if p.Title != "title" {
		t.Fail()
	}
	project.LanguageID = 100
	t.Logf("%+v\n", p)
	p, err = repo.CreateProject(context.Background(), project)
	if err == nil {
		t.Fatal(err)
	}
}
func TestPortfolioRepository_GetProject(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	p, err := repo.GetProject(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if p == nil || p.ID != 1 {
		t.Failed()
	}
	t.Logf("%+v\n", p)
	t.Logf("%+v\n", p.Language)
	p, err = repo.GetProject(context.Background(), 100)
	if err == nil {
		t.Fatal(err)
	}
}

func TestPortfolioRepository_ListProjects(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	filter := &models.ProjectFilter{
		SortField: "lang_id",
		SortOrder: "desc",
	}
	projects, err := repo.ListProjects(context.Background(), filter)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range projects {
		t.Logf("%+v\n", p)
	}
}

func TestPortfolioRepository_UpdateProject(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	project := &models.Project{
		ID:           1,
		Title:        "title",
		Version:      "version",
		Description:  "description",
		IsActive:     true,
		IsArchived:   true,
		IsDeveloping: true,
		GHLink:       null.StringFrom("test"),
		LanguageID:   1,
	}
	err = repo.UpdateProject(context.Background(), project)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPortfolioRepository_DeleteProject(t *testing.T) {
	cfg := config.New("../../configs/.env")
	db, err := postgres.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPortfolioRepository(db)
	if repo == nil {
		t.Fatal("repository.NewPortfolioRepository returned nil")
	}
	err = repo.DeleteProject(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
}
