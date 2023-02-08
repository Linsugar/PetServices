package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	//用户信息表
	gorm.Model
	Username       string `gorm:"not null" json:"username" binding:"required" form:"username"`
	Password       string `json:"password" binding:"required" form:"password"`
	Phone          string `gorm:"unique;index:phone" json:"phone" binding:"required" form:"phone"`
	CreateCity     string `gorm:"default:'成都'" json:"create_city" form:"create_city"`
	CreateAddress  string `gorm:"default:'成都高新区'"`
	InitIp         string
	NowIp          string
	Token          string         `gorm:"column:token;"`
	IsDel          bool           `gorm:"column:isdel;default:false"`                                                          //是否删除
	UserId         sql.NullString `gorm:"unique;unique_index;not null"`                                                        //不重复id
	InvitePerson   string         `gorm:"default:'6666'" json:"invitePerson" form:"invitePerson" binding:"required"`           //邀请人id某人为空
	ProfilePicture string         `gorm:"default:'http://cdn.tlapp.club/pet.png'" json:"profilePicture" form:"profilePicture"` //头像地址
	UserContent    string         `gorm:"default:''" json:"userContent" form:"userContent"`                                    //用户简介
	UserCode       string         `gorm:"" json:"UserCode" form:"UserCode"`
	UserDevice     string         `gorm:"index:devices" json:"userDevice" form:"userDevice" binding:"required"` //用户设备
}

// 自定义表名-默认是结构体名称+s
func (User) TableName() string {
	return "User"
}

type PetDetail struct {
	//宠物资料详细表
	gorm.Model
	PetID       int64         `grom:"not null;unique;index:pet"`
	PetMaster   string        `gorm:"not null" json:"petMaster" binding:"required" form:"petMaster"` //宠物主人的id
	PetName     string        `json:"pet_name" form:"pet_name" binding:"required" `
	PetCall     string        `gorm:"default:'无'" json:"petCall" form:"petCall"`                                                        //联系方式
	Petdetail   string        `gorm:"default:'暂无介绍'" json:"petdetail" form:"petdetail"`                                                 //宠物详细介绍
	PetClass    string        `gorm:"not null;default:'0'" json:"petClass" form:"petClass"`                                             //宠物类型
	PetBuyer    sql.NullInt32 `json:"petBuyer" form:"PetBuyer"`                                                                         //买主id 默认为空
	PetPhoto    string        `gorm:"default:'暂无'" json:"petPhoto" form:"petPhoto"`                                                     //宠物相册
	PetAvatotr  string        `gorm:"default:'http://cdn.tlapp.club/pet.png'" form:"pet_avatotr" json:"pet_avatotr" binding:"required"` //宠物头像
	PetVideo    string        `gorm:"default:'暂无'" json:"pet_video" form:"pet_video"`                                                   //视频地址
	PetMoney    float64       `gorm:"default:'0.0'" json:"petMoney" form:"petMoney"`                                                    //最初定价
	PetBuyMoney float64       `gorm:"default:'0.0'" json:"petBuyMoney" form:"petBuyMoney"`                                              //最终售卖价
	PetContent  string        `json:"petContent" form:"petContent" binding:"required"`                                                  //最终售卖价
	PetAge      float64       `json:"petAge" form:"petAge" binding:"required"`                                                          //最终售卖价
	PetGender   string        `gorm:"defalut:'MALE'" json:"petGender" form:"petGender"`                                                 //最终售卖价
	PetWeight   float32       `gorm:"default:'0.0'" json:"petWeight" form:"petWeight"`                                                  //最终售卖价
	PetLocation string        `gorm:"defalut:'[10.0,20]'" json:"petLocation" form:"petLocation"`
	PetSex      int           `gorm:"default:'0'" json:"pet_sex" form:"pet_sex"`
}

type RegisterCode struct {
	gorm.Model
	Code       string
	CodeDevice string
	CodeIp     string
}

func (RegisterCode) TableName() string {
	return "RegisterCode"
}

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
