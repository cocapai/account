package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"scutrobot.buff/go_demo/common"
)

func main() {
	// 导入设置文件
	InitConfig()
	// 初始化数据库
	db := common.GetDB()
	// 函数结束关闭数据库
	defer db.Close()

	// gin示例中的初始化
	r := gin.Default()
	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":"+port))
	}
	panic(r.Run())
}

func InitConfig()  {
	// 获取工作地址
	workDir, _ := os.Getwd()
	//viper加载
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	// 读取其中的错误，如果有输出
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}