package Routers

import (
	"PetService/Middlewares"
	"PetService/Views"
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once
var Gone *gin.Engine

//实现单例只创建一次
func engine() *gin.Engine {
	once.Do(func() {
		Gone = gin.Default()
	})
	return Gone
}

func Router() {
	R := engine()
	Gone.Use(Middlewares.JWThMiddleware())
	V1Route := R.Group("/UserCenter")
	{
		V1Route.Any("/user", Views.UserController)
		V1Route.Any("/pet", Views.PetController)
		V1Route.Any("/dynamic", Views.DynamicController)
		V1Route.Any("/article", Views.ArticleController)
		V1Route.GET("/weixin", Views.WeixinGet)
		V1Route.POST("/register", Views.Register)
		V1Route.POST("/check_login", Views.Check_login)
		V1Route.Any("/topic", Views.TopicController)
		V1Route.GET("/Info", Views.Person_Info_Controller)
		V1Route.Any("/list", Views.TalkListController)
		V1Route.GET("/newtype", Views.GetNewTopic)
	}
	V2Route := R.Group("/UserConfig")
	{
		V2Route.Any("/QiNiu", Views.SetQINiuToken)     //获取七牛云token
		V2Route.Any("/CodeWith", Views.CodeController) //获取验证码
	}

}
