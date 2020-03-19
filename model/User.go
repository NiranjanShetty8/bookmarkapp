package model

type User struct {
	Base
	Username  string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	Bookmarks []Bookmark `json:"-"`
	// Categories *[]Category `json:"-"`
}
