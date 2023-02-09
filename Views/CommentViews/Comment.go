package CommentViews

import (
	"PetService/Models/Mine"
	"PetService/Untils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetQINiuToken(c *gin.Context) {
	if c.Request.Method == "POST" {
		token := Untils.QiNiuToken()
		Untils.ResponseOkState(c, token)
	}
}

// Check_login 专门检测是否失效
func Check_login(c *gin.Context) {
	appid, _ := c.Get("userID")
	Err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		er := tx.Debug().Model(&Mine.WeiChat{}).Where("app_id=?", appid).Find(&Mine.WeiChat{}).RowsAffected
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
