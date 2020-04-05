package model

//Number of attempts allowed with a wrong password
const loginAttempts = 3

// Represents the user
type User struct {
	Base
	Username       string      `gorm:"unique;not null;type:varchar(30)" json:"username"`
	Password       string      `gorm:"not null" json:"password"`
	Email          string      `gorm:"unique;type:varchar(40)" json:"email,omitempty"`
	ProfilePicture interface{} `gorm:"type:mediumblob" json:"profilePicture"`
	SuperUser      bool        `gorm:"type:varchar(5);DEFAULT:'FALSE'" json:"superUser,boolean"`
	LoginAttempts  int         `gorm:"type:integer(1);not null;DEFAULT:3" json:"loginAttempts"`
	Categories     []Category  `json:"categories"`
}

func GetLoginAttempts() int {
	return loginAttempts
}
