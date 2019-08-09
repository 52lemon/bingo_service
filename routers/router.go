package routers

import (
	"github.com/gin-gonic/gin"
	"bingo_service/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

)

func RegisterRoutes(router *gin.Engine) {

	// 模板路径
	router.LoadHTMLGlob("views/**/*")
	// use ginSwagger middleware to
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//url := ginSwagger.URL("http://localhost:9000/swagger/doc.json")
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/", controllers.IndexHome)
	router.GET("/ping", controllers.IndexHome)

	router.GET("/user/:id", controllers.UserGet)
	router.GET("/users", controllers.UserGetList)
	router.POST("/user", controllers.UserPost)
	router.PUT("/user/:id", controllers.UserPut)
	router.PATCH("/user/:id", controllers.UserPut)
	router.DELETE("/user/:id", controllers.UserDelete)


}

