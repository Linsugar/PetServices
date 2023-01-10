package Views

import (
	"PetService/Untils"
	"github.com/gin-gonic/gin"
)

//type GetToken struct {
//	UserId string `json:"userId" form:"userId"`
//}

func SetQINiuToken(c *gin.Context) {
	//Data := GetToken{}
	//err := c.ShouldBind(&Data)
	//if err != nil {
	//	Untils.ResponseBadState(c, err)
	//}
	if c.Request.Method == "POST" {
		token := Untils.QiNiuToken()
		Untils.ResponseOkState(c, token)
	}

}
