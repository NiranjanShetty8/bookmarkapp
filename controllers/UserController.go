package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/security"
	"github.com/NiranjanShetty8/bookmarkapp/services"
	"github.com/NiranjanShetty8/bookmarkapp/web"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

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
	fmt.Println("login", actualUser)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("mismatch", map[string]string{"error": err.Error()}))
		return
	}
	uc.getToken(&actualUser, &w)
}

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

func (uc *UserController) getToken(user *model.User, w *http.ResponseWriter) {
	const session int64 = 600
	fmt.Println("token", user)
	claims := jwt.MapClaims{
		"username": user.Username,
		"userID":   user.ID,
		"IssuedAt": time.Now().Unix() + session,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	fmt.Println(token)
	fmt.Println(security.GetSecretKey())
	tokenString, err := token.SignedString(security.GetSecretKey())
	if err != nil {
		web.RespondError(w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	web.RespondJSON(w, http.StatusOK, security.Response{Token: tokenString, User: *user})
}

func (uc *UserController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/bookmarkapp/user/register", uc.register).Methods("POST")
	r.HandleFunc("/api/bookmarkapp/user/login", uc.login).Methods("POST")
}
