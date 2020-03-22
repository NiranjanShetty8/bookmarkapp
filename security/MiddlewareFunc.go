package security

import (
	"fmt"
	"net/http"

	"github.com/NiranjanShetty8/bookmarkapp/web"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var secretKey = []byte("Private_Key")

// Middleware func to check token and provide or deny access to resources
func AuthMiddleWareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := request.HeaderExtractor{"token"}.ExtractToken(r)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})
		if err != nil {
			web.RespondError(&w, web.NewHTTPError("Access Denied.", http.StatusForbidden))
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var actualUserID, userID uuid.UUID
			var err error
			actualUserID, err = uuid.FromString(mux.Vars(r)["userid"])
			if err != nil {
				web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}
			id, ok := claims["userID"].(string)
			if !ok {
				web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}
			userID, err = uuid.FromString(id)
			if err != nil {
				web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}

			if userID != actualUserID {
				web.RespondError(&w, web.NewHTTPError("Access Denied", http.StatusForbidden))
				return
			}
			next.ServeHTTP(w, r)
		} else {
			web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
		}
	})
}

// To get the secret key
func GetSecretKey() []byte {
	return secretKey
}
