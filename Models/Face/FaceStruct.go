package Face

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

type SessionData struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}
