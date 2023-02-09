package Routers

import (
	"PetService/Middlewares"
	"PetService/Views/CommentViews"
	"PetService/Views/FaceViews"
	"PetService/Views/HomeViews"
	"PetService/Views/MineViews"
	"PetService/Views/RunViews"
	"PetService/Views/SaleViews"
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
		//V1Route.Any("/user", Views.UserController)
		//V1Route.Any("/pet", Views.PetController)
		//V1Route.Any("/dynamic", Views.DynamicController)
		//V1Route.Any("/article", Views.ArticleController)
		//V1Route.GET("/weixin", Views.WeixinGet)
		V1Route.POST("/register", MineViews.Register)
		V1Route.POST("/check_login", CommentViews.Check_login)
		V1Route.Any("/topic", HomeViews.TopicController)
		V1Route.Any("/Info", MineViews.Person_Info_Controller)
		V1Route.Any("/list", HomeViews.TalkListController)
		V1Route.GET("/new_messages", HomeViews.GetNewTopic)
		V1Route.Any("/sale", SaleViews.FriendController)
		V1Route.Any("/comment", SaleViews.CommentController)
		V1Route.Any("/detail", SaleViews.FriendDetail)
		V1Route.GET("/most_new_sale_friend", SaleViews.GetNewFriends)
		V1Route.POST("/TencentFace", FaceViews.Tencent)
		V1Route.POST("/TencentAge", FaceViews.TencentAge)
		V1Route.POST("/run_data", RunViews.GetWeiXinRunData)
		V1Route.POST("/follow", HomeViews.CollectPostList)
		V1Route.GET("/run_steps", RunViews.GetListRunData)
		V1Route.GET("/run_statistic", RunViews.GetUserRunData)
		V1Route.GET("/session", RunViews.GetSession)
	}
	V2Route := R.Group("/UserConfig")
	{
		V2Route.Any("/QiNiu", CommentViews.SetQINiuToken) //获取七牛云token
	}

}
