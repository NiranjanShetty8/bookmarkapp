package model

import uuid "github.com/satori/go.uuid"

//Represents a bookmark
type Bookmark struct {
	Base
	Name       string    `gorm:"type:varchar(30);unique;not null" json:"name"`
	URL        string    `gorm:"not null" json:"url"`
	CategoryID uuid.UUID `gorm:"type:varchar(36);not null" json:"categoryID"`
}
