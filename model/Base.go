package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// A base type for tables which will have these columns
type Base struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key;" json:"id"`
	CreatedAt time.Time  `gorm:"column:createdOn" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedOn" json:"-"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deletedOn" json:"-"`
}
