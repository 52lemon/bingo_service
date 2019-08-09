package main

import (
	_ "bingo_service/docs"
	"github.com/gin-gonic/gin"
	"bingo_service/servers"
	"bingo_service/jobs"
)

var GinEngine *gin.Engine
var Config = make(map[string]string)

// @title bingo_service
// @version 1.0
// @description go web 骨架.
// @termsOfService http://swagger.io/terms/

// @contact.name jinchunguang
// @contact.email jin-chunguang@foxmail.com

// @host localhost:9000
func main() {

	// 运行 task
	jobs.TaskRun()

	// 运行 nsq
	servers.NsqRun()

	// 运行server
	servers.HttpRun(GinEngine)
}
