package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
	"testing"
)

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

import (
	"github.com/dubbogo-test/consumer-app/api"
)

var grpcGreeterImpl1 = new(api.GreeterClientImpl)

func BenchmarkDubbo_Mosn(T *testing.B) {
	rc := config.NewRootConfigBuilder().
		SetConsumer(config.NewConsumerConfigBuilder().
			SetReferences(
				map[string]*config.ReferenceConfig{
					"GreeterClientImpl": &config.ReferenceConfig{
						InterfaceName: "com.apache.dubbo.sample.basic.IGreeter",
						Check:         nil,
						URL:           fmt.Sprintf("tri://127.0.0.1:2045/com.apache.dubbo.sample.basic.IGreeter"),
						Protocol:      "tri",
					},
				}).
			Build()).
		SetLogger(&config.LoggerConfig{ZapConfig: config.ZapConfig{
			Level: "INFO"}}).
		Build()
	config.SetConsumerService(grpcGreeterImpl1)
	hessian.RegisterPOJO(&api.User{})
	// load config
	if err := rc.Init(); err != nil {
		panic(err)
	}

	for i := 0; i < T.N; i++ {

		req := &api.HelloRequest{
			Name: "bobtthp",
		}
		_, err := grpcGreeterImpl1.SayHello(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkDubbo(T *testing.B) {
	rc := config.NewRootConfigBuilder().
		SetConsumer(config.NewConsumerConfigBuilder().
			SetReferences(
				map[string]*config.ReferenceConfig{
					"GreeterClientImpl": &config.ReferenceConfig{
						InterfaceName: "com.apache.dubbo.sample.basic.IGreeter",
						Check:         nil,
						URL:           fmt.Sprintf("tri://127.0.0.1:20000/com.apache.dubbo.sample.basic.IGreeter"),
						Protocol:      "tri",
					},
				}).
			Build()).
		SetLogger(&config.LoggerConfig{ZapConfig: config.ZapConfig{
			Level: "INFO"}}).
		Build()
	config.SetConsumerService(grpcGreeterImpl1)
	hessian.RegisterPOJO(&api.User{})
	// load config
	if err := rc.Init(); err != nil {
		panic(err)
	}

	for i := 0; i < T.N; i++ {

		req := &api.HelloRequest{
			Name: "bobtthp",
		}
		_, err := grpcGreeterImpl1.SayHello(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		}
	}
}
