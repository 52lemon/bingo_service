package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"bingo_service/models"
)

// 通过Id获取用户
func UserGet(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	userModel := models.User{}

	data, err := userModel.UserGet(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// 获取用户列表
func UserGetList(ctx *gin.Context) {

	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "20")

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	userModel := models.User{}

	users, err := userModel.UserGetList(pageInt, pageSizeInt)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": users,
	})
}

// @Description 增加用户
// @Accept  json
// @Produce json
// @Param Name query string false "string valid"
// @Param Age query int false "int valid"
// @Router /user [post]
func UserPost(ctx *gin.Context) {
	userModel := models.User{}
	if err := ctx.ShouldBind(&userModel); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	id, err := userModel.UserAdd()

	if nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "success",
		"uid": id,
	})
}

// 更新用户
func UserPut(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if nil != err || 0 == idInt {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "resource identifier not found",
		})
		return
	}

	userModel := models.User{}

	if err := ctx.ShouldBind(&userModel); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	_, err = userModel.UserUpdate(idInt)

	if nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 更新成功返回 204
	ctx.JSON(http.StatusNoContent, gin.H{"code":"0","msg": "ok"})

}

// @Description 删除用户
// @Accept  json
// @Produce json
// @Param Id query string true "用户id"
// @Router /user/:id [delete]
func UserDelete(ctx *gin.Context) {

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if nil != err || 0 == idInt {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "resource identifier not found",
		})
		return
	}

	userModel := models.User{}

	_, err = userModel.UserDelete(idInt)

	if nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 删除成功返回 204
	ctx.JSON(http.StatusOK, gin.H{"code":"0","msg": "ok"})
}
