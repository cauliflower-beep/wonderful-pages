package routers

import (
	"github.com/gin-gonic/gin"
	"tomatoList/controller"
	"tomatoList/setting"
)

// SetupRouter
//
//	@Description: 初始化路由
func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// 旋转导航页
	r.GET("/rotating", controller.RotatingNavigation)

	return r
}
