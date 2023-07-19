package dao

import (
	"bytetrack_server/repository/dal"
	"bytetrack_server/repository/model"

	"golang.org/x/crypto/bcrypt"
)

const (
	PassWordCost = 12 //密码加密难度
)

// SetPassword 设置密码
func SetPassword(user *model.User, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func CheckPassword(user *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// CreateUser 创建User
func CreateUser(user *model.User) (err error) {
	err = dal.User.Create(user)
	return err
}

// FindUserByAccount 根据用户名找到用户
func FindUserByAccount(account string) (user *model.User, err error) {
	acc, err := dal.User.Where(dal.User.Account.Eq(account)).First()
	return acc, err
}
