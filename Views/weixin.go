package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	appid  string
	secret string
)

func GetWeiXinSession(code string) Models.SessionData {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, code)
	get, err := http.Get(url)
	if err != nil {
		Untils.Error.Println(err.Error())
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		Untils.Error.Println(err.Error())
	}
	var getSession Models.SessionData
	err = json.Unmarshal(all, &getSession)
	if err != nil {
		Untils.Error.Println(err.Error())
	}
	Untils.Info.Println(getSession.SessionKey)
	return getSession
}

// WeiXinDecode 加密的数据解密
func WeiXinDecode(encryptedData, sessionKey, iv string) any {
	encrypted, _ := base64.StdEncoding.DecodeString(encryptedData)
	keyB, _ := base64.StdEncoding.DecodeString(sessionKey)
	ivB, _ := base64.StdEncoding.DecodeString(iv)
	block, _ := aes.NewCipher(keyB)                 // 分组秘钥
	blockMode := cipher.NewCBCDecrypter(block, ivB) // 加密模式
	decrypted := make([]byte, len(encrypted))       // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)     // 解密
	Untils.Info.Println(string(decrypted))
	value := strings.Split(string(decrypted), "\u000E")
	Untils.Info.Println(value[0])
	return value[0]

}

func GetWeiXinRunData(c *gin.Context) {
	run := Models.RunData{}
	err := c.Bind(&run)
	var runData Models.WeiXinRunData
	var userInfo Models.WeiChat
	if err != nil {
		Untils.Error.Println(err.Error())
		return
	}
	valueSession := GetWeiXinSession(run.Code)
	Untils.Db.Debug().Model(&userInfo).Where("open_id=?", valueSession.Openid).First(&userInfo)
	runData.AppId = valueSession.Openid
	runData.WeiChat = userInfo
	Result := WeiXinDecode(run.EncryptedData, valueSession.SessionKey, run.Iv)
	returnMap := Models.RunResultData{}
	json.Unmarshal([]byte(Result.(string)), &returnMap)
	Untils.Info.Println("进入这里2：", valueSession.Openid)
	affected := Untils.Db.Model(&runData).Where("app_id=?", valueSession.Openid).Find(&runData).RowsAffected
	step := len(returnMap.StepInfoList)
	runData.Today = returnMap.StepInfoList[step-1].Step
	totalnum := runData.TotalNum
	for i := 0; i < step; i++ {
		totalnum += returnMap.StepInfoList[i].Step
	}
	runData.TotalNum = totalnum
	if affected < 1 {
		Untils.Db.Debug().Model(&runData).Create(&runData)
		Untils.ResponseOkState(c, runData)
		return
	} else {
		Untils.Db.Model(&runData).Update(map[string]any{
			"today":     runData.Today,
			"total_num": runData.TotalNum,
		}).Where("app_id=?", valueSession.Openid)
		Untils.ResponseOkState(c, runData)
		return
	}

}

// GetUserRunData 独立获取专属用户的累计数据
func GetUserRunData(c *gin.Context) {
	var weiRunData Models.WeiXinRunData
	user_id, _ := c.GetQuery("user_id")
	Untils.Db.Model(&weiRunData).Where("wei_chat_id=?", user_id).First(&weiRunData)
	Untils.ResponseOkState(c, weiRunData)
}

// GetListRunData 获取用户排行榜的累计数
func GetListRunData(c *gin.Context) {
	page_size, _ := strconv.Atoi(c.Query("page_size"))
	page_number, _ := strconv.Atoi(c.Query("page_number"))
	page_number = page_number - 1
	order_by := c.DefaultQuery("order_by", "today")
	sort_by := c.DefaultQuery("sort_by", "desc")
	oreder := fmt.Sprintf("%s %s", order_by, sort_by)
	var weiRunData []Models.WeiXinRunData
	Untils.Db.Model(&weiRunData).Preload("WeiChat").Find(&weiRunData).Limit(page_size).Offset(page_number).Order(oreder).Find(&weiRunData)
	var page = make(map[string]any)
	var data = make(map[string]any)
	total := 0
	if len(weiRunData) < 10 && len(weiRunData) != 0 {
		total = 1
	} else {
		total = len(weiRunData) / 10
	}
	page["number"] = "1"
	page["size"] = "10"
	page["total-pages"] = total
	page["total_items"] = len(weiRunData)
	data["page"] = page
	data["page_data"] = weiRunData
	Untils.ResponseOkState(c, data)
}

func GetSession(c *gin.Context) {
	code := c.Query("code")
	Untils.Info.Println("得到的code:=", code)
	var getSession Models.SessionData
	getSession = GetWeiXinSession(code)
	Untils.ResponseOkState(c, getSession)
}

func init() {
	appid = "wx1c5dc5540564df37"
	secret = "90462a31842d82938161d2c3676c8cc3"
}
