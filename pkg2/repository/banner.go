package repository

import (
	"time"

	"github.com/Rosaniline/gorm-ut/pkg2/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RepositoryBanner interface {
	SaveBanner2(id uuid.UUID, fakertime time.Time) error
}

type repositoryBanner struct {
	db *gorm.DB
}

func (r *repositoryBanner) SaveBanner2(id uuid.UUID, fakertime time.Time) error {
	mappedBannerToModel := &model.Banner{
		ID:               id,
		BannerCategoryID: 1,
		Status:           1,
		FileName:         "123-gambar.png",
		BannerOrder:      1,
		Link:             "https://google.com",
		CreatedById:      1234,
		CreatedByName:    "indra",
		CreatedAt:        fakertime,
		UpdatedById:      1234,
		UpdatedByName:    "indra",
		UpdatedAt:        fakertime,
		DeletedById:      1234,
		DeletedByName:    "indra",
		DeletedAt:        fakertime,
	}

	// Save To Database
	return r.db.Create(mappedBannerToModel).Error

}

func NewRepositoryBanner(db *gorm.DB) RepositoryBanner {
	return &repositoryBanner{
		db: db,
	}
}
