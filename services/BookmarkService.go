package service

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

func NewBookmarkService(db *gorm.DB, repos *repository.GormRepository) *BookmarkService {
	return &BookmarkService{
		DB:         db,
		Repository: repos,
	}
}

func (bms *BookmarkService) GetAllBookmarks(uid uuid.UUID, bookmarks *[]model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, true)
	err := bms.Repository.GetAll(uow, uid, bookmarks, []string{})
	return err

}

func (bms *BookmarkService) GetBookmarkById(userId, bookmarkId uuid.UUID, bookmark *model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, true)
	err := bms.Repository.Get(uow, userId, bookmarkId, bookmark, []string{})
	return err
}

func (bms *BookmarkService) AddBookmark(bookmark *model.Bookmark) error {
	uow := repository.NewUnitOfWork(bms.DB, false)
	err := bms.Repository.Add(uow, bookmark)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

func (bms *BookmarkService) DeleteBookmark(userId, bookmarkId uuid.UUID) error {
	uow := repository.NewUnitOfWork(bms.DB, false)
	err := bms.Repository.Delete(uow, userId, bookmarkId, model.Bookmark{})
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

//uuid.Nil try to make to category name
func (bms *BookmarkService) GetBookmarksByCategory(userId, categoryId uuid.UUID, bookmark *model.Bookmark) {
	uow := repository.NewUnitOfWork(bms.DB, true)
	err := bms.Repository.GetAllByCategory(uow, categoryId, userId, bookmark, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}
