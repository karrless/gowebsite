package models

import (
	"github.com/guregu/null/v5"
)

type Language struct {
	ID   int64       `json:"ID" db:"id"`
	Name string      `json:"Name" db:"name"`
	Svg  null.String `json:"Svg" db:"svg"`
}

type Project struct {
	ID           int64       `json:"ID" db:"id"`
	Title        string      `json:"Title" db:"title"`
	Version      string      `json:"Version" db:"version"`
	Description  string      `json:"Description" db:"description"`
	LanguageID   int64       `json:"LanguageID" db:"lang_id"`
	Language     *Language   `json:"Language" db:"-"`
	IsActive     bool        `json:"IsActive" db:"is_active"`
	IsArchived   bool        `json:"IsArchived" db:"is_archived"`
	IsDeveloping bool        `json:"IsDeveloping" db:"is_developing"`
	GHLink       null.String `json:"GHLink" db:"gh_link"`
	TGLink       null.String `json:"TGLink" db:"tg_link"`
	HTTPLink     null.String `json:"HTTPLink" db:"http_link"`
}

type ProjectFilter struct {
	LanguageID   *int64 `json:"LanguageID" db:"lang_id"`
	IsActive     *bool  `json:"IsActive" db:"is_active"`
	IsArchived   *bool  `json:"IsArchived" db:"is_archived"`
	IsDeveloping *bool  `json:"IsDevelop" db:"is_developing"`
	SortField    string `json:"SortField" db:"sort_field"`
	SortOrder    string `json:"SortOrder" db:"sort_order"`
	Limit        uint64 `json:"Limit" db:"limit"`
	Offset       uint64 `json:"Offset" db:"offset"`
}
