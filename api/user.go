package api

import (
	"bytetrack_server/repository/model"
	"bytetrack_server/service"
	"bytetrack_server/util/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.User
		err := ctx.ShouldBind(&data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.RespError(err, "格式错误"))
		}
		resp, err := service.Register(&data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.RespError(err, "注册失败"))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.User
		err := ctx.ShouldBind(&data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.RespError(err, "格式错误"))
		}
		resp, err := service.Login(&data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.RespError(err, "登录失败"))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
