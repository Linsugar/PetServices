package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

//文章控制器

func ArticleController(c *gin.Context) {
	if c.Request.Method == "POST" {
		ArticlePost(c)
	} else if c.Request.Method == "GET" {
		ArticleAll(c)
	}
}

func ArticleAll(c *gin.Context) {
	fmt.Println("进入")
	var Article []Models.Article
	if err := Untils.Db.Model(&Models.Article{}).Find(&Article).Error; err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	Untils.ResponseOkState(c, Article)
}

func ArticlePost(c *gin.Context) {
	FormData := Models.Article{}
	errs := c.Bind(&FormData)
	fmt.Println("绑定有误:", errs)
	if errs != nil {
		Untils.ResponseBadState(c, errs)
	}
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		if e1 := tx.Where("user_id=?", FormData.ArticleAuthor).Find(&Models.User{}).Error; e1 != nil {
			return e1
		}
		if e2 := Untils.Db.Model(&Models.Article{}).Create(&FormData).Error; e2 != nil {
			return e2
		}
		return nil
	})
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}

	Untils.ResponseOkState(c, FormData)
}

// SendReleaseTopic 添加话题
func SendReleaseTopic(c *gin.Context) {
	FormData := Models.ReleaseTopic{}
	var user Models.WeiChat
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
		e2 := Untils.Db.Model(&Models.ReleaseTopic{}).Create(&FormData).Error
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
	FormData := Models.ReleaseTopic{}
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		e1 := tx.Model(&Models.ReleaseTopic{}).Order(oreder).Find(&FormData).Error
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

// AddTopicList 添加当前发布信息
func AddTopicList(c *gin.Context) {
	model := Models.TopicDiscuss{}
	info := Models.WeiChat{}
	errs := c.Bind(&model)
	if errs != nil {
		fmt.Println("绑定失败")
		Untils.ResponseBadState(c, errs)
		return
	}
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		e1 := tx.Debug().Model(&Models.WeiChat{}).Where("app_code=?", model.AppCode).First(&info).Error
		if e1 != nil {
			return errors.New("非法闯入-当前账号不存在")
		}
		model.WeiChat = info
		model.PosterId = info.ID
		e2 := tx.Model(&Models.TopicDiscuss{}).Create(&model).Error
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

// GetTopicList 获取当前所有发布信息
func GetTopicList(c *gin.Context) {
	var model []Models.TopicDiscuss
	page_size, _ := strconv.Atoi(c.Query("page_size"))
	page_number, _ := strconv.Atoi(c.Query("page_number"))
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
		err = Untils.Db.Debug().Model(&Models.TopicDiscuss{}).Preload("WeiChat", func(tx *gorm.DB) *gorm.DB {
			return tx.Model(Models.WeiChat{}).Where("id=?", user_id)
		}).Where("type=? AND app_code=?", getType, app_code).Limit(page_size).Offset(page_number).Order(oreder).Find(&model).Error
		if err != nil {
			Untils.ResponseBadState(c, err)
		} else {
			goto label
		}
	}
	if app_code != "" {
		err = Untils.Db.Debug().Model(&Models.TopicDiscuss{}).Preload("WeiChat").Where("type=? AND app_code=?", getType, app_code).Limit(page_size).Offset(page_number).Order(oreder).Find(&model).Error
		if err != nil {
			Untils.ResponseBadState(c, err)
		} else {
			goto label
		}
	}
	err = Untils.Db.Debug().Model(&Models.TopicDiscuss{}).Preload("WeiChat").Where("type=?", getType).Limit(page_size).Offset(page_number).Order(oreder).Find(&model).Error
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

func TalkListController(c *gin.Context) {
	if c.Request.Method == "POST" {
		AddTopicList(c)
	} else if c.Request.Method == "GET" {
		GetTopicList(c)
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
