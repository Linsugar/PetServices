package Comment

import (
	"PetService/Models/Mine"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

// Comment 评论
type Comment struct {
	gorm.Model
	CommenterId    uint       `json:"commenter_id"`
	ObjId          uint       `json:"obj_id"`
	CollegeId      int        `json:"college_id"`
	Content        string     `json:"content"`
	Attachments    arraySlice `json:"attachments" gorm:"type:text"`
	RefCommentId   uint       `json:"ref_comment_id"`
	ObjType        int        `json:"obj_type"`
	Type           string     `json:"type"`
	Status         int        `json:"status"`
	Author         int        `json:"author"`
	CanDelete      bool       `json:"can_delete"`
	WeiChatID      uint
	WeiChat        Mine.WeiChat  `json:"commenter" binding:"-"`
	RefComment     []RefComment  `json:"ref_comment"`
	SubComments    []SaleComment `json:"sub_comments"`
	SaleFriendID   uint          `json:"-"`
	TopicDiscussID uint          `json:"-"`
}

type RefComment struct {
	gorm.Model
	Attachments   arraySlice `json:"attachments" gorm:"type:text"`
	CommenterId   uint       `json:"commenter_id"`
	CollegeId     int        `json:"college_id"`
	Content       string     `json:"content"`
	ObjType       int        `json:"obj_type"`
	ObjId         uint       `json:"obj_id"`
	Type          string     `json:"type"`
	Status        int        `json:"status"`
	WeiChatID     uint
	WeiChat       Mine.WeiChat `json:"refCommenter" binding:"-"`
	SaleFriendID  uint
	SaleCommentID uint
	//ListCommentID uint
	CommentID uint
}

type SaleComment struct {
	gorm.Model
	Attachments  arraySlice `json:"attachments" gorm:"type:text"`
	CommenterId  uint       `json:"commenter_id"`
	CollegeId    int        `json:"college_id"`
	Content      string     `json:"content"`
	ObjType      int        `json:"obj_type"`
	ObjId        uint       `json:"obj_id"`
	Type         string     `json:"type"`
	Status       int        `json:"status"`
	WeiChatID    uint
	WeiChat      Mine.WeiChat `json:"refCommenter" binding:"-"`
	CommentID    uint
	RefComment   []RefComment `json:"ref_comment"`
	SaleFriendID uint
}

type ListComment struct {
	gorm.Model
	Attachments arraySlice `json:"attachments" gorm:"type:text"`
	CommenterId uint       `json:"commenter_id"`
	CollegeId   int        `json:"college_id"`
	Content     string     `json:"content"`
	ObjType     int        `json:"obj_type"`
	ObjId       uint       `json:"obj_id"`
	Type        string     `json:"type"`
	Status      int        `json:"status"`
	WeiChatID   uint
	WeiChat     Mine.WeiChat `json:"refCommenter" binding:"-"`
	CommentID   uint
	RefComment  []RefComment `json:"ref_comment"`
}

func (Comment) TableName() string {
	return "Comment"
}

func (RefComment) TableName() string {
	return "RefComment"
}

func (SaleComment) TableName() string {
	return "SaleComment"
}

func (ListComment) TableName() string {
	return "ListComment"
}

type intArraySlice []int

// Scan是为了扫描数据库里面的字段然后根据设定进行返回
func (a *intArraySlice) Scan(value any) error {
	fmt.Println("结果：", value)
	var intlist = make([]int, 0)
	str, ok := value.(string)
	if !ok {
		return errors.New("数据类型解析失败")
	} else {
		v := []byte(str)
		err := json.Unmarshal(v, &intlist)
		if err != nil {
			return err
		}
	}
	*a = intlist
	return nil
}

// Values是为了存进数据库存进去的内容
func (a intArraySlice) Value() (driver.Value, error) {
	if len(a) > 0 {
		marshal, err := json.Marshal(a)
		if err != nil {
			fmt.Println("解析失败：", err.Error())
			return nil, err
		}
		value := string(marshal)
		return value, nil
	} else {
		return "", nil
	}

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

//CollectHomeList 收藏首页用户动态
type CollectHomeList struct {
	gorm.Model
	List      intArraySlice `json:"list" gorm:"type:text"`
	WeiChatID uint
	WeiChat   Mine.WeiChat `json:"user" binding:"-"`
	Type      int          `json:"type"`
}

//CollectSaleList 收藏分享动态
type CollectSaleList struct {
	gorm.Model
	List      intArraySlice `json:"list" gorm:"type:text"`
	WeiChatID uint
	WeiChat   Mine.WeiChat `json:"user" binding:"-"`
	Type      int          `json:"type"`
}

//CollectUserList 关注的用户
type CollectUserList struct {
	gorm.Model
	List      intArraySlice `json:"list" gorm:"type:text"`
	WeiChatID uint
	WeiChat   Mine.WeiChat `json:"user" binding:"-"`
	Type      int          `json:"type"`
}
