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

type UserController struct {
	userService *services.UserService
}

// Returns instance of UserController
func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

// Does validations and adds user to database
func (uc *UserController) register(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Can't parse",
			map[string]string{"error": err.Error()}))
		return
	}
	err = uc.validateUser(w, &user)
	if err != nil {
		web.RespondError(&w,
			web.NewValidationError("Invalid", map[string]string{"error": err.Error()}))
		return
	}
	err = uc.userService.Register(&user)

	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, user.ID)

}

// Handles validation,Authentication & Authorization of user
func (uc *UserController) login(w http.ResponseWriter, r *http.Request) {

	user := model.User{}
	actualUser := model.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Can't parse",
			map[string]string{"error": err.Error()}))
		return
	}
	err = uc.validateUser(w, &user)
	if err != nil {
		web.RespondError(&w,
			web.NewValidationError("Invalid", map[string]string{"error": err.Error()}))
		return
	}

	err = uc.userService.Login(&user, &actualUser)
	if actualUser.LoginAttempts == 1 {
		web.RespondErrorMessage(&w, http.StatusForbidden, err.Error()+
			" Please send an e-mail to niranjan@swabhavtechlabs.com for account unlock.")
		return
	}
	if err != nil {
		web.RespondError(&w, web.NewValidationError("mismatch",
			map[string]string{"error": err.Error()}))
		return
	}
	security.GetToken(&actualUser, &w)
}

// Does the actual validation
func (uc *UserController) validateUser(w http.ResponseWriter, user *model.User) error {
	username := strings.TrimSpace(user.Username)
	if username == "" {
		return errors.New("Username is required.")
	}

	if len(user.Password) == 0 {
		return errors.New("Password is required.")
	}
	if len(user.Password) < 8 {
		return errors.New("Password should be 8 or more than 8 characters.")
	}
	return nil
}

// Registers routes in router
func (uc *UserController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/bookmarkapp/user/register", uc.register).Methods("POST")
	r.HandleFunc("/api/bookmarkapp/user/login", uc.login).Methods("POST")
}
