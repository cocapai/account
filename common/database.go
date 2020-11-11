package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
	"scutrobot.buff/go_demo/model"
)
// 新建数据库
var DB *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB{
	// 从viper读取其中的配置
	// 数据库类型
	driverName := viper.GetString("datasourse.driverName")
	// 数据库地址
	host := viper.GetString("datasource.host")
	// 数据库端口
	port := viper.GetString("datasource.port")
	// 数据库名
	database := viper.GetString("datasource.database")
	// 用户账号密码
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	// 数据库字符集
	charset := viper.GetString("datasource.charset")
	// 数据库时区，用于统一时间时区
	loc := viper.GetString("datasource.loc")
	// 综合配置 %s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	// 打开数据库
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect , error:" + err.Error())
	}
	// 自动补全表单
	db.AutoMigrate(&model.User{})

	// 赋值给DB
	DB = db
	return db
}

// 输出数据库
func GetDB() *gorm.DB{
	return DB
}