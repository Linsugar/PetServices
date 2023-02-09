package MineViews

import (
	"PetService/Middlewares"
	"PetService/Models/Mine"
	"PetService/Untils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserCollectHomeEvent 将首页事件变为自己独有的收藏
func UserCollectHomeEvent(c *gin.Context) {

}

// UserCollectSaleEvent 将分享事件变为自己独有的收藏
func UserCollectSaleEvent(c *gin.Context) {

}

// UserCollectOtherEvent 将别人列为自己关注的对象
func UserCollectOtherEvent(c *gin.Context) {

}

func Register(c *gin.Context) {
	//用户注册
	NowIp := c.ClientIP()
	value := ""
	var WeiUser Mine.WeiChat
	err2 := c.Bind(&WeiUser)
	if err2 != nil {
		Untils.ResponseBadState(c, err2)
		return
	}
	WeiUser.VisitIP = NowIp
	Err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		er := tx.Model(&Mine.WeiChat{}).Where("app_id=?", WeiUser.AppId).First(&Mine.WeiChat{}).RowsAffected
		if er == 0 {
			if err3 := tx.Model(&Mine.WeiChat{}).Create(&WeiUser).Error; err3 != nil {
				// 返回任何错误都会回滚事务
				return err3
			}
		} else {
			if err3 := tx.Model(&Mine.WeiChat{}).Update(&WeiUser).Error; err3 != nil {
				// 返回任何错误都会回滚事务
				return err3
			}
		}
		token, toker := Middlewares.GenToken(WeiUser.AppId, WeiUser.NickName)
		if toker != nil {
			return toker
		}
		if err2 := tx.Model(&Mine.WeiChat{}).Where("app_id=?", WeiUser.AppId).Update("token", token).Error; err2 != nil {
			return err2
		}
		value = token
		return nil
	})
	if Err != nil {
		Untils.ResponseBadState(c, Err)
	} else {

		Untils.ResponseOkState(c, value)
	}

}

// Update_UserInfo 更新用户信息
func Update_UserInfo(c *gin.Context) {
	user := Mine.WeiChat{}
	value := UpdateInfo{}
	err := c.Bind(&value)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	Err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		er := tx.Model(user).Where("id=?", value.Uid).Find(&user).RowsAffected
		if er != 0 {
			v := tx.Model(user).Updates(map[string]interface{}{"avatar": value.Avatar, "signature": value.Signature, "nick_name": value.Nickname}).Where("id=?", value.Uid).Error
			if v != nil {
				return v
			}
			return nil
		}
		return nil
	})
	if Err != nil {
		Untils.ResponseBadState(c, Err)
	} else {
		Untils.ResponseOkState(c, user)
	}

}

type UpdateInfo struct {
	Signature string `json:"signature"`
	Nickname  string `json:"nickname"`
	Uid       int    `json:"id"`
	Avatar    string `json:"avatar"`
}

// Get_Personal_Info 获取用户详情
func Get_Personal_Info(c *gin.Context) {
	var info Mine.WeiChat
	id := c.Query("user_id")
	var err error
	if id == "" {
		err = Untils.Db.Transaction(func(tx *gorm.DB) error {
			userID, _ := c.Get("userID")
			tx.Model(Mine.WeiChat{}).Where("app_id=?", userID).First(&info)
			return nil
		})
	} else {
		err = Untils.Db.Transaction(func(tx *gorm.DB) error {
			tx.Model(Mine.WeiChat{}).Where("id=?", id).First(&info)
			return nil
		})
	}
	if err != nil {
		Untils.ResponseBadState(c, err)
	} else {
		Untils.ResponseOkState(c, info)
	}
}

// 用户信息更新控制
func Person_Info_Controller(c *gin.Context) {
	if c.Request.Method == "POST" {
		Update_UserInfo(c)
	}
	if c.Request.Method == "GET" {
		Get_Personal_Info(c)
	}
}
