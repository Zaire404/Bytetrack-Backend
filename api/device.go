package api

import (
	"bytetrack_server/repository/model"
	"bytetrack_server/service"
	"bytetrack_server/util/applog"
	"bytetrack_server/util/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeviceCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type Request struct {
			model.DeviceInfo
			UserID int32 `json:"userID"`
		}
		var request Request
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, response.RespError(err, "userID格式错误"))
			return
		}
		resp, err := service.CreateDevice(&request.DeviceInfo, request.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.RespError(err, "添加失败"))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func DeviceInfoGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := strconv.Atoi(ctx.Param("userID"))
		if err != nil {
			applog.Logger("device").Info().Str("userID", ctx.Param("userID")).Send()
			applog.Logger("device").Error().Err(err).Send()
			return
		}
		resp, err := service.GetDeviceInfo(int32(userID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.RespError(err, "查询失败"))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
