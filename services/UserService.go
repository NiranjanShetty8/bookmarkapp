package services

import (
	"fmt"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/repository"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	DB         *gorm.DB
	Repository *repository.GormRepository
}

func (us *UserService) Register(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	user.ID = uuid.NewV4()
	err := us.Repository.Add(uow, user)
	if err != nil {
		uow.Complete()
		return err
	}
	fmt.Print("register ID: ", user.ID)
	fmt.Printf("%T", user.ID)
	uow.Commit()
	return err
}

func (us *UserService) Login(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, true)
	actualUser := model.User{}
	err := us.Repository.Get(uow, user.ID, uuid.Nil, &actualUser, []string{"Bookmarks"})
	if err != nil {
		return err
	}
	if actualUser.Username == "" {
		return fmt.Errorf("Incorrect UserName")
	}
	if user.Password != actualUser.Password {
		return fmt.Errorf("Incorrect Password")
	}
	fmt.Println("Access Granted")
	return err
}

func NewUserService(db *gorm.DB, repos *repository.GormRepository) *UserService {
	db.AutoMigrate(model.User{})
	return &UserService{
		DB:         db,
		Repository: repos,
	}
}
