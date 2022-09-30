package service

import (
	"golang.org/x/crypto/bcrypt"
)

// 加密
func HashAndSalt(pwdStr string) (pwdHash string, err error) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, 10)
	if err != nil {
		return
	}
	pwdHash = string(hash)
	return
}

// 驗證
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	return err == nil

}
