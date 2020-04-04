package services

import (
	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/repository"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SuperUserService struct {
	DB          *gorm.DB
	Respository *repository.GormRepository
}

func NewSuperUserService(db *gorm.DB, repository *repository.GormRepository) *SuperUserService {
	return &SuperUserService{
		DB:          db,
		Respository: repository,
	}
}

func (sus *SuperUserService) GetAllUsers(users *[]model.User) error {
	uow := repository.NewUnitOfWork(sus.DB, true)
	err := sus.Respository.GetAll(uow, uuid.Nil, users, []string{"Categories", "Categories.Bookmarks"})
	return err
}

func (sus *SuperUserService) DeleteUser(userID uuid.UUID) error {
	uow := repository.NewUnitOfWork(sus.DB, false)
	user := model.User{}
	err := sus.Respository.Get(uow, uuid.Nil, userID, &user, []string{"Categories", "Categories.Bookmarks"})
	if err != nil {
		uow.Complete()
		return err
	}
	for _, category := range user.Categories {
		err = sus.Respository.Delete(uow, user.ID, category.ID, &model.Category{})
		if err != nil {
			uow.Complete()
			return err
		}
		for _, bookmark := range category.Bookmarks {
			err = sus.Respository.Delete(uow, category.ID, bookmark.ID, model.Bookmark{})
			if err != nil {
				uow.Complete()
				return err
			}
		}
	}
	err = sus.Respository.Delete(uow, uuid.Nil, userID, &model.User{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}
