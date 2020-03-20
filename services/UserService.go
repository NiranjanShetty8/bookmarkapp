package services

import (
	"fmt"
	"strings"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/repository"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	DB         *gorm.DB
	Repository *repository.GormRepository
}

func (us *UserService) Register(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	// user.ID = uuid.NewV4()
	err := us.Repository.Add(uow, user)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

func (us *UserService) Login(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, true)
	actualUser := model.User{}
	err := us.Repository.GetByName(uow, user.Username, &actualUser,
		[]string{"Categories", "Categories.Bookmarks"})
	if err != nil {
		if strings.EqualFold(fmt.Sprintf("%v", err), "Record not found") {
			return fmt.Errorf("Incorrect Username")
		}
		return err
	}
	if user.Password != actualUser.Password {
		return fmt.Errorf("Incorrect Password")
	}
	fmt.Println("Access Granted", actualUser.Username,
		actualUser.Categories)
	return err
}

func NewUserService(db *gorm.DB, repos *repository.GormRepository) *UserService {
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Bookmark{})
	return &UserService{
		DB:         db,
		Repository: repos,
	}
}
