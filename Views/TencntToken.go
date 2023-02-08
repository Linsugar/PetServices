package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"github.com/gin-gonic/gin"
)

func Tencent(c *gin.Context) {
	value := Models.FaceMerge{}
	err := c.Bind(&value)
	if err != nil {
		Untils.Error.Println(err.Error())
		return
	}
	faceValue := Untils.FacePk(value.YourFace, value.HisFace)
	Untils.ResponseOkState(c, faceValue)
}

func TencentAge(c *gin.Context) {
	value := Models.FaceAge{}
	err := c.Bind(&value)
	if err != nil {
		Untils.Error.Println(err.Error())
		return
	}
	faceValue := Untils.FaceAge(value.Image, value.Age)
	Untils.ResponseOkState(c, faceValue)
}
