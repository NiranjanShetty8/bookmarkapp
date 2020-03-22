package services

import (
	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/repository"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type BookmarkService struct {
	DB         *gorm.DB
	Repository *repository.GormRepository
}

func (bms *BookmarkService) GetAllBookmarks(uid uuid.UUID, bookmarks *[]model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, true)
	err := bms.Repository.GetAll(uow, uid, bookmarks, []string{})
	return err

}

func (bms *BookmarkService) GetBookmarkById(categoryID, bookmarkID uuid.UUID, bookmark *model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, true)
	err := bms.Repository.Get(uow, categoryID, bookmarkID, bookmark, []string{})
	return err
}

func (bms *BookmarkService) AddBookmark(bookmark *model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, false)
	bookmark.ID = uuid.NewV4()
	err := bms.Repository.Add(uow, bookmark)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

func (bms *BookmarkService) DeleteBookmark(categoryID, bookmarkId uuid.UUID) error {
	uow := repository.NewUnitOfWork(bms.DB, false)
	err := bms.Repository.Delete(uow, categoryID, bookmarkId, model.Bookmark{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

func (bms *BookmarkService) UpdateBookmark(bookmark *model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, false)
	err := bms.Repository.Update(uow, bookmark)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

func (bms *BookmarkService) GetBookmarkByName(bookmarkName string, categoryID uuid.UUID,
	bookmark *model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, true)
	err := bms.Repository.GetByName(uow, bookmarkName, categoryID, bookmark, []string{})
	return err
}

func NewBookmarkService(db *gorm.DB, repos *repository.GormRepository) *BookmarkService {
	db = db.AutoMigrate(&model.Category{}, &model.Bookmark{})
	db.Model(&model.Bookmark{}).
		AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
		// AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	return &BookmarkService{
		DB:         db,
		Repository: repos,
	}
}
