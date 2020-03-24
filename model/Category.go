package model

import uuid "github.com/satori/go.uuid"

// Represents category
type Category struct {
	Base
	Name      string     `gorm:"type:varchar(40)"`
	UserID    uuid.UUID  `gorm:"type:varchar(36);not null"`
	Bookmarks []Bookmark `json:"bookmarks"`
}
