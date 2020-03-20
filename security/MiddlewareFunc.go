package security

import (
	"fmt"
	"net/http"

	"github.com/NiranjanShetty8/bookmarkapp/web"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var secretKey = []byte("Private_Key")

func AuthMiddleWareFunc(h http.Handler) http.Handler {
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
			web.RespondError(&w, web.NewHTTPError("Access Denied.", http.StatusForbidden))
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var actualUserID, userID *uuid.UUID
			var err error
			*actualUserID, err = uuid.FromString(mux.Vars(r)["userid"])
			if err != nil {
				web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}
			//ID check
			id, ok := claims["userID"].(string)
			if !ok {
				web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}
			*userID, err = uuid.FromString(id)
			if err != nil {
				web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}

			if *userID != *actualUserID {
				web.RespondError(&w, web.NewHTTPError("Access Denied", http.StatusForbidden))
				return
			}
			h.ServeHTTP(w, r)
		} else {
			// w.WriteHeader(http.StatusForbidden)
			// w.Write([]byte(err.Error()))
			web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
		}
	})
}
