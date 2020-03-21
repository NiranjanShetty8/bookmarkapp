package repository

import (
	"fmt"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

//Unit of Work represents a transaction
type UnitOfWork struct {
	DB        *gorm.DB
	readOnly  bool
	committed bool
}

//NewUnitOfWork creates new UnitOfWork
func NewUnitOfWork(db *gorm.DB, readonly bool) *UnitOfWork {
	if readonly {
		return &UnitOfWork{
			DB:        db.New(),
			readOnly:  true,
			committed: false,
		}
	}
	return &UnitOfWork{
		DB:        db.New().Begin(),
		readOnly:  false,
		committed: false,
	}
}

//Complete function RollBacks transaction
func (uow *UnitOfWork) Complete() {
	if !uow.readOnly && !uow.committed {
		uow.DB.Rollback()
	}
}

//Commit function commits the transaction
func (uow *UnitOfWork) Commit() {
	if !uow.readOnly && !uow.committed {
		uow.DB.Commit()
	}
}

//Repository represents generic interface for interacting with DB
type Repository interface {
	GetAll(uow *UnitOfWork, uid uuid.UUID, out interface{}, preloadAssociations []string) error
	Get(uow *UnitOfWork, userId, bookmarkID uuid.UUID, out interface{},
		preloadAssociations []string) error
	GetByName(uow *UnitOfWork, name string, out interface{},
		preloadAssociations []string) error
	Add(uow *UnitOfWork, input interface{}) error
	Delete(uow *UnitOfWork, userId, bookmarkID uuid.UUID, out interface{}) error
	Update(uow *UnitOfWork, uid uuid.UUID, entity interface{}) error
}

//GormRepository implements Repository
type GormRepository struct{}

func (repos *GormRepository) GetAll(uow *UnitOfWork, uid uuid.UUID, out interface{},
	preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = uow.DB.Preload(association)
	}
	switch out.(type) {
	case *model.User:
		return db.Model(out).Find(out).Error
	case *model.Category:
		return db.Model(out).Find(out, "user_id = ?", uid).Error
	default:
		return db.Model(out).Find(out, "category_id = ?", uid).Error
	}
}

func (repos *GormRepository) Get(uow *UnitOfWork, parentID, childID uuid.UUID, out interface{},
	preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association).Where("id = ?", parentID)
	}
	switch out.(type) {
	case *model.User:
		return db.Model(out).First(out, "id = ?", parentID).Error
	case *model.Category:
		return db.Model(out).First(out, "user_id = ? AND id = ?", parentID, childID).Error
	default:
		return db.Model(out).First(out, "category_id = ? AND id = ?", parentID, childID).Error
	}
}

func (repos *GormRepository) GetByName(uow *UnitOfWork, name string, out interface{},
	preloadAssociations []string) error {
	db := uow.DB
	switch out.(type) {
	case *model.User:
		fmt.Println("in user switch")
		for _, association := range preloadAssociations {
			db = db.Preload(association).Where("username = ?", name)
		}
		return db.Model(out).First(out, "username = ?", name).Error

	case *model.Category:
		for _, association := range preloadAssociations {
			db = db.Preload(association).Where("name = ?", name)
		}
		return db.Model(out).First(out, "name = ?", name).Error

	default:
		for _, association := range preloadAssociations {
			db = db.Preload(association).Where("name = ?", name)
		}
		return db.Model(out).First(out, "name = ?", name).Error
	}
}

func (repos *GormRepository) Add(uow *UnitOfWork, value interface{}) error {
	return uow.DB.Create(value).Error
}

func (repos *GormRepository) Delete(uow *UnitOfWork, parentID,
	childID uuid.UUID, value interface{}) error {
	switch value.(type) {
	case *model.User:
		return uow.DB.Model(value).Delete(value, "id = ?", parentID).Error
	case *model.Category:
		return uow.DB.Model(value).Delete(value, "user_id = ? AND id = ?", parentID, childID).Error
	default:
		return uow.DB.Model(value).Delete(value, "category_id = ? AND id = ?", parentID, childID).Error
	}
}

func (repos *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Model(entity).Update(entity).Error
}

//NewGormRepository creates a new GormRepository
func NewGormRepository() *GormRepository {
	return &GormRepository{}
}
