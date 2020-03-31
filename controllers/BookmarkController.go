package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/security"
	"github.com/NiranjanShetty8/bookmarkapp/services"
	"github.com/NiranjanShetty8/bookmarkapp/web"
	"github.com/gorilla/mux"
)

type BookmarkController struct {
	BookmarkService *services.BookmarkService
}

func (bmkController *BookmarkController) RegisterRoutes(router *mux.Router) {
	router.Use(security.AuthMiddleWareFunc)
	router.HandleFunc("/api/bookmarkapp/user/{userid}/bookmark/all",
		bmkController.GetAllBookmarksOfUser).Methods("GET")
	subRouter := router.PathPrefix("/api/bookmarkapp/user/{userid}/category/{categoryid}").Subrouter()
	subRouter.Use(security.AuthMiddleWareFunc)
	subRouter.HandleFunc("/bookmark", bmkController.GetAllBookmarks).Methods("GET")
	subRouter.HandleFunc("/bookmark/{bookmarkid}", bmkController.GetBookmarkByID).Methods("GET")
	subRouter.HandleFunc("/bookmark/name/{bookmarkname}", bmkController.GetBookmarkByName).Methods("GET")
	subRouter.HandleFunc("/bookmark", bmkController.AddBookmark).Methods("POST")
	subRouter.HandleFunc("/bookmark/{bookmarkid}", bmkController.UpdateBookmark).Methods("PUT")
	subRouter.HandleFunc("/bookmark/{bookmarkid}", bmkController.DeleteBookmark).Methods("DELETE")
}

func (bmkController *BookmarkController) GetAllBookmarks(w http.ResponseWriter, r *http.Request) {
	cid, err := ParseID(&w, r, "categoryid")
	if err != nil {
		return
	}
	bookmarks := []model.Bookmark{}
	err = bmkController.BookmarkService.GetAllBookmarks(cid, &bookmarks)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, bookmarks)
}

func (bmkController *BookmarkController) GetAllBookmarksOfUser(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseID(&w, r, "userid")
	if err != nil {
		return
	}
	tempArray := []model.Bookmark{}
	bookmarks := []model.Bookmark{}
	categories := []model.Category{}
	csv := services.NewCategoryService(bmkController.BookmarkService.DB,
		bmkController.BookmarkService.Repository)
	err = csv.GetAllCategories(userID, &categories)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	for _, category := range categories {
		err = bmkController.BookmarkService.GetAllBookmarks(category.ID, &tempArray)
		bookmarks = append(bookmarks, tempArray...)
	}
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, bookmarks)
}

func (bmkController *BookmarkController) GetBookmarkByID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := ParseID(&w, r, "categoryid")
	if err != nil {
		return
	}
	bookmark := model.Bookmark{}
	bookmarkID, err := ParseID(&w, r, "bookmarkid")
	if err != nil {
		return
	}
	err = bmkController.BookmarkService.GetBookmarkById(categoryID, bookmarkID, &bookmark)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, bookmark)

}

func (bmkController *BookmarkController) GetBookmarkByName(w http.ResponseWriter, r *http.Request) {
	categoryID, err := ParseID(&w, r, "categoryid")
	if err != nil {
		return
	}
	bookmark := model.Bookmark{}
	err = bmkController.BookmarkService.GetBookmarkByName(mux.Vars(r)["bookmarkname"], categoryID,
		&bookmark)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, bookmark)

}

func (bmkController *BookmarkController) AddBookmark(w http.ResponseWriter, r *http.Request) {
	categoryID, err := ParseID(&w, r, "categoryid")
	if err != nil {
		return
	}
	bookmark := model.Bookmark{}
	err = web.UnmarshalJSON(r, &bookmark)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	err = validateBookmark(&bookmark)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	bookmark.CategoryID = categoryID
	err = bmkController.BookmarkService.AddBookmark(&bookmark)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, bookmark.ID)

}

func (bmkController *BookmarkController) UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	categoryID, err := ParseID(&w, r, "categoryid")
	if err != nil {
		return
	}
	bookmark := model.Bookmark{}
	err = web.UnmarshalJSON(r, &bookmark)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	err = validateBookmark(&bookmark)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	bookmarkID, err := ParseID(&w, r, "bookmarkid")
	if err != nil {
		return
	}
	bookmark.ID = bookmarkID
	bookmark.CategoryID = categoryID
	err = bmkController.BookmarkService.UpdateBookmark(&bookmark)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, "Bookmark Updated")

}

func (bmkController *BookmarkController) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	categoryID, err := ParseID(&w, r, "categoryid")
	if err != nil {
		return
	}
	bookmarkID, err := ParseID(&w, r, "bookmarkid")
	if err != nil {
		return
	}
	err = bmkController.BookmarkService.DeleteBookmark(categoryID, bookmarkID)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, "Bookmark Deleted")

}

func NewBookmarkController(bms *services.BookmarkService) *BookmarkController {
	return &BookmarkController{
		BookmarkService: bms,
	}
}

func validateBookmark(bookmark *model.Bookmark) error {
	name := strings.TrimSpace(bookmark.Name)
	if name == "" {
		return errors.New("Name is required")
	}
	url := strings.TrimSpace(bookmark.URL)
	if url == "" {
		return errors.New("URL is required")
	}
	// _,err := http.Get(bookmark.URL)
	// if err != nil {
	// 	return errors.New("URL is invalid")
	// }
	return nil
}
