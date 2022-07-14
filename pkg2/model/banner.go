package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Banner struct {
	ID               uuid.UUID
	BannerCategoryID int
	Status           int
	FileName         string
	BannerOrder      int
	Link             string
	CreatedById      int
	CreatedByName    string
	CreatedAt        time.Time
	UpdatedById      int
	UpdatedByName    string
	UpdatedAt        time.Time
	DeletedById      int
	DeletedByName    string
	DeletedAt        time.Time
}

func (p *Banner) TableName() string {
	return "banners"
}
