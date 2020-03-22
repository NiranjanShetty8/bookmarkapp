package services

import (
	"errors"
	"strings"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/repository"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	DB         *gorm.DB
	Repository *repository.GormRepository
}

//Service used to register new user
func (us *UserService) Register(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	user.ID = uuid.NewV4()
	err := us.Repository.Add(uow, user)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

//Service used to login an existing user
func (us *UserService) Login(user, actualUser *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, true)
	err := us.Repository.GetByName(uow, user.Username, uuid.Nil, actualUser,
		[]string{"Categories", "Categories.Bookmarks"})
	if err != nil {
		if strings.EqualFold(err.Error(), "Record not found") {
			return errors.New("Incorrect Username")
		}
		return err
	}
	if user.Password != actualUser.Password {
		return errors.New("Incorrect Password")
	}
	return err
}

//Returns instance of UserService
func NewUserService(db *gorm.DB, repos *repository.GormRepository) *UserService {
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Bookmark{})
	return &UserService{
		DB:         db,
		Repository: repos,
	}
}
