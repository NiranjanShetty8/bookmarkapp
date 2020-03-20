package model

import uuid "github.com/satori/go.uuid"

type Bookmark struct {
	Base
	Description string `gorm:type:varchar(200)"`
	URL         string `gorm:type:varchar(200)"`
	// UserID      uuid.UUID `gorm:"type:varchar(36);not_null"`
	CategoryID uuid.UUID `gorm:"type:varchar(36);not_null"`
}
