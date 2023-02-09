package Home

import (
	"PetService/Models/Comment"
	"PetService/Models/Mine"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type TopicDiscuss struct {
	gorm.Model
	PosterId      uint              `json:"poster_id"`
	CollegeId     string            `json:"college_id"`
	Content       string            `json:"content"`
	Attachments   arraySlice        `json:"attachments" gorm:"type:text"`
	Topic         string            `json:"topic"`
	Type          int               `json:"type"`
	Status        int               `json:"status"`
	Private       int               `json:"private"`
	CommentNumber int               `json:"comment_number"`
	PraiseNumber  int               `json:"praise_number"`
	Mobile        string            `json:"mobile"`
	NewColumn     string            `json:"new_column"`
	Praises       arraySlice        `json:"praises" gorm:"type:text"`
	Comments      []Comment.Comment `json:"comments"`
	Follow        bool              `json:"follow"`
	CanDelete     bool              `json:"can_delete"`
	CanChat       bool              `json:"can_chat"`
	Supertube     int               `json:"supertube"`
	WeiChatID     uint
	WeiChat       Mine.WeiChat `binding:"-" json:"poster"`
}

func (TopicDiscuss) TableName() string {
	return "TopicDiscuss"
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

type FollowBind struct {
	ObjId   int `json:"obj_id"`
	ObjType int `json:"obj_type"`
	UserId  int `json:"user_id"`
}
