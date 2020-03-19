package model

import uuid "github.com/satori/go.uuid"

type Category struct {
	Base
	Name      string     `gorm:"type:varchar(40)"`
	UserId    uuid.UUID  `gorm:"type:varchar(36)"`
	Bookmarks []Bookmark `json:"-"`
}
