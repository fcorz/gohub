package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func SetupRoute(router *gin.Engine) {

	// 註冊全局中間件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	routes.RegisterAPIRoutes(router)

	// 404
	setup404handler(router)
}

// 註冊全局中間件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		gin.Recovery(),
	)
}

// 處理404請求
func setup404handler(router *gin.Engine) {

	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面404")
		} else {
			// 返回json
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
