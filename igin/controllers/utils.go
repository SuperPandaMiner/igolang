package controllers

import (
	"github.com/gin-gonic/gin"
	"igin/models"
	"net/http"
	"utils"
)

// 能获取到值 ok 为 true，否则为 false
func getBool(ctx *gin.Context, key string) (value bool, ok bool) {
	val := ctx.Query(key)
	if val == "" {
		return false, false
	}
	if val == "true" || val == "1" {
		return true, true
	} else {
		return false, true
	}
}

func getInt64(ctx *gin.Context, key string) (int64, bool) {
	if i64, err := utils.ParseInt64(ctx.Query(key), 0); err != nil {
		return 0, false
	} else {
		return i64, true
	}
}

func getUint64(ctx *gin.Context, key string) (uint64, bool) {
	if ui64, err := utils.ParseUint64(ctx.Query(key), 0); err != nil {
		return 0, false
	} else {
		return ui64, true
	}
}

func abort(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse(err.Error()))
}

func abortWithCode(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, models.ErrorResponseWithCode(code, err.Error()))
}

func writeSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, models.OkResponse(data))
}
