package Views

import (
	"PetService/Middlewares"
	"PetService/Models"
	"PetService/Untils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserController(c *gin.Context) {
	if c.Request.Method == "POST" {
		UserPost(c)
	} else if c.Request.Method == "GET" {
		UserGet(c)
	}
}

func UserGet(c *gin.Context) {
	//获取所有的用户列表
	NowIp := c.ClientIP()
	fmt.Printf("得到的访问ip：%v", NowIp)
	var us []Models.User
	err := Untils.Db.Model(&Models.User{}).Find(&us).Error
	if err != nil {
		Untils.ResponseBadState(c, err)
	}
	Untils.ResponseOkState(c, us)
}

type Login struct {
	//必须大写
	//form:"username" json:"user" uri:"user" xml:"user" binding:"required"
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	//Token    string `form:"token" json:"token" binding:"required"`
	Nowip string `form:"nowip" json:"nowip"`
}

func UserPost(c *gin.Context) {
	//用户登录
	var formData Login
	var Data Models.User
	err := c.Bind(&formData)
	formData.Nowip = c.ClientIP()
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	err3 := Untils.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Models.User{}).Where("phone =? and password=?", formData.Phone, formData.Password).First(&Data).Error; err != nil {
			return err
		}
		token, toker := Middlewares.GenToken(Data.UserId.String, Data.Username)
		if toker != nil {
			return toker
		}
		if err2 := tx.Model(&Models.User{}).Where("user_id=?", Data.UserId.String).Update("token", token).Error; err2 != nil {
			return err2
		}
		return nil

	})
	if err3 != nil {
		Untils.ResponseBadState(c, err3)
		return
	}
	Untils.ResponseOkState(c, Data)
}

func Register(c *gin.Context) {
	//用户注册
	NowIp := c.ClientIP()
	value := ""
	var WeiUser Models.WeiChat
	err2 := c.Bind(&WeiUser)
	if err2 != nil {
		Untils.ResponseBadState(c, err2)
		return
	}
	WeiUser.VisitIP = NowIp
	Err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		er := tx.Model(&Models.WeiChat{}).Where("app_id=? AND app_code=?", WeiUser.AppId, WeiUser.AppCode).Find(&Models.WeiChat{}).RowsAffected
		if er == 0 {
			if err3 := tx.Model(&Models.WeiChat{}).Create(&WeiUser).Error; err3 != nil {
				// 返回任何错误都会回滚事务
				return err3
			}
		} else {
			if err3 := tx.Model(&Models.WeiChat{}).Update(&WeiUser).Error; err3 != nil {
				// 返回任何错误都会回滚事务
				return err3
			}
		}
		token, toker := Middlewares.GenToken(WeiUser.AppId, WeiUser.AppCode)
		if toker != nil {
			return toker
		}
		if err2 := tx.Model(&Models.WeiChat{}).Where("app_id=?", WeiUser.AppId).Update("token", token).Error; err2 != nil {
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

// Check_login 专门检测是否失效
func Check_login(c *gin.Context) {
	appid, _ := c.Get("userID")
	Err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		er := tx.Debug().Model(&Models.WeiChat{}).Where("app_id=?", appid).Find(&Models.WeiChat{}).RowsAffected
		fmt.Println(er)
		if er != 0 {
			return nil
		} else {
			return errors.New("请重新登录")
		}
	})
	if Err != nil {
		Untils.ResponseBadState(c, Err)
	} else {
		Untils.ResponseOkState(c, "登录成功")
	}
}

func Get_UserInfo(c *gin.Context) {
	user := Models.WeiChat{}
	app_id, _ := c.Get("userID")
	Err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		er := tx.Model(&Models.WeiChat{}).Where("app_id=?", app_id).Find(&user)
		if er != nil {
			return er.Error
		}
		return nil
	})
	if Err != nil {
		Untils.ResponseBadState(c, Err)
	} else {
		Untils.ResponseOkState(c, user)
	}

}

// Update_UserInfo 更新用户信息
func Update_UserInfo(c *gin.Context) {
	user := Models.WeiChat{}
	value := UpdateInfo{}
	err := c.Bind(&value)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	Err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		er := tx.Model(&Models.WeiChat{}).Where("app_code=?", value.AppCode).Find(&user).RowsAffected
		if er != 0 {
			v := tx.Model(&Models.WeiChat{}).Updates(map[string]interface{}{"avatar": value.Avator, "signature": value.Signature, "nick_name": value.Nickname}).Where("app_code=?", value.AppCode).Error
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
	AppCode   string `json:"app_code"`
	Avator    string `json:"avator"`
}

// Get_Personal_Info 获取用户详情
func Get_Personal_Info(c *gin.Context) {
	var info Models.WeiChat
	app_code := c.Query("app_code")
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		tx.Model(Models.WeiChat{}).Where("app_code=?", app_code).First(&info)
		return nil
	})
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
