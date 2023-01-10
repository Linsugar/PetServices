package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

// 发布朋友信息
func SaleFriend(c *gin.Context) {
	var sale Models.SaleFriend
	var info Models.WeiChat
	err := c.Bind(&sale)
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	err = Untils.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Debug().Model(&Models.WeiChat{}).Where("app_code = ?", sale.AppCode).First(&info).RowsAffected
		if res > 0 {
			fmt.Println("进入1")
			sale.WeiChat = info
			sale.OwnerId = info.ID
			sale.Type = 1
			//毫秒
			sale.CreatedAt = time.Now().UnixMilli()
			sale.UpdatedAt = time.Now().UnixMilli()
			e1 := tx.Debug().Model(&Models.SaleFriend{}).Create(&sale).Error
			if e1 != nil {
				return errors.New(e1.Error())
			}
			return nil
		} else {
			return errors.New("无此用户")
		}
	})
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	Untils.ResponseOkState(c, sale)
}

// 获取朋友信息
func GetSaleFriend(c *gin.Context) {
	var model []Models.SaleFriend
	page_size, _ := strconv.Atoi(c.Query("page_size"))
	page_number, _ := strconv.Atoi(c.Query("page_number"))
	page_number = page_number - 1
	getType := c.DefaultQuery("type", "1")
	order_by := c.DefaultQuery("order_by", "created_at")
	sort_by := c.DefaultQuery("sort_by", "desc")
	app_code := c.Query("app_code")
	//just := c.Query("just")
	user_id := c.Query("user_id")

	oreder := fmt.Sprintf("%s %s", order_by, sort_by)
	var page = make(map[string]any)
	var data = make(map[string]any)
	total := 0

	var err error

	if user_id != "" {
		err = Untils.Db.Debug().Model(&Models.SaleFriend{}).Preload("Comments").Preload("WeiChat", func(tx *gorm.DB) *gorm.DB {
			return tx.Model(Models.WeiChat{}).Where("id=?", user_id)
		}).Where("type=? AND app_code=?", getType, app_code).Limit(page_size).Offset(page_number).Order(oreder).Find(&model).Error
		if err != nil {
			Untils.ResponseBadState(c, err)
		} else {
			goto label
		}
	}
	if app_code != "" {
		err = Untils.Db.Debug().Model(&Models.SaleFriend{}).Preload("Comments").Preload("WeiChat").Where("type=? AND app_code=?", getType, app_code).Limit(page_size).Offset(page_number).Order(oreder).Find(&model).Error
		if err != nil {
			Untils.ResponseBadState(c, err)
		} else {
			goto label
		}
	}
	err = Untils.Db.Debug().Model(&Models.SaleFriend{}).Preload("Comments").Preload("WeiChat").Where("type=?", getType).Limit(page_size).Offset(page_number).Order(oreder).Find(&model).Error
	if err != nil {
		Untils.ResponseBadState(c, err)
	} else {
		goto label
	}
label:
	if len(model) < 10 && len(model) != 0 {
		total = 1
	} else {
		total = len(model) / 10
	}
	page["number"] = "1"
	page["size"] = "10"
	page["total-pages"] = total
	page["total_items"] = len(model)
	data["page"] = page
	data["page_data"] = model
	Untils.ResponseOkState(c, data)

}

func FriendDetail(c *gin.Context) {
	detail := Models.SaleFriend{}
	id := c.Query("id")
	Untils.Db.Debug().Model(&Models.SaleFriend{}).Where("id=?", id).Preload("Comments.WeiChat").Preload("Comments").First(&detail)
	Untils.ResponseOkState(c, detail)
}

// AddComment 添加评论
func AddComment(c *gin.Context) {
	comment := Models.Comment{}
	app_code := "04rNbDIGuBoYcsQn"
	info := Models.WeiChat{}
	err := c.Bind(&comment)
	if err != nil {
		return
	}
	Untils.Db.Debug().Model(&Models.WeiChat{}).Where("app_code=?", app_code).First(&info)
	comment.WeiChat = info
	comment.CommenterId = info.ID
	comment.SaleFriendID = comment.ObjId
	Untils.Db.Debug().Model(&Models.Comment{}).Create(&comment)

}

// GetComment 获取评论
func GetComment(c *gin.Context) {
	comment := Models.Comment{}
	Untils.Db.Model(&Models.Comment{}).Preload("WeiChat").First(&comment)
	Untils.ResponseOkState(c, comment)
}

func CommentController(c *gin.Context) {
	if c.Request.Method == "POST" {
		AddComment(c)
	}
	if c.Request.Method == "GET" {
		GetComment(c)
	}
}

func FriendController(c *gin.Context) {
	if c.Request.Method == "POST" {
		SaleFriend(c)
	}
	if c.Request.Method == "GET" {
		GetSaleFriend(c)
	}
}

func GetNewFriends(c *gin.Context) {
	newTime := c.Query("time")
	app_code := c.Query("app_code")
	fmt.Println("当前时间：", newTime)
	var value []Models.SaleFriend
	var info Models.WeiChat
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		fmt.Println(1)
		affected := tx.Debug().Model(&Models.WeiChat{}).Where("app_code=?", app_code).First(&info).RowsAffected
		fmt.Println(affected)
		if affected >= 1 {
			var LOC, _ = time.LoadLocation("Asia/Shanghai")
			location, err := time.ParseInLocation("2006/01/02 15:04:05", newTime, LOC)
			if err != nil {
				return err
			}
			fmt.Println("x:", location.UnixMilli())
			Untils.Db.Debug().Model(&Models.SaleFriend{}).Where("created_at >= ?", location.UnixMilli()).Find(&value)
			return nil
		} else {
			return errors.New("时间不对")
		}
	})
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	Untils.ResponseOkState(c, value)
}
