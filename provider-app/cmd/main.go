package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

import (
	"bob.com/dubbogo-test-app/provider-app/cmd/server/api"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)




func main() {
	config.SetProviderService(&api.GreeterProvider{})
	if err := config.Load(); err != nil {
		panic(err)
	}

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// health check
	r.GET("/health.check", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
