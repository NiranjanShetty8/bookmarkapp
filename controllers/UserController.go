package controllers

import (
	"net/http"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/services"
	"github.com/NiranjanShetty8/bookmarkapp/web"
	"github.com/gorilla/mux"
)

type UserController struct {
	userService *services.UserService
}

func (uc *UserController) register(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Can't parse",
			map[string]string{"error": err.Error()}))
		return
	}
	if len(user.Username) == 0 {
		web.RespondError(&w,
			web.NewValidationError("fatal", map[string]string{"error": "Username is required"}))
		return
	}
	if len(user.Password) == 0 {
		web.RespondError(&w,
			web.NewValidationError("fatal", map[string]string{"error": "Password is required"}))
		return
	}
	if len(user.Password) < 8 {
		web.RespondError(&w, web.NewValidationError("required",
			map[string]string{"error": "Password should be more than 8 characters"}))
		return
	}
	err = uc.userService.Register(&user)

	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, user.ID)

}

func (uc *UserController) login(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Can't parse",
			map[string]string{"error": err.Error()}))
		return
	}
	if len(user.Username) == 0 {
		web.RespondError(&w,
			web.NewValidationError("fatal", map[string]string{"error": "Username is required"}))
		return
	}
	if len(user.Password) == 0 {
		web.RespondError(&w,
			web.NewValidationError("fatal", map[string]string{"error": "Password is required"}))
		return
	}
	if len(user.Password) < 8 {
		web.RespondError(&w, web.NewValidationError("required",
			map[string]string{"error": "Password should be more than 8 characters"}))
		return
	}
	err = uc.userService.Login(&user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("mismatch", map[string]string{"error": err.Error()}))
		return
	}
}

func (uc *UserController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("api/bookmarkapp/user/register", uc.register).Methods("POST")
	r.HandleFunc("api/bookmarkapp/user/login", uc.login).Methods("POST")
}
