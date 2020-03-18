package model

type User struct {
	Base
	Nickname string               `gorm:"type:varchar(100)"`
	Username string               `gorm:"unique;not null"`
	Password string               `gorm:"not null"`
	Bookmark *[]bookmark.Bookmark `json:"-"`
}
