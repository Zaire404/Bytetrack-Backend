package service

import (
	"bytetrack_server/repository/dao"
	"bytetrack_server/repository/model"
	"bytetrack_server/util/applog"
	"bytetrack_server/util/response"
)

func CreateDevice(device *model.DeviceInfo, userID int32) (resp interface{}, err error) {

	if err = dao.CreateDeviceInfo(device); err != nil {
		applog.Logger("device").Error().Err(err).Send()
		return nil, err
	}

	if err = dao.CreateDeviceBelong(device.ID, userID); err != nil {
		applog.Logger("device").Error().Err(err).Send()
		return nil, err
	}

	return response.RespSuccess("Create Successfully!"), nil
}

func GetDeviceInfo(userID int32) (resp interface{}, err error) {
	device_id_list, err := dao.GetDeviceIDByUserID(userID)
	if err != nil {
		return nil, err
	}

	device_info_list, err := dao.GetDeviceInfoByIDList(device_id_list)
	if err != nil {
		return nil, err
	}

	return response.RespSuccessWithData(device_info_list), nil

}
