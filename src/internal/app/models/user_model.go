package models

import (
	hashpassword "mazano-server/src/internal/app/auth/hash_password"
)

type User struct {
	Uuid            string
	Username        string
	Password        string
	Email           string
	Favorite_Genres []string
	Avatar          string
}

func (user *User) ComparePasswords(password string) bool {
	return hashpassword.CheckPasswordHash(password, user.Password)
}
