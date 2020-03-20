package controllers

import (
	"fmt"

	"github.com/NiranjanShetty8/bookmarkapp/services"
	"github.com/gorilla/mux"
)

type BookmarkController struct {
	BookmarkService *services.BookmarkService
	CategoryService *services.CategoryService
}

func (bmc *BookmarkController) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/bookmarkapp/user/{userid}").Subrouter()
	// subRouter.Use(bmc.Auth)
	fmt.Print(subRouter)
}

func NewBookmarkController(bms *services.BookmarkService,
	cs *services.CategoryService) *BookmarkController {
	return &BookmarkController{
		BookmarkService: bms,
		CategoryService: cs,
	}
}
