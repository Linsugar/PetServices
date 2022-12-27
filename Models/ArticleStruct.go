package Models

import (
	"github.com/jinzhu/gorm"
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
	PosterId      uint   `json:"poster_id"`
	CollegeId     string `json:"college_id"`
	Content       string `json:"content"`
	Attachments   string `json:"attachments"`
	Topic         string `json:"topic"`
	Type          int    `json:"type"`
	Status        int    `json:"status"`
	Private       int    `json:"private"`
	CommentNumber int    `json:"comment_number"`
	PraiseNumber  int    `json:"praise_number"`
	Mobile        string `json:"mobile"`
	NewColumn     string `json:"new_column"`
	Praises       int    `json:"praises"`
	Comments      int    `json:"comments"`
	Follow        bool   `json:"follow"`
	CanDelete     bool   `json:"can_delete"`
	CanChat       bool   `json:"can_chat"`
	Supertube     int    `json:"supertube"`
	AppCode       string `json:"app_code" binding:"required"`
	WeiChatID     uint
	WeiChat       WeiChat `binding:"-" json:"poster"`
}

func (TopicDiscuss) TableName() string {
	return "TopicDiscuss"
}
