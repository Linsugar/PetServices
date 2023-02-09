package main

import (
	"PetService/Routers"
	"PetService/Tasks"
	"PetService/Untils"
)

func main() {
	Routers.Router()
	//监听端口默认为8080
	err := Routers.Gone.Run(":8000")
	if err != nil {
		return
	}
	Tasks.TaskInitAll()
	defer Untils.Db.Close()
	//每天凌晨1点执行一次
}
