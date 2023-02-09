package HomeViews

import (
	"PetService/Models/Comment"
	"PetService/Models/Home"
	"PetService/Models/Mine"
	"PetService/Untils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

// SendReleaseTopic 添加话题
func SendReleaseTopic(c *gin.Context) {
	FormData := Mine.ReleaseTopic{}
	var user Mine.WeiChat
	errs := c.Bind(&FormData)
	if errs != nil {
		Untils.ResponseBadState(c, errs)
	}
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		e1 := tx.Where("app_code=?", FormData.AppCode).Find(&user).Error
		if e1 != nil {
			return errors.New("非法闯入-当前账号不存在")
		}
		FormData.WeiChatID = user.ID
		FormData.UserType = user.Type
		e2 := Untils.Db.Model(&Mine.ReleaseTopic{}).Create(&FormData).Error
		if e2 != nil {
			return errors.New("写入失败")
		}
		return nil
	})
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	Untils.ResponseOkState(c, FormData)
}

func GetReleaseTopic(c *gin.Context) {
	order_by := c.DefaultQuery("order_by", "created_at")
	sort_by := c.DefaultQuery("sort_by", "desc")

	oreder := fmt.Sprintf("%s %s", order_by, sort_by)
	FormData := Mine.ReleaseTopic{}
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		e1 := tx.Model(&Mine.ReleaseTopic{}).Order(oreder).Find(&FormData).Error
		if e1 != nil {
			return errors.New("非法闯入-当前账号不存在")
		}
		return nil
	})
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	Untils.ResponseOkState(c, FormData)
}

// AddList 添加当前发布信息
func AddList(c *gin.Context) {
	model := Home.TopicDiscuss{}
	info := Mine.WeiChat{}
	errs := c.Bind(&model)
	if errs != nil {
		fmt.Println("绑定失败")
		Untils.ResponseBadState(c, errs)
		return
	}
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		uid, _ := c.Get("userID")
		e1 := tx.Debug().Model(&Mine.WeiChat{}).Where("open_id=?", uid).First(&info).Error
		if e1 != nil {
			return errors.New("非法闯入-当前账号不存在")
		}
		model.WeiChat = info
		model.PosterId = info.ID
		model.Type = 1
		e2 := tx.Model(&Home.TopicDiscuss{}).Debug().Create(&model).Error
		if e2 != nil {
			return errors.New("写入失败")
		}
		return nil
	})
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	Untils.ResponseOkState(c, model)
}

// GetList 获取当前所有发布信息
func GetList(c *gin.Context) {
	var model []Home.TopicDiscuss
	page_size, _ := strconv.Atoi(c.Query("page_size"))
	page_number, _ := strconv.Atoi(c.Query("page_number"))
	getType := c.DefaultQuery("type", "1")
	order_by := c.DefaultQuery("order_by", "created_at")
	sort_by := c.DefaultQuery("sort_by", "desc")
	//just := c.Query("just")
	user_id := c.Query("user_id")
	filter := c.Query("filter")
	order := fmt.Sprintf("%s %s", order_by, sort_by)
	var page = make(map[string]any)
	var data = make(map[string]any)
	total := 0

	var count int
	if page_number <= 0 {
		page_number = 1
	}
	if page_size <= 0 {
		page_size = 10
	}
	var svalue string
	if user_id != "" && filter == "" {
		svalue = fmt.Sprintf("type=%s AND poster_id=%s", getType, user_id)
	} else {
		svalue = fmt.Sprintf("type=%s", getType)
	}
	if filter != "" {
		svalue = fmt.Sprintf("type=%s AND AND content LIKE %s", getType, filter+"%")
	}
	db := Untils.Db.Debug().Model(&Home.TopicDiscuss{}).Preload("Comments").
		Preload("Comments.WeiChat").
		Preload("Comments.RefComment", "type=1").
		Preload("Comments.RefComment.WeiChat").
		Preload("WeiChat").Where(svalue)
	db.Count(&count)
	err := db.Limit(page_size).Offset((page_number - 1) * page_size).Order(order).Find(&model).Error
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
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

func TalkListController(c *gin.Context) {
	if c.Request.Method == "POST" {
		AddList(c)
	} else if c.Request.Method == "GET" {
		GetList(c)
	}
}

func TopicController(c *gin.Context) {
	if c.Request.Method == "POST" {
		SendReleaseTopic(c)
	} else if c.Request.Method == "GET" {
		GetReleaseTopic(c)
	}
}

func GetNewTopic(c *gin.Context) {
	app_cdde := c.Query("app_code")
	fmt.Println(app_cdde)
	value := make(map[string]any, 0)
	value["data"] = 0
	value["error_code"] = 0
	value["error_message"] = "success"
	Untils.ResponseOkState(c, value)
}

func CollectPostList(c *gin.Context) {
	var collectData Home.FollowBind
	var userInfo Mine.WeiChat
	var homeList Comment.CollectHomeList
	err := c.Bind(&collectData)
	if err != nil {
		Untils.Error.Println(err.Error())
		return
	}
	Untils.Db.Model(userInfo).Where("id=?", collectData.UserId).First(&userInfo)
	homeList.WeiChat = userInfo
	rowsAffected := Untils.Db.Model(homeList).Where("wei_chat_id=?", userInfo.ID).Find(&homeList).RowsAffected
	var us []int
	if rowsAffected == 0 {
		us = append(us, collectData.ObjId)
		homeList.List = us
		homeList.Type = collectData.ObjType
		Untils.Db.Create(&homeList)
	} else {
		Untils.Db.Model(homeList).Update(map[string]any{"list": us}).Where("wei_chat_id=?", userInfo.ID)
	}
}
