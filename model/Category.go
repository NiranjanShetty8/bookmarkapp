package model

import uuid "github.com/satori/go.uuid"

// Represents category
type Category struct {
	Base
	Name      string     `gorm:"not null;type:varchar(30)" json:"name"`
	UserID    uuid.UUID  `gorm:"type:varchar(36);not null" json:"userID"`
	Bookmarks []Bookmark `json:"bookmarks"`
}
