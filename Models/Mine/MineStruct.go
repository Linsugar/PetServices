package Mine

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type UserToPic struct {
	//话题-列表
	gorm.Model
	UserId        int      `json:"user_id"`
	AppId         int      `json:"app_id"`
	UserType      int      `json:"user_type"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	Attachments   []string `json:"attachments"`
	PraiseNumber  int      `json:"praise_number"`
	ViewNumber    int      `json:"view_number"`
	CommentNumber int      `json:"comment_number"`
	Status        int      `json:"status"`
}

func (UserToPic) TableName() string {
	return "UserToPic"
}

// WeiChat 微信用户信息
type WeiChat struct {
	gorm.Model
	Code        string `json:"code"`
	Iv          string `json:"iv"`
	AppId       string `json:"app_id" binding:"required"`
	AppCode     string `json:"app_code"`
	Token       string `json:"token"`
	Avatar      string `json:"avatar" gorm:"default:'http://cdn.tlapp.club/pet.png'"`
	NickName    string `json:"nickName" gorm:"default:'访客'"`
	VisitIP     string `json:"visitIP"`
	Signature   string `json:"personal_signature"`
	Gender      bool   `json:"gender"`
	Type        string `json:"type" gorm:"default:'0'"`
	UnionId     string `json:"Union_id"`
	Status      string `json:"status"`
	ActiveValue int    `json:"active_value"`
	City        string `json:"city"`
	ClockNum    int    `json:"clock_num"`
	Country     string `json:"country"`
	FansNum     int    `json:"fans_num"`
	FollowNum   int    `json:"follow_num"`
	Language    string `json:"language"`
	Mobile      string `json:"mobile"`
	OpenId      string `json:"open_id"`
	PostNum     int    `json:"post_num"`
	Province    string `json:"province"`
	Supertube   int    `json:"supertube"`
}

func (WeiChat) TableName() string {
	return "WeiChat"
}

// ReleaseTopic 发布话题
type ReleaseTopic struct {
	gorm.Model
	Content       string     `json:"content" binding:"required" gorm:"not null"`
	Attachments   arraySlice `json:"attachments" gorm:"type:text"`
	Private       bool       `json:"private"`
	AppCode       string     `json:"app_code" binding:"required" gorm:"not null"`
	Type          int        `json:"type" gorm:"type:text"`
	PraiseNumber  int        `json:"praise_number"`
	Status        int        `json:"status"`
	Title         string     `json:"title"`
	UserType      string     `json:"user_type"`
	ViewNumber    int        `json:"view_number"`
	CommentNumber int        `json:"comment_number"`
	WeiChatID     uint       `json:"user_id"`
	WeiChat       WeiChat    `binding:"-" json:"-"`
}

func (ReleaseTopic) TableName() string {
	return "ReleaseTopic"
}

type arraySlice []string

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
