package main

import (
	"log"
	"wechat_server/database"
	"wechat_server/process"

	"github.com/gin-gonic/gin"
)

// 与填写的服务器配置中的Token一致
const Token = "wechatserver"

func main() {

	database.InitMysql()

	InitServer()
}

func InitServer() {
	router := gin.Default()

	// router.GET("/wx", WXCheckSignature)
	router.POST("/wx", process.WXMsgReceive)

	log.Fatalln(router.Run(":80"))
}
