package Models

import (
	"github.com/jinzhu/gorm"
)

// FaceMerge 人脸融合传入的参数
type FaceMerge struct {
	YourFace string `json:"your_face"`
	HisFace  string `json:"his_face"`
	AppCode  string `json:"app_code"`
}

// FaceAge 人脸年龄传入的参数
type FaceAge struct {
	Image   string `json:"image"`
	AppCode string `json:"app_code"`
	Age     int    `json:"age"`
}

type RunData struct {
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
	Code          string `json:"code"`
	AppCode       string `json:"app_code"`
}

type SessionData struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

type RunResultData struct {
	StepInfoList []struct {
		Timestamp int `json:"timestamp"`
		Step      int `json:"step"`
	}
}

type WeiXinRunData struct {
	gorm.Model
	AppId     string `gorm:"not null" json:"app_id" binding:"required"`
	Today     int    `json:"today"`
	TotalNum  int    `json:"total_num"`
	WeiChatID uint
	WeiChat   WeiChat `binding:"-" json:"info"`
}

func (*WeiXinRunData) TableName() string {
	return "WeiXinRunData"
}
