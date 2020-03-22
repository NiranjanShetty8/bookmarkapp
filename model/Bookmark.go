package model

import uuid "github.com/satori/go.uuid"

//Represents a bookmark
type Bookmark struct {
	Base
	Name       string    `gorm:type:varchar(200)"`
	URL        string    `gorm:type:varchar(200)"`
	CategoryID uuid.UUID `gorm:"type:varchar(36);not_null"`
}
