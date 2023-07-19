package dao

import (
	"bytetrack_server/repository/dal"
	"bytetrack_server/repository/model"
	"bytetrack_server/util/applog"
)

func CreateDeviceInfo(device *model.DeviceInfo) (err error) {
	err = dal.DeviceInfo.Create(device)
	return err
}

func DeleteDeviceInfoByID(id int32) (err error) {
	_, err = dal.DeviceInfo.Where(dal.DeviceInfo.ID.Eq(id)).Delete()
	return err
}
func GetDeviceInfoByIDList(id_list []int32) (device_info_list []*model.DeviceInfo, err error) {
	for _, val := range id_list {
		acc, err := GetDeviceInfoByID(val)
		if err != nil {
			applog.Logger("device").Error().Err(err).Send()
			continue
		}
		device_info_list = append(device_info_list, acc)
	}
	return device_info_list, err
}

func GetDeviceInfoByID(id int32) (device_info *model.DeviceInfo, err error) {
	device_info, err = dal.DeviceInfo.Where(dal.DeviceInfo.ID.Eq(id)).First()
	return device_info, err
}
