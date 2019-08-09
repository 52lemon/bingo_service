package servers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"bingo_service/pkgs/config"
	"bingo_service/routers"
)

func init()  {
	//
}
func HttpRun(httpServer *gin.Engine) {

	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 服务器日志
	now := time.Now()
	logFileName:= fmt.Sprintf("%d_%02d_%02d.log", now.Year(), now.Month(), now.Day() )
	logFile, _ := os.Create("logs/"+logFileName)
	gin.DefaultWriter = io.MultiWriter(logFile)

	//httpServer = gin.Default()
	httpServer = gin.New()
	httpServer.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	httpServer.Use(gin.Recovery())

	// 服务配置
	serverConfig := config.GetServerConfig()

	// gin 运行模式
	gin.SetMode(serverConfig["env"])

	// 注册路由
	routers.RegisterRoutes(httpServer)

	// 端口设置
	serverAddr := serverConfig["host"] + ":" + serverConfig["port"]
	// 读取超时
	readTimeOut, _ := strconv.Atoi(serverConfig["read_time_out"])
	// 写入超时
	writeTimeOut, _ := strconv.Atoi(serverConfig["write_time_out"])

	// HTTP配置
	server := &http.Server{
		Addr:           serverAddr,
		Handler:        httpServer,
		ReadTimeout:    time.Duration(readTimeOut) * time.Second,
		WriteTimeout:   time.Duration(writeTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20, // HTTP请求头的最大允许值
	}
	err := server.ListenAndServe()
	if nil != err {
		panic("server run error: " + err.Error())
	}

}
