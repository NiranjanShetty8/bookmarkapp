package controllers

import (
	"net/http"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/services"
	"github.com/NiranjanShetty8/bookmarkapp/web"
	"github.com/gorilla/mux"
)

type SuperUserController struct {
	superUserService *services.SuperUserService
}

// Registers routes of super user
func (suc *SuperUserController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/bookmarkapp/user/all", suc.getAllUsers).Methods("GET")
	r.HandleFunc("/api/bookmarkapp/user/{userid}", suc.deleteUser).Methods("DELETE")
}

// Gets all users in db
func (suc *SuperUserController) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	err := suc.superUserService.GetAllUsers(&users)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	web.RespondJSON(&w, http.StatusOK, users)
}

// Deletes a speific user
func (suc *SuperUserController) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseID(&w, r, "userid")
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	err = suc.superUserService.DeleteUser(userID)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	web.RespondJSON(&w, http.StatusOK, "User Deleted")
}

// Returns instance of SuperUserController
func NewSuperUserController(sus *services.SuperUserService) *SuperUserController {
	return &SuperUserController{
		superUserService: sus,
	}
}
