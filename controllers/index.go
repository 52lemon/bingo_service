package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 默认页面
// @version 1.0
// @Accept application/x-json-stream
// @Router /ping [get]
func IndexHome(ctx *gin.Context) {

	// ctx.JSON(http.StatusOK, gin.H{
	//	"msg": "ping",
	// })

	//ctx.String(http.StatusOK, "pong")

	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"msg": "bingo",
	})


}
