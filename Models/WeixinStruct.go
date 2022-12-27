package Models

import "github.com/jinzhu/gorm"

type T2 struct {
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
	Item       []struct {
		MediaId string `json:"media_id"`
		Content struct {
			NewsItem []struct {
				Title            string `json:"title"`
				ThumbMediaId     string `json:"thumb_media_id"`
				ShowCoverPic     int    `json:"show_cover_pic"`
				Author           string `json:"author"`
				Digest           string `json:"digest"`
				Content          string `json:"content"`
				Url              string `json:"url"`
				ContentSourceUrl string `json:"content_source_url"`
				Thumb_url        string `json:"thumb_url"`
			} `json:"news_item"`
		} `json:"content"`
		UpdateTime int `json:"update_time"`
	} `json:"item"`
}
type Test1 struct {
	gorm.Model
	Name    string
	Test2ID uint
	Test2   Test2
}

func (Test1) TableName() string {
	return "Test1"
}

type Test2 struct {
	gorm.Model
	Code string
	Name string
}

func (Test2) TableName() string {
	return "Test2"
}
