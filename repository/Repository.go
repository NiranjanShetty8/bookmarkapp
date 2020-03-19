package repository

import (
	"fmt"

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
	GetAllByCategory(ufw *UnitOfWork, value interface{}, condition string,
		preloadAssociations []string) error
	Get(uow *UnitOfWork, userId, bookmarkID uuid.UUID, out interface{},
		preloadAssociations []string) error
	Add(uow *UnitOfWork, input interface{}) error
	Delete(uow *UnitOfWork, userId, bookmarkID uuid.UUID, out interface{}) error
	Update(uow *UnitOfWork, uid uuid.UUID, entity interface{}) error
}

//GormRepository implements Repository
type GormRepository struct{}

func (repos *GormRepository) GetAll(uow *UnitOfWork, uid uuid.UUID, out interface{},
	preloadAssociations []string) error {
	//change preload
	db := uow.DB
	for _, association := range preloadAssociations {
		db = uow.DB.Preload(association)
	}
	return db.Model(out).Find(out, "user_id = ?", uid).Error
}

func (repos *GormRepository) Get(uow *UnitOfWork, userId, bookmarkID uuid.UUID, out interface{},
	preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association).Where("id = ?", userId)
	}
	if bookmarkID == uuid.Nil {
		return db.Model(out).First(out, "id = ?", userId).Error
	}
	return db.Model(out).First(out, "user_id = ? AND id = ?", userId, bookmarkID).Error
}

func (repos *GormRepository) Add(uow *UnitOfWork, value interface{}) error {
	fmt.Print(value)
	return uow.DB.Create(value).Error
}

func (repos *GormRepository) Delete(uow *UnitOfWork, userId,
	bookmarkID uuid.UUID, value interface{}) error {
	return uow.DB.Model(value).Delete(value, "user_id = ? AND id = ?", userId, bookmarkID).Error
}

func (repos *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Model(entity).Update(entity).Error
}

func (repo *GormRepository) GetAllByCategory(uow *UnitOfWork, value interface{},
	uid interface{}, out interface{}, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	if uid == nil {
		return db.Model(out).First(out, "category_id  = ?", value).Error
	}
	return db.Model(out).First(out, "category_id = ? AND user_id = ?", value, uid).Error
}

//NewGormRepository creates a new GormRepository
func NewGormRepository() *GormRepository {
	return &GormRepository{}
}
