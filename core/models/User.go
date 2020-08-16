package models

import (
	"bytes"
	"crypto/sha512"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:unique_index";json:"username"`
	Email    string `gorm:unique_index;json:"email"`
	Salt     []byte
	Password []byte
	LastDate time.Time `json:"last_date"`
}

func (user *User) ValidPassword(password string) bool {
	dk := pbkdf2.Key([]byte(password), user.Salt, 100000, 128, sha512.New)
	return bytes.Equal(dk, user.Password)
}