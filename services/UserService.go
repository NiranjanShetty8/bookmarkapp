package services

import (
	"errors"
	"strconv"
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
	if actualUser.LoginAttempts == 0 {
		return errors.New("Login attemps over. Account locked")
	}
	if user.Password != actualUser.Password {
		actualUser.LoginAttempts--
		us.Repository.Save(uow, actualUser)
		if actualUser.LoginAttempts == 0 {
			return errors.New("Login attemps over. Account locked")
		}
		attempts := strconv.Itoa(actualUser.LoginAttempts)
		return errors.New("Incorrect Password. Attempts remaining: " + attempts)
	}
	actualUser.LoginAttempts = model.GetLoginAttempts()
	us.Repository.Save(uow, actualUser)
	return err
}

func (us *UserService) UpdateUser(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	err := us.Repository.Update(uow, user)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
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
