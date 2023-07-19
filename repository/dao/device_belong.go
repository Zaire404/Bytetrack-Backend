package dao

import (
	"bytetrack_server/repository/dal"
	"bytetrack_server/repository/model"
	"bytetrack_server/util/applog"
)

func CreateDeviceBelong(deviceID int32, userID int32) (err error) {
	data := model.DeviceBelong{UserID: userID, DeviceID: deviceID}
	err = dal.DeviceBelong.Create(&data)
	return err
}

func GetDeviceIDByUserID(userID int32) (device_id_list []int32, err error) {
	data_list, err := dal.DeviceBelong.Where(dal.DeviceBelong.UserID.Eq(userID)).Find()
	if err != nil {
		applog.Logger("device").Error().Err(err).Send()
		return nil, err
	}
	for _, data := range data_list {
		device_id_list = append(device_id_list, data.DeviceID)
	}
	return device_id_list, err
}
