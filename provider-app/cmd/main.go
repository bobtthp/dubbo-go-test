package main

import (
	"flag"
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/dubbogo-test/provider-app/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var zk string

func init()  {
	flag.StringVar(&zk, "zk", "127.0.0.1:2181","-zk=127.0.0.1:2181")
}


func main() {
	flag.Parse()
	rc := config.NewRootConfigBuilder().
		SetProvider(config.NewProviderConfigBuilder().
			AddService("GreeterProvider", config.NewServiceConfigBuilder().
				SetInterface("com.apache.dubbo.sample.basic.IGreeter").
				SetProtocolIDs("triple").
				Build()).
			Build()).
		AddProtocol("triple", config.NewProtocolConfigBuilder().
			SetName("tri").
			SetPort("20000").
			Build()).
		AddRegistry("bob", &config.RegistryConfig{
		 Protocol: "zookeeper",
		 Address:  zk,
		 Timeout:  "3s",
	 }).
		Build()

	//rc.Init()

	config.SetProviderService(&api.GreeterProvider{})
	 hessian.RegisterPOJO(&api.User{})

	if err := rc.Init(); err != nil {
		fmt.Println(err)

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
