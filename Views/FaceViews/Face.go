package FaceViews

import (
	"PetService/Models/Face"
	"PetService/Untils"
	"github.com/gin-gonic/gin"
)

func Tencent(c *gin.Context) {
	value := Face.FaceMerge{}
	err := c.Bind(&value)
	if err != nil {
		Untils.Error.Println(err.Error())
		return
	}
	faceValue := FacePk(value.YourFace, value.HisFace)
	Untils.ResponseOkState(c, faceValue)
}

func TencentAge(c *gin.Context) {
	value := Face.FaceAge{}
	err := c.Bind(&value)
	if err != nil {
		Untils.Error.Println(err.Error())
		return
	}
	faceValue := FaceAge(value.Image, value.Age)
	Untils.ResponseOkState(c, faceValue)
}
