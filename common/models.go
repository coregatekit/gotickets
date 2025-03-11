package common

import (
	"time"

	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = cuid2.Generate()
	return nil
}
