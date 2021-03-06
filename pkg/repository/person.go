package repository

import (
	"github.com/Rosaniline/gorm-ut/pkg/model"
	uuid "github.com/satori/go.uuid"

	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(id uuid.UUID, name string) error
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) Create(id uuid.UUID, name string) error {
	person := &model.Person{
		ID:   id,
		Name: name,
	}
	return p.DB.Create(person).Error
}

func CreateRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

// type Repository interface {
// 	Create(id int, name string) error
// }

// type repo struct {
// 	DB *gorm.DB
// }

// func (p *repo) Create(id int, name string) error {
// 	person := &model.Person{
// 		ID:   id,
// 		Name: name,
// 	}
// 	return p.DB.Create(person).Error
// }

// func CreateRepository(db *gorm.DB) Repository {
// 	return &repo{
// 		DB: db,
// 	}
// }
