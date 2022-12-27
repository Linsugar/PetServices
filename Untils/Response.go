package Untils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseOkState[T any](c *gin.Context, Data T) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "请求成功",
		"data":    Data,
		"status":  http.StatusOK,
	})
}

func ResponseBadState(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"Message": "请求有误,请检查",
		"data":    err.Error(),
		"status":  http.StatusInternalServerError,
	})
}
