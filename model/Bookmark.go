package model

import uuid "github.com/satori/go.uuid"

type Bookmark struct {
	Base
	Description string    `gorm:type:varchar(200)"`
	URL         string    `gorm:type:varchar(200)"`
	UserId      uuid.UUID `gorm:"type:varchar(36);not_null"`
	CategoryId  uuid.UUID `gorm:"type:varchar(36);not_null"`
}
