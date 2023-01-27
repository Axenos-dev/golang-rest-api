package models

import (
	hashpassword "mazano-server/src/internal/app/auth/hash_password"

	"github.com/google/uuid"
)

type NewUserRequestData struct {
	Username string
	Email    string
	Password string
}

type NewUser struct {
	Uuid            string
	Username        string
	Password        string
	Email           string
	Avatar          string
	Favorite_Genres []string
}

type NewUserResponse struct {
	Code        int16
	Description string
	Uuid        string
}

func (user *NewUser) CreateUser(user_req NewUserRequestData) {
	pass_hash, _ := hashpassword.HashPassword(user_req.Password)

	user.Uuid = uuid.New().String()
	user.Avatar = ""
	user.Email = user_req.Email
	user.Favorite_Genres = []string{}
	user.Password = pass_hash
	user.Username = user_req.Username
}

func (response *NewUserResponse) CreateResponse(code int16, description string, data NewUser) {
	response.Code = code
	response.Description = description
	response.Uuid = data.Uuid
}
