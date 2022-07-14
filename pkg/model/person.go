package model

import uuid "github.com/satori/go.uuid"

// import "github.com/google/uuid"

type Person struct {
	// ID   uuid.UUID `gorm:"column:id;primary_key" json:"id"`
	ID uuid.UUID
	// ID   int
	Name string
}

func (p *Person) TableName() string {
	return "person"
}
