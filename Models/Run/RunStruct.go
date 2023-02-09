package Run

import (
	"PetService/Models/Mine"
	"github.com/jinzhu/gorm"
)

type RunData struct {
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
	Code          string `json:"code"`
	AppCode       string `json:"app_code"`
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
	WeiChat   Mine.WeiChat `binding:"-" json:"info"`
}

func (*WeiXinRunData) TableName() string {
	return "WeiXinRunData"
}
