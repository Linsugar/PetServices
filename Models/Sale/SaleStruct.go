package Sale

import (
	"PetService/Models/Comment"
	"PetService/Models/Mine"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

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

type SaleFriend struct {
	OwnerId       uint              `json:"owner_id"`
	CollegeId     int               `json:"college_id"`
	Name          string            `json:"name"`
	Gender        int               `json:"gender"`
	Major         string            `json:"major"`
	Expectation   string            `json:"expectation"`
	Introduce     string            `json:"introduce"`
	Attachments   arraySlice        `json:"attachments" gorm:"type:text"`
	ImageArray    string            `json:"imageArray" gorm:"type:text"`
	CommentNumber int               `json:"comment_number"`
	PraiseNumber  int               `json:"praise_number"`
	Type          int               `json:"type" grom:"default:1"`
	Status        int               `json:"status"`
	ID            uint              `gorm:"primary_key"`
	UpdatedAt     int64             `json:"updated_at" binding:"-"`
	CreatedAt     int64             `json:"created_at" binding:"-"`
	CanDelete     bool              `json:"can_delete"`
	CanChat       bool              `json:"can_chat"`
	Comments      []Comment.Comment `json:"comments"`
	Follow        bool              `json:"follow"`
	FollowNumber  int               `json:"follow_number"`
	WeiChatID     uint
	WeiChat       Mine.WeiChat `binding:"-" json:"poster"`
}

func (SaleFriend) TableName() string {
	return "SaleFriend"
}
