package SaleViews

import (
	"PetService/Models/Comment"
	"PetService/Models/Mine"
	"PetService/Models/Sale"
	"PetService/Untils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

// 发布朋友信息
func SaleFriend(c *gin.Context) {
	var sale Sale.SaleFriend
	var info Mine.WeiChat
	err := c.Bind(&sale)
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	err = Untils.Db.Transaction(func(tx *gorm.DB) error {
		uid, _ := c.Get("userID")
		res := tx.Debug().Model(&Mine.WeiChat{}).Where("open_id = ?", uid).First(&info).RowsAffected
		if res > 0 {
			sale.WeiChat = info
			sale.OwnerId = info.ID
			sale.Type = 1
			//毫秒
			sale.CreatedAt = time.Now().UnixMilli()
			sale.UpdatedAt = time.Now().UnixMilli()
			e1 := tx.Debug().Model(&Sale.SaleFriend{}).Create(&sale).Error
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
	var model []Sale.SaleFriend
	page_size, _ := strconv.Atoi(c.Query("page_size"))
	page_number, _ := strconv.Atoi(c.Query("page_number"))
	Untils.Info.Println(page_number)
	getType := c.DefaultQuery("type", "1")
	order_by := c.DefaultQuery("order_by", "created_at")
	sort_by := c.DefaultQuery("sort_by", "desc")
	//just := c.Query("just")
	user_id := c.Query("user_id")
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
	var err error
	var db *gorm.DB
	var svalue string
	if user_id != "" {
		svalue = fmt.Sprintf("type=%s AND owner_id=%s", getType, user_id)
		db = Untils.Db.Debug().Model(&Sale.SaleFriend{}).Preload("Comments").
			Preload("WeiChat", func(tx *gorm.DB) *gorm.DB {
				return tx.Model(Mine.WeiChat{})
			}).Where(svalue)
		db.Count(&count)
	} else {
		svalue = fmt.Sprintf("type=%s", getType)
		Untils.Info.Println(page_number)
		db = Untils.Db.Debug().Model(&Sale.SaleFriend{}).Preload("Comments").Preload("Comments.SubComments").
			Preload("Comments.SubComments.WeiChat").Preload("WeiChat").Where(svalue)
	}
	err = db.Limit(page_size).Offset((page_number - 1) * page_size).Order(order).Find(&model).Error
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

func FriendDetail(c *gin.Context) {
	detail := Sale.SaleFriend{}
	id := c.Query("id")
	Untils.Db.Debug().Model(&detail).Where("id=?", id).
		Preload("Comments").Preload("Comments.WeiChat").
		Preload("Comments.SubComments").Preload("Comments.SubComments.WeiChat").
		Preload("Comments.SubComments.RefComment").Preload("Comments.SubComments.RefComment.WeiChat").First(&detail)
	Untils.ResponseOkState(c, detail)
}

// AddComment 添加评论
func AddComment(c *gin.Context) {
	comment := Comment.Comment{}
	appCode := "04rNbDIGuBoYcsQn"
	info := Mine.WeiChat{}
	err := c.ShouldBindBodyWith(&comment, binding.JSON)
	if err != nil {
		Untils.Error.Println(err.Error())
		return
	}
	Untils.Db.Debug().Model(&info).Where("app_code=?", appCode).First(&info)
	comment.WeiChat = info
	comment.CommenterId = info.ID
	Untils.Info.Println("id:===> ", comment.Type)
	if comment.Type == "1" {
		comment.TopicDiscussID = comment.ObjId
		if comment.RefCommentId > 0 {
			ref := &Comment.RefComment{}
			er1 := c.ShouldBindBodyWith(&ref, binding.JSON)
			if er1 != nil {
				fmt.Println(er1.Error())
				return
			}
			Untils.Info.Println(comment.RefCommentId)
			ref.WeiChat = info
			ref.CommenterId = info.ID
			ref.CommentID = comment.RefCommentId
			Untils.Info.Println("找到多少条数据之前：", comment.RefCommentId)
			Untils.Db.Debug().Create(&ref)
			Untils.ResponseOkState(c, ref)
			return
		}
	}
	if comment.Type == "2" || comment.Type == "4" {
		comment.SaleFriendID = comment.ObjId
		if comment.RefCommentId > 0 {
			saleInfo := Comment.SaleComment{}
			affected := Untils.Db.Debug().Model(&saleInfo).Where("id =?", comment.RefCommentId).Find(&saleInfo).RowsAffected
			if comment.RefCommentId == comment.ObjId && affected == 0 {
				sale := &saleInfo
				er1 := c.ShouldBindBodyWith(&sale, binding.JSON)
				sale.WeiChat = info
				sale.CommenterId = info.ID
				sale.SaleFriendID = comment.RefCommentId
				sale.CommentID = comment.RefCommentId
				if er1 != nil {
					fmt.Println(er1.Error())
					return
				}
				Untils.Db.Debug().Model(&saleInfo).Create(&sale)
				Untils.ResponseOkState(c, sale)
				return
			} else {
				ref := &Comment.RefComment{}
				er1 := c.ShouldBindBodyWith(&ref, binding.JSON)
				ref.WeiChat = info
				ref.CommenterId = info.ID
				ref.SaleFriendID = comment.RefCommentId
				ref.SaleCommentID = comment.RefCommentId
				if er1 != nil {
					fmt.Println(er1.Error())
					return
				}
				Untils.Info.Println("找到多少条数据之前：", comment.RefCommentId)
				Untils.Db.Debug().Model(ref).Create(&ref)
				Untils.ResponseOkState(c, ref)
				return
			}
		}
	}
	Untils.Db.Debug().Model(&Comment.Comment{}).Create(&comment)
	Untils.ResponseOkState(c, comment)

}

// GetComment 获取评论
func GetComment(c *gin.Context) {
	comment := Comment.Comment{}
	Untils.Db.Model(comment).Preload("WeiChat").First(&comment)
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
	var value []Sale.SaleFriend
	var info Mine.WeiChat
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		fmt.Println(1)
		affected := tx.Debug().Model(info).Where("app_code=?", app_code).First(&info).RowsAffected
		fmt.Println(affected)
		if affected >= 1 {
			var LOC, _ = time.LoadLocation("Asia/Shanghai")
			location, err := time.ParseInLocation("2006/01/02 15:04:05", newTime, LOC)
			if err != nil {
				return err
			}
			fmt.Println("x:", location.UnixMilli())
			Untils.Db.Debug().Model(value).Where("created_at >= ?", location.UnixMilli()).Find(&value)
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
