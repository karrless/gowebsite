package models

import (
	"github.com/guregu/null/v5"
)

type Language struct {
	ID   int64       `form:"id" json:"id" db:"id"`
	Name string      `form:"name" json:"name" db:"name"`
	Svg  null.String `form:"svg" json:"svg" db:"svg"`
}

type Project struct {
	ID           int64       `form:"id" json:"id" db:"id"`
	Title        string      `form:"title" json:"title" db:"title"`
	Version      string      `form:"version" json:"version" db:"version"`
	Description  string      `form:"dscription" json:"dscription" db:"description"`
	LanguageID   int64       `form:"languageID" json:"languageID" db:"lang_id"`
	Language     *Language   `form:"language" json:"language" db:"-"`
	IsActive     null.Bool   `form:"isActive" json:"isActive" db:"is_active"`
	IsArchived   null.Bool   `form:"isArchived" json:"isArchived" db:"is_archived"`
	IsDeveloping null.Bool   `form:"isDeveloping" json:"isDeveloping" db:"is_developing"`
	GHLink       null.String `form:"GHLink" json:"GHLink" db:"gh_link"`
	TGLink       null.String `form:"TGLink" json:"TGLink" db:"tg_link"`
	HTTPLink     null.String `form:"HTTPLink" json:"HTTPLink" db:"http_link"`
}

type ProjectFilter struct {
	LanguageID   *int64 `form:"lang_id" db:"lang_id"`
	IsActive     *bool  `form:"is_active" db:"is_active"`
	IsArchived   *bool  `form:"is_archived" db:"is_archived"`
	IsDeveloping *bool  `form:"is_developing" db:"is_developing"`
	SortField    string `form:"sort_field" db:"sort_field"`
	SortOrder    string `form:"sort_order" db:"sort_order"`
	Limit        uint64 `form:"limit" db:"limit"`
	Offset       uint64 `form:"offset" db:"offset"`
}
