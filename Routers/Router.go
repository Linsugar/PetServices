package Routers

import (
	"PetService/Middlewares"
	"PetService/Views"
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once
var Gone *gin.Engine

// 实现单例只创建一次
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
		V1Route.Any("/Info", Views.Person_Info_Controller)
		V1Route.Any("/list", Views.TalkListController)
		V1Route.GET("/new_messages", Views.GetNewTopic)
		V1Route.Any("/sale", Views.FriendController)
		V1Route.Any("/comment", Views.CommentController)
		V1Route.Any("/detail", Views.FriendDetail)
		V1Route.GET("/most_new_sale_friend", Views.GetNewFriends)
		V1Route.POST("/TencentFace", Views.Tencent)
		V1Route.POST("/TencentAge", Views.TencentAge)
		V1Route.POST("/run_data", Views.GetWeiXinRunData)
		V1Route.GET("/run_steps", Views.GetListRunData)
		V1Route.GET("/run_statistic", Views.GetUserRunData)
		V1Route.GET("/session", Views.GetSession)
	}
	V2Route := R.Group("/UserConfig")
	{
		V2Route.Any("/QiNiu", Views.SetQINiuToken)     //获取七牛云token
		V2Route.Any("/CodeWith", Views.CodeController) //获取验证码
	}

}
