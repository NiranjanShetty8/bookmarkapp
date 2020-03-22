package services

import (
	"fmt"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/repository"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CategoryService struct {
	DB         *gorm.DB
	Repository *repository.GormRepository
}

func (cs *CategoryService) GetAllCategories(uid uuid.UUID, categories *[]model.Category) error {
	uow := repository.NewUnitOfWork(cs.DB, true)
	err := cs.Repository.GetAll(uow, uid, categories, []string{"Bookmarks"})
	return err
}

func (cs *CategoryService) GetCategory(userID, categoryID uuid.UUID, category *model.Category) error {
	uow := repository.NewUnitOfWork(cs.DB, true)
	fmt.Println("uid:", userID, categoryID)
	err := cs.Repository.Get(uow, userID, categoryID, category, []string{"Bookmarks"})
	fmt.Print("error", err)
	return err
}

func (cs *CategoryService) GetCategoryByName(categoryName string, userID uuid.UUID,
	category *model.Category) error {
	uow := repository.NewUnitOfWork(cs.DB, true)
	err := cs.Repository.GetByName(uow, categoryName, userID, category, []string{"Bookmarks"})
	return err
}

func (cs *CategoryService) AddCategory(category *model.Category) error {
	uow := repository.NewUnitOfWork(cs.DB, false)
	category.ID = uuid.NewV4()
	err := cs.Repository.Add(uow, category)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

func (cs *CategoryService) DeleteCategory(userId, categoryId uuid.UUID) error {
	uow := repository.NewUnitOfWork(cs.DB, false)
	err := cs.Repository.Delete(uow, userId, categoryId, &model.Category{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

func (cs *CategoryService) UpdateCategory(category *model.Category) error {
	uow := repository.NewUnitOfWork(cs.DB, false)
	err := cs.Repository.Update(uow, category)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

func NewCategoryService(db *gorm.DB, repos *repository.GormRepository) *CategoryService {
	db.AutoMigrate(&model.User{}, &model.Category{})
	db.Model(&model.Category{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	return &CategoryService{
		DB:         db,
		Repository: repos,
	}
}
