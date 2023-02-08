package Untils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	facefusion "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/facefusion/v20220927"
	ft "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ft/v20200304"
	iai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai/v20200303"
)

var (
	secretId  string
	secretKey string
)

func FaceMerge(leftUrl, retUrl string) any {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	//人脸融合
	cpf.HttpProfile.Endpoint = "facefusion.tencentcloudapi.com"

	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := facefusion.NewClient(credential, "ap-chengdu", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := facefusion.NewFuseFaceRequest()
	request.ProjectId = common.StringPtr("at_1619894828715380736")
	request.ModelId = common.StringPtr("mt_1620323556490645504")
	request.RspImgType = common.StringPtr("url")

	request.MergeInfos = []*facefusion.MergeInfo{
		{
			Url: common.StringPtr(leftUrl),
		},
	}

	// 返回的resp是一个FuseFaceResponse的实例，与请求对象对应
	response, err := client.FuseFace(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
	return response.Response
}

func FacePk(urla, urlb string) *iai.CompareFaceResponseParams {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "iai.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := iai.NewClient(credential, "ap-chengdu", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := iai.NewCompareFaceRequest()

	//request.ImageB = common.StringPtr(urla)
	request.UrlA = common.StringPtr(urla)
	request.UrlB = common.StringPtr(urlb)

	// 返回的resp是一个CompareFaceResponse的实例，与请求对象对应
	response, err := client.CompareFace(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
	if *response.Response.Score >= 60 {
		*response.Response.FaceModelVersion = "很有夫妻相哦!"
	} else {
		*response.Response.FaceModelVersion = "请继续加油!"
	}

	return response.Response
}

func FaceAge(url string, age int) *ft.ChangeAgePicResponseParams {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ft.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := ft.NewClient(credential, "ap-chengdu", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := ft.NewChangeAgePicRequest()

	request.Url = common.StringPtr(url)
	request.AgeInfos = []*ft.AgeInfo{
		&ft.AgeInfo{
			Age: common.Int64Ptr(int64(age)),
		},
	}
	request.RspImgType = common.StringPtr("base64")

	// 返回的resp是一个ChangeAgePicResponse的实例，与请求对象对应
	response, err := client.ChangeAgePic(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response.Response
}

func AesDecrypt(crypted, key, iv string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(sessionKeyBytes)
	if err != nil {
		return nil, err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, ivBytes)
	origData := make([]byte, len(decodeBytes))
	blockMode.CryptBlocks(origData, decodeBytes)
	//获取的数据尾端有'/x0e'占位符,去除它
	for i, ch := range origData {
		if ch == '\x0e' {
			origData[i] = ' '
		}
	}
	Error.Println("结果数据：", string(origData))
	//{"phoneNumber":"15082726017","purePhoneNumber":"15082726017","countryCode":"86","watermark":{"timestamp":1539657521,"appid":"wx4c6c3ed14736228c"}}//<nil>
	return origData, nil
}

func init() {
	secretId = "AKIDKagmgQzKQM67wZ1GH3cw1tZ4CNM7cstf"
	secretKey = "bx4wZ0FH7k8HC3kHetyqUt54rWN2Xv86"
}
