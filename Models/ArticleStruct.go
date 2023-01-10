package Models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type Article struct {
	gorm.Model
	ArticleAuthor  int64  `gorm:"not null" json:"author" form:"author" binding:"required"`
	ArticleTitle   string `gorm:"not null" json:"title" form:"title" binding:"required"`
	ArticleContent string `gorm:"not null" json:"content" form:"content" binding:"required"`
	ArticleViews   int64  `gorm:"default:0" json:"views" form:"views"`
	ArticleGoods   int64  `gorm:"default:0" json:"goods" form:"goods"`
	ArticleImage   string `gorm:"not null" json:"article_image" form:"article_image" binding:"required"`
}

func (Article) TableName() string {
	return "Article"
}

type TopicDiscuss struct {
	gorm.Model
	PosterId      uint       `json:"poster_id"`
	CollegeId     string     `json:"college_id"`
	Content       string     `json:"content"`
	Attachments   arraySlice `json:"attachments" gorm:"type:text"`
	Topic         string     `json:"topic"`
	Type          int        `json:"type"`
	Status        int        `json:"status"`
	Private       int        `json:"private"`
	CommentNumber int        `json:"comment_number"`
	PraiseNumber  int        `json:"praise_number"`
	Mobile        string     `json:"mobile"`
	NewColumn     string     `json:"new_column"`
	Praises       arraySlice `json:"praises" gorm:"type:text"`
	Comments      arraySlice `json:"comments" gorm:"type:text"`
	Follow        bool       `json:"follow"`
	CanDelete     bool       `json:"can_delete"`
	CanChat       bool       `json:"can_chat"`
	Supertube     int        `json:"supertube"`
	AppCode       string     `json:"app_code" binding:"required"`
	WeiChatID     uint
	WeiChat       WeiChat `binding:"-" json:"poster"`
}

func (TopicDiscuss) TableName() string {
	return "TopicDiscuss"
}

type arraySlice []string
type arrayStruct struct {
}

// Scan是为了扫描数据库里面的字段然后根据设定进行返回
func (a *arraySlice) Scan(value any) error {
	fmt.Println("结果：", value)
	str, ok := value.([]byte)
	if !ok {
		return errors.New("数据类型解析失败")
	}
	newStr := string(str)
	*a = strings.Split(newStr, ",")
	return nil
}

// Values是为了存进数据库存进去的内容
func (a arraySlice) Value() (driver.Value, error) {
	if len(a) > 0 {
		value := strings.Join(a, ",")
		return value, nil
	} else {
		return "", nil
	}

}

type SaleFriend struct {
	OwnerId       uint       `json:"owner_id"`
	CollegeId     int        `json:"college_id"`
	Name          string     `json:"name"`
	Gender        int        `json:"gender"`
	Major         string     `json:"major"`
	Expectation   string     `json:"expectation"`
	Introduce     string     `json:"introduce"`
	Attachments   arraySlice `json:"attachments" gorm:"type:text"`
	ImageArray    string     `json:"imageArray" gorm:"type:text"`
	CommentNumber int        `json:"comment_number"`
	PraiseNumber  int        `json:"praise_number"`
	Type          int        `json:"type" grom:"default:1"`
	Status        int        `json:"status"`
	ID            uint       `gorm:"primary_key"`
	UpdatedAt     int64      `json:"updated_at" binding:"-"`
	CreatedAt     int64      `json:"created_at" binding:"-"`
	CanDelete     bool       `json:"can_delete"`
	CanChat       bool       `json:"can_chat"`
	Comments      []Comment  `json:"comments"`
	Follow        bool       `json:"follow"`
	AppCode       string     `json:"app_code"`
	FollowNumber  int        `json:"follow_number"`
	WeiChatID     uint
	WeiChat       WeiChat `binding:"-" json:"poster"`
}

func (SaleFriend) TableName() string {
	return "SaleFriend"
}

// Comment 评论
type Comment struct {
	gorm.Model
	CommenterId  uint       `json:"commenter_id"`
	ObjId        uint       `json:"obj_id"`
	CollegeId    int        `json:"college_id"`
	Content      string     `json:"content"`
	Attachments  arraySlice `json:"attachments" gorm:"type:text"`
	RefCommentId string     `json:"ref_comment_id"`
	ObjType      int        `json:"obj_type"`
	Type         string     `json:"type"`
	Status       int        `json:"status"`
	Author       int        `json:"author"`
	WeiChatID    uint
	WeiChat      WeiChat    `json:"commenter" binding:"-"`
	RefComment   string     `json:"ref_comment"`
	CanDelete    bool       `json:"can_delete"`
	SubComments  arraySlice `json:"sub_comments" gorm:"type:text"`
	SaleFriendID uint
}

func (Comment) TableName() string {
	return "Comment"
}
