package model

//Number of attempts allowed with a wrong password
//It is 1 more than the acutal number because of the default value update problem in gorm
//eg:= if loginAttempts = 4, the user gets 3 attempts
const loginAttempts = 3

// Represents the user
type User struct {
	Base
	Username      string     `gorm:"unique;not null" json:"username"`
	Password      string     `gorm:"not null" json:"password"`
	LoginAttempts int        `gorm:"type:integer(1);not null;DEFAULT:3" json:"-"`
	Categories    []Category `json:"-"`
}

func GetLoginAttempts() int {
	return loginAttempts
}
