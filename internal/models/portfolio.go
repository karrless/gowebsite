package models

import (
	"github.com/volatiletech/null/v9"
)

// Technology model
type Technology struct {
	ID   int64       `form:"id" json:"id" db:"id"`
	Name string      `form:"name" json:"name" db:"name"`
	Svg  null.String `form:"svg" json:"svg" db:"svg" swaggertype:"string"`
}

type Project struct {
	ID            int64         `form:"id" json:"id" db:"id"`
	Title         string        `form:"title" json:"title" db:"title"`
	Version       string        `form:"version" json:"version" db:"version"`
	Description   string        `form:"dscription" json:"dscription" db:"description"`
	TechnologyIDs []int64       `form:"tech_id" json:"tech_id" db:"-"`
	Technologies  []*Technology `form:"technologies" json:"technologies" db:"-"`
	IsActive      null.Bool     `form:"isActive" json:"isActive" db:"is_active" swaggertype:"boolean"`
	IsArchived    null.Bool     `form:"isArchived" json:"isArchived" db:"is_archived" swaggertype:"boolean"`
	IsDeveloping  null.Bool     `form:"isDeveloping" json:"isDeveloping" db:"is_developing" swaggertype:"boolean"`
	Links         []string      `form:"links" json:"links" db:"links"`
}

type ProjectFilter struct {
	TechnologiesID *[]int64 `form:"tech_id" db:"tech_id"`
	IsActive       *bool    `form:"is_active" db:"is_active"`
	IsArchived     *bool    `form:"is_archived" db:"is_archived"`
	IsDeveloping   *bool    `form:"is_developing" db:"is_developing"`
	SortField      string   `form:"sort_field" db:"sort_field"`
	SortOrder      string   `form:"sort_order" db:"sort_order"`
	Limit          uint64   `form:"limit" db:"limit"`
	Offset         uint64   `form:"offset" db:"offset"`
}

type TechnologyFilter struct {
	TechnologiesID *[]int64 `form:"tech_id" db:"tech_id"`
	SortField      string   `form:"sort_field" db:"sort_field"`
	SortOrder      string   `form:"sort_order" db:"sort_order"`
	Limit          uint64   `form:"limit" db:"limit"`   //nolint:tagliatelle
	Offset         uint64   `form:"offset" db:"offset"` //nolint:tagliatelle
}
