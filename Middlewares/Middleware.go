package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func FirstCheck(MapValue map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		IP := c.ClientIP()
		v, ok := MapValue[IP]
		if ok {
			t1 := v.(int64)
			if t2 := time.Now().Unix(); t2-t1 > 5 {
				//可以继续操作
				c.Next()
			} else {
				//
				c.Abort()
			}
		} else {
			MapValue[IP] = time.Now().Unix()
			fmt.Printf("当前ip名单：%v", MapValue)
			c.Next()
		}
	}
}

func JWThMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的token中
		fmt.Println("进入了")
		url := c.Request.URL
		method := c.Request.Method
		if url.Path == "/UserCenter/register" && method == "POST" {
			c.Next()
			return
		}
		token := c.Request.Header.Get("token")
		if token == "" {
			// 处理 没有token的时候
			c.JSON(403, gin.H{
				"err":    "丢失token",
				"status": http.StatusInternalServerError,
			})
			c.Abort() // 不会继续停止
			return
		}
		// 解析
		mc, err := ParseToken(token)
		if err != nil {
			// 处理 解析失败
			fmt.Printf("解析失败：%v\n", err.Error())
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userID", mc.UserID)
		c.Next()
	}
}
