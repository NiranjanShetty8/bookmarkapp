package security

import "github.com/NiranjanShetty8/bookmarkapp/model"

// Response is a representation of JSON response for JWT
type Response struct {
	model.User
	Token string `json:"token"`
}
