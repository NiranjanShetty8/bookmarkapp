package security

import (
	"net/http"
	"time"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/web"
	"github.com/dgrijalva/jwt-go"
)

//30 minute session
const session int64 = 1800000

// Response is a representation of JSON response for JWT.
type Response struct {
	model.User
	Token string `json:"token"`
}

// Responds with unique token after login
func GetToken(user *model.User, w *http.ResponseWriter) {
	claims := jwt.MapClaims{
		"username":  user.Username,
		"userID":    user.ID,
		"issuedAt":  time.Now().Unix(),
		"validTill": time.Now().Unix() + session,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		web.RespondError(w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	web.RespondJSON(w, http.StatusOK, Response{Token: tokenString, User: *user})
}
