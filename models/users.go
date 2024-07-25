package models

import (
	"crypto/md5"
	"encoding/hex"
)

type LOGINUSER struct {
	Email    string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// hashPassword calculates the MD5 hash of a given password
func HashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Login(input LOGINUSER) (Users, error) {
	input.Password = HashPassword(input.Password)
	var user Users
	err := DB.Model(&Users{}).Where(&Users{Email: input.Email, Password: input.Password}).First(&user).Error
	return user, err
}
