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

// Service to get all bookmarks
func (bms *BookmarkService) GetAllBookmarks(uid uuid.UUID, bookmarks *[]model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, true)
	err := bms.Repository.GetAll(uow, uid, bookmarks, []string{})
	return err

}

// Gets bookmark by ID
func (bms *BookmarkService) GetBookmarkById(categoryID, bookmarkID uuid.UUID, bookmark *model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, true)
	err := bms.Repository.Get(uow, categoryID, bookmarkID, bookmark, []string{})
	return err
}

// Adds a new bookmark
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

// Deletes specific bookmark
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

//Updates specific bookmark
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

// Gets bookmark by name
func (bms *BookmarkService) GetBookmarkByName(bookmarkName string, categoryID uuid.UUID,
	bookmark *model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, true)
	err := bms.Repository.GetByName(uow, bookmarkName, categoryID, bookmark, []string{})
	return err
}

// returns instance of BookmarkService
func NewBookmarkService(db *gorm.DB, repos *repository.GormRepository) *BookmarkService {
	db = db.AutoMigrate(&model.Category{}, &model.Bookmark{})
	db.Model(&model.Bookmark{}).
		AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	return &BookmarkService{
		DB:         db,
		Repository: repos,
	}
}
