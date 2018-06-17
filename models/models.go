package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type BaseModel struct {
	ID        int64
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}

func (m *BaseModel) BeforeInsert(db orm.DB) error {
	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return nil
}

func (m *BaseModel) BeforeUpdate(db orm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

type Post struct {
	BaseModel
	Creator string `sql:"type:varchar(255)" json:"creator" binding:"required"`
	Content string `json:"content"`
}
