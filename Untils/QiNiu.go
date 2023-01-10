package Untils

import (
	"PetService/Conf"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

func QiNiuToken() string {
	bucket := "tang-chat"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
		//CallbackBody: "key=$(key)&hash=$(etag)&width=$(imageInfo.width)&height=$(imageInfo.height)&imageURL=$(imageURL)&size=$(fsize)",
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)","width":"$(imageInfo.width)","height":"$(imageInfo.height)"}`,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := auth.New(Conf.AccessKey, Conf.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	//SetRedisValue("qiniu",upToken,3600);
	fmt.Println("token:" + upToken)
	return upToken
}
