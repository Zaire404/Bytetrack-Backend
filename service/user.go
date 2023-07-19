package service

import (
	"bytetrack_server/repository/dao"
	"bytetrack_server/repository/model"
	"bytetrack_server/util/applog"
	"bytetrack_server/util/response"
	"errors"

	"gorm.io/gorm"
)

func Register(user *model.User) (resp interface{}, err error) {
	_, err = dao.FindUserByAccount(user.Account)

	switch err {
	case gorm.ErrRecordNotFound:
		if err = dao.SetPassword(user, user.Password); err != nil {
			applog.Logger("user").Error().Err(err).Send()
			return nil, err
		}

		if err = dao.CreateUser(user); err != nil {
			applog.Logger("user").Error().Err(err).Send()
			return nil, err
		}

		return response.RespSuccess("Register Successfully!"), nil

	case nil:
		err = errors.New("用户已存在")
		return nil, err
	default:
		return nil, nil
	}
}

func Login(user *model.User) (resp interface{}, err error) {
	acc, err := dao.FindUserByAccount(user.Account)
	if err != nil {
		err = errors.New("用户不存在")
		return nil, err
	}
	if dao.CheckPassword(acc, user.Password) {
		return response.RespSuccess("Login Successfully"), nil
	} else {
		err = errors.New("账号/密码错误")
		return nil, err
	}

}