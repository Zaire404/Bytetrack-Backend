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
	r.GET("device/get_device/:userID", api.DeviceInfoGet())
	return r
}
