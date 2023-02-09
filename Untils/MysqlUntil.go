package Untils

import (
	"PetService/Conf"
	"PetService/Models/Comment"
	"PetService/Models/Home"
	"PetService/Models/Mine"
	"PetService/Models/Run"
	"PetService/Models/Sale"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var Db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", Conf.UserName, Conf.PassWord, Conf.Host, Conf.Port, Conf.Database, Conf.CharSet)
	dataDase, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("有误", err.Error())
		panic(err)
		//return
	}
	sqlDB := dataDase.DB()
	sqlDB.SetMaxIdleConns(Conf.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(Conf.OpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	Db = dataDase
	var ModelsArrary = []interface{}{&Mine.WeiChat{}, &Mine.ReleaseTopic{},
		&Home.TopicDiscuss{}, &Sale.SaleFriend{}, &Comment.Comment{}, &Comment.RefComment{}, &Comment.SaleComment{},
		&Comment.ListComment{}, &Run.WeiXinRunData{}, &Comment.CollectHomeList{}, &Comment.CollectUserList{}, &Comment.CollectSaleList{}}

	Db.AutoMigrate(ModelsArrary...)
	fmt.Println("链接成功", err)
}
