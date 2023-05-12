package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 url     --> controller  --> logic   -->    model
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// RotatingNavigation
// @Description: 旋转导航页
func RotatingNavigation(c *gin.Context) {
	c.HTML(http.StatusOK, "Rotating_Navigation_Animation.html", nil)
}
