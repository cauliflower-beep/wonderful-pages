package routers

import (
	"github.com/gin-gonic/gin"
	"wonderful-pages/controller"
	"wonderful-pages/setting"
)

// SetupRouter
// @Description: 初始化路由
func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找 抵一个参数 /status 是访问路径 第二个参数status是文件根目录
	r.Static("/static", "static") // 这样做的好处是可以将静态文件与代码分离，方便管理和维护
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*") // 这个方法会将所有模板文件加载到内存中，以便后续使用
	/*
		r.Static()与r.LoadHTMLGlob()的区别：
		r.Static()方法用于告诉gin框架在访问某个路径时去哪里查找静态文件。比如我在xxx.html中访问了一个静态资源"/static/css,gin就会去static中寻找css静态文件
		r.LoadHTMLGlob()方法用于告诉gin框架去哪里查找模板文件。例如某个handle方法中返回了xxx.html模板文件，gin就会去templates中找xxx.html
	*/

	// 旋转导航页
	r.GET("/rotating", controller.RotatingNavigation)

	return r
}
