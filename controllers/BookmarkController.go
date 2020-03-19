package controllers

import (
	"fmt"
	"net/http"

	"github.com/NiranjanShetty8/bookmarkapp/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
)

type BookmarkController struct {
	BookmarkService *services.BookmarkService
	CategoryService *services.CategoryService
}

func (bmc *BookmarkController) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/bookmarkapp/user/{userid}").Subrouter()
	// subRouter.Use(bmc.Auth)
}

func (bmc *BookmarkController) AuthMiddleWareFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := request.HeaderExtractor{"token"}.ExtractToken(r)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return secretKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Access Denied; Please Login First."))
			return
		}
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			// web.HeaderWrite(&w, http.StatusForbidden, err)
		}
	})
}

func NewBookmarkController(bms *service.BookmarkService,
	cs *service.CategoryService) *BookmarkController {
	return &BookmarkController{
		BookmarkService: bms,
		CategoryService: cs,
	}
}
