package main

import (
	"flag"
	"fmt"
	"github.com/dubbogo-test/provider-app/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var nacos string

func init() {
	flag.StringVar(&nacos, "nacos", "127.0.0.1:8848", "-nacos 127.0.0.1:8848")
}

func main() {
	flag.Parse()
	rc := config.NewRootConfigBuilder().
		SetApplication(config.NewApplicationConfigBuilder().
			SetName("provider-test-app").SetModule("opensource").
			Build()).
		SetProvider(config.NewProviderConfigBuilder().
			AddService("GreeterProvider", config.NewServiceConfigBuilder().
				SetInterface("com.apache.dubbo.sample.basic.IGreeter").
				SetProtocolIDs("triple").
				Build()).
			//SetFilter("go2sky-tracing-server").
			Build()).
		AddProtocol("triple", config.NewProtocolConfigBuilder().
			SetName("tri").
			//SetIp("127.0.0.1").
			SetPort("20000").
			Build()).
		SetLogger(&config.LoggerConfig{ZapConfig: config.ZapConfig{
			Level: "DEBUG"}}).
		AddRegistry("bob.test", &config.RegistryConfig{
			Group:        "DEFAULT_GROUP",
			Protocol:     "nacos",
			Address:      nacos,
			Timeout:      "3s",
			RegistryType: "interface",
		}).
		Build()

	config.SetProviderService(&api.GreeterProvider{})
	//hessian.RegisterPOJO(&api.User{})

	if err := rc.Init(); err != nil {
		fmt.Println(err)

		panic(err)
	}

	//// setup reporter, use gRPC reporter for production
	//report, err := reporter.NewGRPCReporter("192.168.196.223:11800")
	//if err != nil {
	//	log.Fatalf("new reporter error: %v \n", err)
	//}
	//
	//// setup tracer
	//tracer, err := go2sky.NewTracer("dubbo-go", go2sky.WithReporter(report))
	//if err != nil {
	//	log.Fatalf("crate tracer error: %v \n", err)
	//}
	//
	//// set dubbogo plugin server tracer
	//err = dubbo_go.SetServerTracer(tracer)
	//if err != nil {
	//	log.Fatalf("set tracer error: %v \n", err)
	//}
	//
	//// set extra tags and report tags
	//dubbo_go.SetServerExtraTags("extra-tags", "server")
	//dubbo_go.SetServerReportTags("release")

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
