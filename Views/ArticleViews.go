package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

func SendReleaseTopic(c *gin.Context) {
	FormData := Models.ReleaseTopic{}
	errs := c.Bind(&FormData)
	if errs != nil {
		Untils.ResponseBadState(c, errs)
	}
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		e1 := tx.Where("app_code=?", FormData.AppCode).Find(&Models.WeiChat{}).Error
		if e1 != nil {
			return errors.New("非法闯入-当前账号不存在")
		}
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

//TopicList 获取当前话题信息
func TopicList(c *gin.Context) {
	model := Models.TopicDiscuss{}
	info := Models.WeiChat{}
	errs := c.Bind(&model)
	if errs != nil {
		Untils.ResponseBadState(c, errs)
	}
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		e1 := tx.Debug().Model(&Models.WeiChat{}).Where("app_code=?", model.AppCode).Find(&info).Error
		if e1 != nil {
			return errors.New("非法闯入-当前账号不存在")
		}
		e2 := Untils.Db.Model(&Models.TopicDiscuss{}).Create(&model).Error
		if e2 != nil {
			return errors.New("写入失败")
		}
		return nil
	})
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	model.Poster.Nickname = info.NickName
	model.Poster.Id = info.ID
	model.Poster.Avatar = info.Avator
	model.Poster.CreatedAt = info.CreatedAt
	model.PosterId = int(info.ID)
	Untils.ResponseOkState(c, model)
}

func GetTopicList(c *gin.Context) {
	//var model []Models.TopicDiscuss
	//info := []Models.WeiChat{}
	//page_size, _ := strconv.Atoi(c.Query("page_size"))
	//page_number, _ := strconv.Atoi(c.Query("page_number"))
	//getType := c.Query("type")
	//order_by := c.Query("order_by")
	//sort_by := c.Query("sort_by")
	//app_code := c.Query("app_code")

	var users []Models.WeiChat
	err := Untils.Db.Model(&Models.WeiChat{}).Preload("TopicDiscuss").Find(&users).Error
	fmt.Println(err)
	fmt.Println(users)
	//err := Untils.Db.Transaction(func(tx *gorm.DB) error {
	//	var e2 error
	//	//oreder := fmt.Sprintf("%s %s", order_by, sort_by)
	//	//if app_code != "" {
	//	//	e2 = tx.Model(&Models.TopicDiscuss{}).Where("type=? and app_code=?", getType, app_code).Limit(page_size).Offset(page_number).Order(oreder).Find(&model).Error
	//	//} else {
	//	//e2 = tx.Model(&Models.TopicDiscuss{}).Where("type=?", getType).Limit(page_size).Offset(page_number).Order(oreder).Find(&model).Error
	//	e2 = tx.Model(&Models.WeiChat{}).Preload("TopicDiscusss").Find(&info).Error
	//	//}
	//
	//	//e2 = tx.Model(&Models.TopicDiscuss{}).Association("TopicDiscuss").Find(&info.TopicDiscusss).Error
	//
	//	if e2 != nil {
	//		return errors.New("写入失败")
	//	}
	//	return info, err
	//	//for _, discuss := range model {
	//	//	e2 = tx.Debug().Model(&Models.WeiChat{}).Where("app_code=?", discuss.AppCode).First(&info).Error
	//	//	discuss.Poster.Nickname = info.NickName
	//	//	discuss.Poster.Id = info.ID
	//	//	discuss.Poster.Avatar = info.Avator
	//	//	discuss.Poster.CreatedAt = info.CreatedAt
	//	//	discuss.PosterId = int(info.ID)
	//	//}
	//	return nil
	//})
	//if err != nil {
	//	Untils.ResponseBadState(c, err)
	//	return
	//}
	Untils.ResponseOkState(c, users)
}

func TopicController(c *gin.Context) {
	if c.Request.Method == "POST" {
		TopicList(c)
	} else if c.Request.Method == "GET" {
		GetTopicList(c)
	}
}

func UserTest(c *gin.Context) {
	//var U1 []Models.UTest
	//var U = Models.UTest{
	//	Name:  "test",
	//	CTest: []Models.CTest{{Code: "222", Name: "c2"}, {Code: "333", Name: "c3"}},
	//}

	//Untils.Db.Create(&U)
	//Untils.Db.Model(&Models.UTest{}).Preload("CTest").Find(&U1, 2)
	//Untils.ResponseOkState(c, &U1)
}
