package routes

import (
	"bytetrack_server/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.POST("user/register", api.UserRegisterHandler())
	r.POST("user/login", api.UserLoginHandler())
	r.POST("device/add", api.DeviceCreateHandler())
	r.POST("device/delete", api.DeviceDeleteHandler())
	r.POST("device/update", api.DeviceUpdateHandler())
	r.GET("device/get/:userID", api.DeviceInfoGet())

	return r
}
