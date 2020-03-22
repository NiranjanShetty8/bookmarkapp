package model

// Represents the user
type User struct {
	Base
	Username   string     `gorm:"unique;not null"`
	Password   string     `gorm:"not null"`
	Categories []Category `json:"-"`
}
