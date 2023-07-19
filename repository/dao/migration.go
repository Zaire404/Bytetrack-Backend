package dao

import (
	"bytetrack_server/repository/model"
)

func migration() {
	err := _db.AutoMigrate(&model.User{}, &model.DeviceInfo{}, &model.DeviceBelong{})
	if err != nil {
		panic(err)
	}
}
